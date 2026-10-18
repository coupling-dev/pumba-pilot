package chaos

import (
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
)

// ListOptsFromFlags reads the container list options from the global command-line flags.
func ListOptsFromFlags(flags cliflags.Flags) container.ListOpts {
	return container.ListOpts{Labels: flags.Global().StringSlice("label")}
}
