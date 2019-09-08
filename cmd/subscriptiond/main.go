package main

import (
	"flag"
	"github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/http"
	"github.com/aaronland/go-mailinglist/server"
	"github.com/aaronland/go-mailinglist/assets/templates"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"		
	"log"
	gohttp "net/http"
	"html/template"
	"strings"
)

func main() {

	subs_dsn := flag.String("subscriptions-dsn", "", "...")
	conf_dsn := flag.String("confirmations-dsn", "", "...")
	sender_dsn := flag.String("sender-dsn", "", "...")
	crumb_dsn := flag.String("crumb-dsn", "", "...")

	protocol := flag.String("protocol", "http", "...")
	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	index_handler := flag.Bool("index-handler", true, "...")
	subscribe_handler := flag.Bool("subscribe-handler", true, "...")
	unsubscribe_handler := flag.Bool("unsubscribe-handler", true, "...")
	confirm_handler := flag.Bool("confirm-handler", true, "...")

	path_index := flag.String("path-subscribe", "/index", "...")
	path_subscribe := flag.String("path-subscribe", "/subscribe", "...")
	path_unsubscribe := flag.String("path-unsubscribe", "/unsubscribe", "...")
	path_confirm := flag.String("path-confirm", "/confirm", "...")

	static_prefix := flag.String("static-prefix", "", "Prepend this prefix to URLs for static assets.")

	var path_templates flags.MultiString
	flag.Var(&path_templates, "templates", "One or more optional strings for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")
	
	path_ping := flag.String("path-ping", "/ping", "...")

	flag.Parse()

	subs_db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*subs_dsn)

	if err != nil {
		log.Fatal(err)
	}

	conf_db, err := mailinglist.NewConfirmationsDatabaseFromDSN(*conf_dsn)

	if err != nil {
		log.Fatal(err)
	}

	sender, err := mailinglist.NewSenderFromDSN(*sender_dsn)

	if err != nil {
		log.Fatal(err)
	}

	t := template.New("subscriptiond").Funcs(template.FuncMap{
	})

	if len(path_templates) > 0 {

		for _, p := range path_templates {
			
			t, err = t.ParseGlob(p)
			
			if err != nil {
				log.Fatal(err)
			}
		}
		
	} else {

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				log.Fatal(err)
			}

			t, err = t.Parse(string(body))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if *static_prefix != "" {

		*static_prefix = strings.TrimRight(*static_prefix, "/")

		if !strings.HasPrefix(*static_prefix, "/") {
			log.Fatal("Invalid -static-prefix value")
		}
	}
	
	mux := gohttp.NewServeMux()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}
	
	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	crumb_cfg, err := crumb.NewCrumbConfigFromDSN(*crumb_dsn)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(*path_ping, ping_handler)

	if *index_handler {

		opts := &http.IndexHandlerOptions{
			// paths for other things
		}

		index_handler, err := http.IndexHandler(opts)

		if err != nil {
			log.Fatal(err)
		}

		index_handler = bootstrap.AppendResourcesHandler(index_handler, bootstrap_opts)
		
		mux.Handle(*path_index, index_handler)
	}

	if *subscribe_handler {

		opts := &http.SubscribeHandlerOptions{
			Subscriptions: subs_db,
			Confirmations: conf_db,
			Sender:        sender,
		}

		subscribe_handler, err := http.SubscribeHandler(opts)

		if err != nil {
			log.Fatal(err)
		}

		subscribe_handler = bootstrap.AppendResourcesHandler(subscribe_handler, bootstrap_opts)				
		subscribe_handler = crumb.EnsureCrumbHandler(crumb_cfg, subscribe_handler)
		
		mux.Handle(*path_subscribe, subscribe_handler)
	}

	if *unsubscribe_handler {

		opts := &http.UnsubscribeHandlerOptions{
			Subscriptions: subs_db,
			Confirmations: conf_db,
			Sender:        sender,
		}

		unsubscribe_handler, err := http.UnsubscribeHandler(opts)

		if err != nil {
			log.Fatal(err)
		}

		unsubscribe_handler = bootstrap.AppendResourcesHandler(unsubscribe_handler, bootstrap_opts)		
		unsubscribe_handler = crumb.EnsureCrumbHandler(crumb_cfg, unsubscribe_handler)

		mux.Handle(*path_unsubscribe, unsubscribe_handler)
	}

	if *confirm_handler {

		opts := &http.ConfirmHandlerOptions{
			Subscriptions: subs_db,
			Confirmations: conf_db,
		}

		confirm_handler, err := http.ConfirmHandler(opts)

		if err != nil {
			log.Fatal(err)
		}

		confirm_handler = bootstrap.AppendResourcesHandler(confirm_handler, bootstrap_opts)				
		confirm_handler = crumb.EnsureCrumbHandler(crumb_cfg, confirm_handler)

		mux.Handle(*path_confirm, confirm_handler)
	}

	s, err := server.NewServer(*protocol, *host, *port)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}
}
