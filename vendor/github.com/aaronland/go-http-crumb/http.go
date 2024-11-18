package crumb

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"log/slog"
	go_http "net/http"

	"github.com/aaronland/go-http-rewrite"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-http-fault/v2"
)

// EnsureCrumbHandler wraps 'next_handler' with a middleware `http.Handler` for assigning and validating
// crumbs using the default `fomuseum/go-http-fault/v2.FaultHandler` as an error handler. Any errors that
// trigger the error handler can be retrieved using `sfomuseum/go-http-fault/v2.RetrieveError()`.
func EnsureCrumbHandler(cr Crumb, next_handler go_http.Handler) go_http.Handler {

	logger := slog.Default()
	fault_logger := slog.NewLogLogger(logger.Handler(), slog.LevelError)

	fault_handler := fault.FaultHandler(fault_logger)
	return EnsureCrumbHandlerWithErrorHandler(cr, next_handler, fault_handler)
}

// EnsureCrumbHandlerWithFaultWrapper wraps 'next_handler' with a middleware `http.Handler` for assigning and validating
// crumbs. Error handling is assumed to be handled by a separate middleware handler provided by a `sfomuseum/go-http-fault/v2.FaultWrapper`.
// instance. Any errors that are triggered as recorded using the `sfomuseum/go-http-fault/v2.AssignError()` method and can
// be retrieved using `sfomuseum/go-http-fault/v2.RetrieveError()` method.
func EnsureCrumbHandlerWithFaultWrapper(cr Crumb, next_handler go_http.Handler) go_http.Handler {

	fn := func(rsp go_http.ResponseWriter, req *go_http.Request) {

		switch req.Method {

		case "POST", "PUT":

			var crumb_var string
			var crumb_err error

			if req.Method == "POST" {
				crumb_var, crumb_err = sanitize.PostString(req, "crumb")
			} else {
				crumb_var, crumb_err = sanitize.GetString(req, "crumb")
			}

			if crumb_err != nil {
				err := Error(UnsanitizedCrumb, crumb_err)
				fault.AssignError(req, err, go_http.StatusBadRequest)
				rsp.WriteHeader(go_http.StatusBadRequest)
				return
			}

			if crumb_var == "" {
				err := Error(MissingCrumb, fmt.Errorf("Missing crumb"))
				fault.AssignError(req, err, go_http.StatusBadRequest)
				rsp.WriteHeader(go_http.StatusBadRequest)
				return
			}

			ok, err := cr.Validate(req, crumb_var)

			if err != nil {
				err := Error(InvalidCrumb, err)
				fault.AssignError(req, err, go_http.StatusBadRequest)
				rsp.WriteHeader(go_http.StatusBadRequest)
				return
			}

			if !ok {
				err := Error(ExpiredCrumb, fmt.Errorf("Expired"))
				fault.AssignError(req, err, go_http.StatusForbidden)
				rsp.WriteHeader(go_http.StatusForbidden)
				return
			}

		default:
			// pass
		}

		crumb_var, err := cr.Generate(req)

		if err != nil {
			err := Error(GenerateCrumb, fmt.Errorf("Expired"))
			fault.AssignError(req, err, go_http.StatusInternalServerError)
			rsp.WriteHeader(go_http.StatusInternalServerError)
			return
		}

		rewrite_func := NewCrumbRewriteFunc(crumb_var)
		rewrite_handler := rewrite.RewriteHTMLHandler(next_handler, rewrite_func)

		rewrite_handler.ServeHTTP(rsp, req)
	}

	h := go_http.HandlerFunc(fn)
	return h
}

// EnsureCrumbHandlerWithErrorHandler wraps 'next_handler' with a middleware a middleware `http.Handler` for
// assigning and validating crumbs using a custom error handler. Any errors that trigger the error handler can
// be retrieved using `sfomuseum/go-http-fault/v2.RetrieveError()`.
func EnsureCrumbHandlerWithErrorHandler(cr Crumb, next_handler go_http.Handler, error_handler go_http.Handler) go_http.Handler {

	fn := func(rsp go_http.ResponseWriter, req *go_http.Request) {

		switch req.Method {

		case "POST", "PUT":

			var crumb_var string
			var crumb_err error

			if req.Method == "POST" {
				crumb_var, crumb_err = sanitize.PostString(req, "crumb")
			} else {
				crumb_var, crumb_err = sanitize.GetString(req, "crumb")
			}

			if crumb_err != nil {
				fault.AssignError(req, Error(UnsanitizedCrumb, crumb_err), go_http.StatusBadRequest)
				error_handler.ServeHTTP(rsp, req)
				return
			}

			if crumb_var == "" {
				fault.AssignError(req, Error(MissingCrumb, fmt.Errorf("Missing crumb")), go_http.StatusBadRequest)
				error_handler.ServeHTTP(rsp, req)
				return
			}

			ok, err := cr.Validate(req, crumb_var)

			if err != nil {
				fault.AssignError(req, Error(InvalidCrumb, err), go_http.StatusInternalServerError)
				error_handler.ServeHTTP(rsp, req)
				return
			}

			if !ok {
				fault.AssignError(req, Error(ExpiredCrumb, fmt.Errorf("Expired")), go_http.StatusForbidden)
				error_handler.ServeHTTP(rsp, req)
				return
			}

		default:
			// pass
		}

		crumb_var, err := cr.Generate(req)

		if err != nil {
			fault.AssignError(req, Error(GenerateCrumb, err), go_http.StatusInternalServerError)
			error_handler.ServeHTTP(rsp, req)
			return
		}

		rewrite_func := NewCrumbRewriteFunc(crumb_var)
		rewrite_handler := rewrite.RewriteHTMLHandler(next_handler, rewrite_func)

		rewrite_handler.ServeHTTP(rsp, req)
	}

	h := go_http.HandlerFunc(fn)
	return h
}

// NewCrumbRewriteFunc returns a `aaronland/go-http-rewrite.RewriteHTMLFunc` used to
// append crumb data to HTML output.
func NewCrumbRewriteFunc(crumb_var string) rewrite.RewriteHTMLFunc {

	var rewrite_func rewrite.RewriteHTMLFunc

	rewrite_func = func(n *html.Node, w io.Writer) {

		if n.Type == html.ElementNode && n.Data == "body" {

			crumb_ns := ""
			crumb_key := "data-crumb"
			crumb_value := crumb_var

			crumb_attr := html.Attribute{crumb_ns, crumb_key, crumb_value}
			n.Attr = append(n.Attr, crumb_attr)
		}

		if n.Type == html.ElementNode && n.Data == "form" {

			ns := ""

			attrs := []html.Attribute{
				html.Attribute{ns, "type", "hidden"},
				html.Attribute{ns, "id", "crumb"},
				html.Attribute{ns, "name", "crumb"},
				html.Attribute{ns, "value", crumb_var},
			}

			i := &html.Node{
				Type:      html.ElementNode,
				DataAtom:  atom.Input,
				Data:      "input",
				Namespace: ns,
				Attr:      attrs,
			}

			n.AppendChild(i)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			rewrite_func(c, w)
		}
	}

	return rewrite_func
}
