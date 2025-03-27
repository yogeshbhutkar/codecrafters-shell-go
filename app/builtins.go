package main

var supported_builtins = []string{
	"exit",
	"echo",
	"pwd",
	"type",
}

var BuiltinsMap = func() map[string]bool {
	var _builtin_map = make(map[string]bool)
	for _, builtin := range supported_builtins {
		_builtin_map[builtin] = true
	}
	return _builtin_map
}()
