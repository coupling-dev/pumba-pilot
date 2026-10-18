package container

import (
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/urfave/cli"
)

// FilterFromApp builds the container filter straight from the root CLI flags,
// so callers no longer pass the names through.
func FilterFromApp(c *cli.Context) FilterFunc {
	names := cliflags.NewV1FromApp(c).Args()
	return func(cn *Container) bool {
		return len(names) == 0 || matchNames(names, cn.Name(), cn.ID())
	}
}
