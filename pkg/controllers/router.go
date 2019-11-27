package controllers

import (
	"fmt"
	"github.com/minzhang2110/smart-home/pkg/oauth"
	"net/http"
)

// Router type
type Router struct {
	handler *Handler
}

// NewRouter returns new router
func NewRouter(h *Handler) *Router {
	o := oauth.New()
	o.HandleAuthorizeToken()

	http.HandleFunc("/health", o.Middleware(func(w http.ResponseWriter, r *http.Request) {
		if userName, ok := r.Context().Value(oauth.UserName).(string); ok {
			fmt.Fprintf(w, "hello：%s", userName)
		}
	}))

	http.HandleFunc("/smarthome", o.Middleware(h.SmartHomeHandler))

	return &Router{
		handler: h,
	}
}

// Start .
func (r *Router) Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
