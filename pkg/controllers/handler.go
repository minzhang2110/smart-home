package controllers

import (
	"log"
	"github.com/minzhang2110/smart-home/pkg/devices"
	"github.com/minzhang2110/smart-home/pkg/devices/outlet"
	"github.com/minzhang2110/smart-home/pkg/oauth"
	"io/ioutil"
	"net/http"

	"github.com/minzhang2110/smart-home/pkg/intents"
)

// Handler type
type Handler struct {
	dvcs *devices.Mgr
}

// NewHandler constructor
func NewHandler() *Handler {
	o := outlet.New("1001", "light")
	return &Handler{
		dvcs: devices.New(o),
	}
}

// SmartHomeHandler .
func (h *Handler) SmartHomeHandler(w http.ResponseWriter, r *http.Request) {
	v, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userName, _ := r.Context().Value(oauth.UserName).(string)
	h.dvcs.SetAgentUserID(userName)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	data := intents.Execute(v, h.dvcs)
	w.Write(data)
	log.Printf("[raw] req: %s, resp: %s", string(v), string(data))
}
