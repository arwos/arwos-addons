package addons

import (
	"fmt"
	"plugin"
)

const SetupFuncName = "ARWOS"

func Load(path string) (setup Setup, plg *plugin.Plugin, err error) {
	defer func() {
		if e := recover(); e != nil {
			setup, plg, err = nil, nil, fmt.Errorf("fail load addon: %+v", e)
		}
	}()

	plg, err = plugin.Open(path)
	if err != nil {
		setup, plg = nil, nil
		return
	}

	var s plugin.Symbol
	s, err = plg.Lookup(SetupFuncName)
	if err != nil {
		setup, plg = nil, nil
		return
	}

	fn, ok := s.(func() Setup)
	if !ok {
		setup, plg, err = nil, nil, fmt.Errorf("func `%s` not found", SetupFuncName)
		return
	}

	setup = fn()

	return
}
