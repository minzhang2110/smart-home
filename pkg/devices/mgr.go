package devices

import (
	"github.com/minzhang2110/smart-home/pkg/devices/outlet"
)

type Mgr struct {
	Outlet *outlet.Outlet
}

func New(o *outlet.Outlet) *Mgr {
	return &Mgr{
		Outlet: o,
	}
}
