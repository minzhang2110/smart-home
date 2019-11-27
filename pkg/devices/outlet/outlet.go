package outlet

import (
	"github.com/minzhang2110/smart-home/pkg/actions/sync"
)

// Outlet .
type Outlet struct {
	ID     string
	Name   string
	Type   string
	Traits []string
	On     bool
	Online bool
}

// New .
func New(id, name string) *Outlet {
	return &Outlet{
		ID:     id,
		Name:   name,
		Type:   sync.TypeOutlet,
		Traits: []string{sync.TraitOnOff},
		On:     false,
		Online: true,
	}
}

// Turn turn on or off
func (o *Outlet) Turn(on bool) error {
	if o.On != on {
		o.On = on
	}
	return nil
}
