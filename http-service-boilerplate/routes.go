package main

func addRoutes(
	mux                 *http.ServeMux,
	logger              *logging.Logger,
	config              Config,
	tenantsStore        *TenantsStore,
	commentsStore       *CommentsStore,
	conversationService *ConversationService,
	chatGPTService      *ChatGPTService,
	authProxy           *authProxy
) {
	mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	mux.Handle("/", http.NotFoundHandler())
}

// func addRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("/api/", handleAPI())
// 	mux.HandleFunc("/about", handleAbout())
// 	mux.HandleFunc("/", handleIndex())
// 	mux.HandleFunc("/admin", adminOnly(handleAdminIndex()))
// }
