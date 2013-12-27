// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated as part of a build step and shouldn't be manually modified
package sublime

import (
	"fmt"
	"github.com/quarnster/util/text"
	"lime/3rdparty/libs/gopy/lib"
	"lime/backend"
)

var (
	_ = backend.View{}
	_ = text.Region{}
	_ = fmt.Errorf
)

func sublime_ActiveWindow() (py.Object, error) {
	ret0 := backend.GetEditor().ActiveWindow()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_Arch() (py.Object, error) {
	ret0 := backend.GetEditor().Arch()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_GetClipboard() (py.Object, error) {
	ret0 := backend.GetEditor().GetClipboard()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_LogCommands(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 bool
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(bool); !ok {
				return nil, fmt.Errorf("Expected type bool for backend.Editor.LogCommands() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	backend.GetEditor().LogCommands(arg1)
	return toPython(nil)
}

func sublime_LogInput(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 bool
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(bool); !ok {
				return nil, fmt.Errorf("Expected type bool for backend.Editor.LogInput() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	backend.GetEditor().LogInput(arg1)
	return toPython(nil)
}

func sublime_NewWindow() (py.Object, error) {
	ret0 := backend.GetEditor().NewWindow()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_PackagesPath() (py.Object, error) {
	ret0 := backend.GetEditor().PackagesPath()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_Platform() (py.Object, error) {
	ret0 := backend.GetEditor().Platform()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_RunCommand(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 string
		arg2 backend.Args
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for backend.Editor.RunCommand() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	arg2 = make(backend.Args)
	if v, err := tu.GetItem(1); err == nil {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(backend.Args); !ok {
				return nil, fmt.Errorf("Expected type backend.Args for backend.Editor.RunCommand() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	backend.GetEditor().RunCommand(arg1, arg2)
	return toPython(nil)
}

func sublime_SetClipboard(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 string
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for backend.Editor.SetClipboard() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	backend.GetEditor().SetClipboard(arg1)
	return toPython(nil)
}

func sublime_Settings() (py.Object, error) {
	ret0 := backend.GetEditor().Settings()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func sublime_Version() (py.Object, error) {
	ret0 := backend.GetEditor().Version()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

var sublime_methods = []py.Method{
	{Name: "register", Func: sublime_Register},
	{Name: "unregister", Func: sublime_Unregister},
	{Name: "error_message", Func: sublime_ErrorMessage},
	{Name: "message_dialog", Func: sublime_MessageDialog},
	{Name: "ok_cancel_dialog", Func: sublime_OkCancelDialog},
	{Name: "status_message", Func: sublime_StatusMessage},
	{Name: "active_window", Func: sublime_ActiveWindow},
	{Name: "arch", Func: sublime_Arch},
	{Name: "get_clipboard", Func: sublime_GetClipboard},
	{Name: "log_commands", Func: sublime_LogCommands},
	{Name: "log_input", Func: sublime_LogInput},
	{Name: "new_window", Func: sublime_NewWindow},
	{Name: "packages_path", Func: sublime_PackagesPath},
	{Name: "platform", Func: sublime_Platform},
	{Name: "run_command", Func: sublime_RunCommand},
	{Name: "set_clipboard", Func: sublime_SetClipboard},
	{Name: "settings", Func: sublime_Settings},
	{Name: "version", Func: sublime_Version},
}
