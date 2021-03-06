package main

/*

go run -mod vendor cmd/subscriptiond/main.go -devel -templates 'templates/html/*.html'

*/

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-mailinglist"
	"github.com/aaronland/go-mailinglist/assets/templates"
	"github.com/aaronland/go-mailinglist/http"
	"github.com/aaronland/go-mailinglist/server"
	"github.com/aaronland/go-string/random"
	"github.com/aaronland/gocloud-runtimevar-string"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
	"html/template"
	"io/ioutil"
	"log"
	gohttp "net/http"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// maybe -mailinglist-config...
	// mailinglist_dsn := flag.String("mailinglist-dsn", "", "...")

	mailinglist_name := flag.String("mailinglist-name", "", "...")
	mailinglist_url := flag.String("mailinglist-url", "", "...")
	mailinglist_sender := flag.String("mailinglist-sender", "", "...")

	subs_dsn := flag.String("subscriptions-dsn", "", "...")
	conf_dsn := flag.String("confirmations-dsn", "", "...")
	invites_dsn := flag.String("invitations-dsn", "", "...")
	logs_dsn := flag.String("eventlogs-dsn", "", "...")

	sender_dsn := flag.String("sender-dsn", "", "...")
	crumb_url := flag.String("crumb-url", "", "...")

	protocol := flag.String("protocol", "http", "...")
	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	enable_subscriptions := flag.Bool("enable-subscribe", true, "...")
	enable_unsubscriptions := flag.Bool("enable-unsubscribe", true, "...")
	enable_confirmations := flag.Bool("enable-confirm", true, "...")
	enable_invitation_requests := flag.Bool("enable-invites-request", true, "...")
	enable_invitation_accepts := flag.Bool("enable-invites-accept", true, "...")

	path_index := flag.String("path-index", "/", "...")
	path_subscribe := flag.String("path-subscribe", "/subscribe", "...")
	path_unsubscribe := flag.String("path-unsubscribe", "/unsubscribe", "...")
	path_invite_request := flag.String("path-invite-request", "/invite", "...")
	path_invite_accept := flag.String("path-invite-accept", "/invite/accept", "...")
	path_confirm := flag.String("path-confirm", "/confirm", "...")

	static_prefix := flag.String("static-prefix", "", "Prepend this prefix to URLs for static assets.")

	var path_templates flags.MultiString
	flag.Var(&path_templates, "templates", "One or more optional strings for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")

	path_ping := flag.String("path-ping", "/ping", "...")

	devel := flag.Bool("devel", false, "...")

	flag.Parse()

	s, err := server.NewServer(*protocol, *host, *port)

	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}

	if *mailinglist_url == "" {
		*mailinglist_url = s.Address()
	}

	site_url, err := url.Parse(*mailinglist_url)

	if err != nil {
		log.Fatalf("Failed to parse site URL (%s): %s", *mailinglist_url, err)
	}

	if *devel {

		root, err := ioutil.TempDir("", "subscriptiond")

		if err != nil {
			log.Fatalf("Failed to create temporary subscriptiond directory: %s", err)
		}

		log.Printf("Started in -devel mode, all databases are filesystem-based. Database root is '%s' and will be removed when this program exits.\n", root)
		defer os.RemoveAll(root)

		subs_dir := filepath.Join(root, "subscriptions")
		conf_dir := filepath.Join(root, "confirmations")
		invites_dir := filepath.Join(root, "invitations")
		logs_dir := filepath.Join(root, "eventlogs")

		err = os.Mkdir(subs_dir, 0700)

		if err != nil {
			log.Fatalf("Failed to create temporary subscriptions directory (%s): %s", subs_dir, err)
		}

		err = os.Mkdir(invites_dir, 0700)

		if err != nil {
			log.Fatalf("Failed to create temporary invitations directory (%s): %s", invites_dir, err)
		}

		err = os.Mkdir(conf_dir, 0700)

		if err != nil {
			log.Fatalf("Failed to create temporary confirmations directory (%s): %s", subs_dir, err)
		}

		err = os.Mkdir(logs_dir, 0700)

		if err != nil {
			log.Fatalf("Failed to create temporary confirmations directory (%s): %s", logs_dir, err)
		}

		opts := random.DefaultOptions()
		opts.AlphaNumeric = true
		opts.Length = 32
		// opts.Chars = 32

		secret, err := random.String(opts)

		if err != nil {
			log.Fatalf("Failed to create crumb secret: %s", err)
		}

		opts.Length = 8
		opts.Chars = 8

		salt, err := random.String(opts)

		if err != nil {
			log.Fatalf("Failed to create crumb salt: %s", err)
		}

		*subs_dsn = fmt.Sprintf("database=fs root=%s", subs_dir)
		*invites_dsn = fmt.Sprintf("database=fs root=%s", invites_dir)
		*conf_dsn = fmt.Sprintf("database=fs root=%s", conf_dir)
		*logs_dsn = fmt.Sprintf("database=fs root=%s", logs_dir)

		*crumb_url = fmt.Sprintf("constant://?val=secret=%s+salt=%s+extra=foo+separator=:+ttl=300", secret, salt)
		*sender_dsn = "sender=stdout"

		*mailinglist_name = "Development"
		*mailinglist_sender = "development@localhost"
	}

	if *mailinglist_name == "" {
		log.Fatal("Missing -mailinglist-name")
	}

	_, err = mail.ParseAddress(*mailinglist_sender)

	if err != nil {
		log.Fatal("Invalid -mailinglist-sender")
	}

	subs_db, err := mailinglist.NewSubscriptionsDatabaseFromDSN(*subs_dsn)

	if err != nil {
		log.Fatalf("Failed to create subscriptions database: %s", err)
	}

	invites_db, err := mailinglist.NewInvitationsDatabaseFromDSN(*invites_dsn)

	if err != nil {
		log.Fatalf("Failed to create invitations database: %s", err)
	}

	conf_db, err := mailinglist.NewConfirmationsDatabaseFromDSN(*conf_dsn)

	if err != nil {
		log.Fatalf("Failed to create confirmations database:", err)
	}

	logs_db, err := mailinglist.NewEventLogsDatabaseFromDSN(*logs_dsn)

	if err != nil {
		log.Fatalf("Failed to create confirmations database:", err)
	}

	sender, err := mailinglist.NewSenderFromDSN(*sender_dsn)

	if err != nil {
		log.Fatalf("Failed to create mail sender: %s", err)
	}

	t := template.New("subscriptiond").Funcs(template.FuncMap{
		"SiteURL": func(path string) string {
			u, _ := url.Parse(site_url.String())
			u.Path = path
			return u.String()
		},
	})

	if len(path_templates) > 0 {

		for _, p := range path_templates {

			t, err = t.ParseGlob(p)

			if err != nil {
				log.Fatalf("Failed to parse templates (%s): %s", p, err)
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
				log.Fatalf("Failed to parse template (%s): %s", name, err)
			}
		}
	}

	if *static_prefix != "" {

		*static_prefix = strings.TrimRight(*static_prefix, "/")

		if !strings.HasPrefix(*static_prefix, "/") {
			log.Fatal("Invalid -static-prefix value")
		}
	}

	feature_flags := &mailinglist.FeatureFlags{
		Subscribe:     *enable_subscriptions,
		Unsubscribe:   *enable_unsubscriptions,
		InviteRequest: *enable_invitation_requests,
		InviteAccept:  *enable_invitation_accepts,
		Confirm:       *enable_confirmations,
	}

	path_cfg := &mailinglist.PathConfig{
		Index:         *path_index,
		Subscribe:     *path_subscribe,
		Unsubscribe:   *path_unsubscribe,
		Confirm:       *path_confirm,
		InviteRequest: *path_invite_request,
		InviteAccept:  *path_invite_accept,
	}

	list_cfg := &mailinglist.MailingListConfig{
		Name:         *mailinglist_name,
		URL:          site_url,
		Sender:       *mailinglist_sender,
		Paths:        path_cfg,
		FeatureFlags: feature_flags,
	}

	crumb_dsn, err := runtimevar.OpenString(context.Background(), *crumb_url)

	if err != nil {
		log.Fatalf("Failed to open crumb URL: %s", err)
	}

	crumb_cfg, err := crumb.NewCrumbConfigFromDSN(crumb_dsn)

	if err != nil {
		log.Fatalf("Failed to create crumb: %s", err)
	}

	mux := gohttp.NewServeMux()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}

	crumb_error_opts := &http.CrumbErrorHandlerOptions{
		Templates: t,
		Config:    list_cfg,
	}

	crumb_error_handler, err := http.CrumbErrorHandler(crumb_error_opts)

	if err != nil {
		log.Fatalf("Failed to crete crumb error handler: %s", err)
	}

	crumb_error_handler = bootstrap.AppendResourcesHandler(crumb_error_handler, bootstrap_opts)

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatalf("Failed to create ping handler:%s", err)
	}

	mux.Handle(*path_ping, ping_handler)

	index_opts := &http.IndexHandlerOptions{
		Templates: t,
		Config:    list_cfg,
	}

	index_handler, err := http.IndexHandler(index_opts)

	if err != nil {
		log.Fatalf("Failed to create index handler: %s", err)
	}

	index_handler = bootstrap.AppendResourcesHandler(index_handler, bootstrap_opts)

	mux.Handle(path_cfg.Index, index_handler)

	subscribe_opts := &http.SubscribeHandlerOptions{
		Config:        list_cfg,
		Templates:     t,
		Subscriptions: subs_db,
		Confirmations: conf_db,
		EventLogs:     logs_db,
		Sender:        sender,
	}

	subscribe_handler, err := http.SubscribeHandler(subscribe_opts)

	if err != nil {
		log.Fatalf("Failed to create subscribe handler: %s", err)
	}

	subscribe_handler = bootstrap.AppendResourcesHandler(subscribe_handler, bootstrap_opts)
	subscribe_handler = crumb.EnsureCrumbHandlerWithErrorHandler(crumb_cfg, subscribe_handler, crumb_error_handler)

	mux.Handle(path_cfg.Subscribe, subscribe_handler)

	unsubscribe_opts := &http.UnsubscribeHandlerOptions{
		Config:        list_cfg,
		Templates:     t,
		Subscriptions: subs_db,
		Confirmations: conf_db,
		EventLogs:     logs_db,
		Sender:        sender,
	}

	unsubscribe_handler, err := http.UnsubscribeHandler(unsubscribe_opts)

	if err != nil {
		log.Fatalf("Failed to create unsubscribe handler: %s", err)
	}

	unsubscribe_handler = bootstrap.AppendResourcesHandler(unsubscribe_handler, bootstrap_opts)
	unsubscribe_handler = crumb.EnsureCrumbHandlerWithErrorHandler(crumb_cfg, unsubscribe_handler, crumb_error_handler)

	mux.Handle(path_cfg.Unsubscribe, unsubscribe_handler)

	invite_request_opts := &http.InviteRequestHandlerOptions{
		Config:        list_cfg,
		Templates:     t,
		Subscriptions: subs_db,
		Invitations:   invites_db,
		EventLogs:     logs_db,
		Sender:        sender,
	}

	invite_request_handler, err := http.InviteRequestHandler(invite_request_opts)

	if err != nil {
		log.Fatalf("Failed to create invite request handler: %s", err)
	}

	invite_request_handler = bootstrap.AppendResourcesHandler(invite_request_handler, bootstrap_opts)
	invite_request_handler = crumb.EnsureCrumbHandlerWithErrorHandler(crumb_cfg, invite_request_handler, crumb_error_handler)

	mux.Handle(path_cfg.InviteRequest, invite_request_handler)

	invite_accept_opts := &http.InviteAcceptHandlerOptions{
		Config:        list_cfg,
		Templates:     t,
		Subscriptions: subs_db,
		Confirmations: conf_db,
		Invitations:   invites_db,
		EventLogs:     logs_db,
		Sender:        sender,
	}

	invite_accept_handler, err := http.InviteAcceptHandler(invite_accept_opts)

	if err != nil {
		log.Fatalf("Failed to create invite accept handler: %s", err)
	}

	invite_accept_handler = bootstrap.AppendResourcesHandler(invite_accept_handler, bootstrap_opts)
	invite_accept_handler = crumb.EnsureCrumbHandlerWithErrorHandler(crumb_cfg, invite_accept_handler, crumb_error_handler)

	mux.Handle(path_cfg.InviteAccept, invite_accept_handler)

	confirm_opts := &http.ConfirmHandlerOptions{
		Config:        list_cfg,
		Templates:     t,
		Subscriptions: subs_db,
		EventLogs:     logs_db,
		Confirmations: conf_db,
	}

	confirm_handler, err := http.ConfirmHandler(confirm_opts)

	if err != nil {
		log.Fatalf("Failed to create confirm handler: %s", err)
	}

	confirm_handler = bootstrap.AppendResourcesHandler(confirm_handler, bootstrap_opts)
	confirm_handler = crumb.EnsureCrumbHandlerWithErrorHandler(crumb_cfg, confirm_handler, crumb_error_handler)

	mux.Handle(path_cfg.Confirm, confirm_handler)

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}
}
