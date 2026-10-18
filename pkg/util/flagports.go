package util

import (
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/urfave/cli"
)

// PortsFromFlags reads and validates the --port flag of the current command.
func PortsFromFlags(c *cli.Context) ([]string, error) {
	return GetPorts(cliflags.NewV1(c).String("port"))
}
