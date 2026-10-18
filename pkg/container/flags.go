package container

import "github.com/alexei-led/pumba/pkg/chaos/cliflags"

// ListOptsFromFlags reads the container list options from the global command-line flags.
func ListOptsFromFlags(flags cliflags.Flags) ListOpts {
	return ListOpts{Labels: flags.Global().StringSlice("label")}
}
