package oauth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RangelReale/osin"
	"github.com/minzhang2110/smart-home/pkg/oauth/storage"
)

// Oauth oauth2 server
type Oauth struct {
	server *osin.Server
}

// New .
func New() *Oauth {
	cfg := osin.NewServerConfig()
	cfg.AllowGetAccessRequest = true
	cfg.AllowClientSecretInParams = true
	cfg.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	cfg.AccessExpiration = 3600*24*7

	sto := storage.NewMemStorage()
	client := &osin.DefaultClient{
		Id:          os.Getenv("CLIENT_ID"),
		Secret:      os.Getenv("CLIENT_SECRET"),
		RedirectUri: os.Getenv("CLIENT_REDIRECT_URI"),
	}
	sto.SetClient(client.Id, client)

	return &Oauth{
		server: osin.NewServer(cfg, sto),
	}
}

// HandleAuthorizeToken handle authorize and token exchange request
func (o *Oauth) HandleAuthorizeToken() {
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := o.server.NewResponse()
		defer resp.Close()

		if ar := o.server.HandleAuthorizeRequest(resp, r); ar != nil {
			if !HandleLoginPage(w, r) {
				return
			}
			ar.Authorized = true
			ar.UserData = os.Getenv("USER_NAME")
			o.server.FinishAuthorizeRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			log.Printf("ERROR: %s\n", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Access token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := o.server.NewResponse()
		defer resp.Close()

		if ar := o.server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			o.server.FinishAccessRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			log.Printf("ERROR: %s\n", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})
}

type key int

// UserName key
const UserName key = 1

// Middleware .
func (o *Oauth) Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := o.server.NewResponse()
		defer resp.Close()

		if ir := o.server.HandleInfoRequest(resp, r); ir != nil {
			if resp.IsError {
				osin.OutputJSON(resp, w, r)
			} else {
				ctx := context.WithValue(r.Context(), UserName, ir.AccessData.UserData)
				handler(w, r.WithContext(ctx))
				return
			}
		}
		osin.OutputJSON(resp, w, r)
	}
}

// HandleLoginPage .
func HandleLoginPage(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	if r.Method == "POST" && r.FormValue("user") == os.Getenv("USER_NAME") &&
		r.FormValue("password") == os.Getenv("USER_PASS") {
		return true
	}

	w.Write([]byte("<html><body>"))

	w.Write([]byte(fmt.Sprintf("<form action=\"/authorize?%s\" method=\"POST\">", r.URL.RawQuery)))

	w.Write([]byte("Login: <input type=\"text\" name=\"user\" /><br/>"))
	w.Write([]byte("Password: <input type=\"password\" name=\"password\" /><br/>"))
	w.Write([]byte("<input type=\"submit\"/>"))

	w.Write([]byte("</form>"))

	w.Write([]byte("</body></html>"))

	return false
}
