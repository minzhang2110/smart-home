package controllers

import (
	"fmt"
	"github.com/minzhang2110/smart-home/pkg/devices"
	"github.com/minzhang2110/smart-home/pkg/devices/outlet"
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
	fmt.Fprint(w, intents.Execute(v, h.dvcs))
}
