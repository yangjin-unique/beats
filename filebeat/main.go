package main

import (
	"os"

	"github.com/elastic/beats/filebeat/cmd"
)

// The basic model of execution:
// - input: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, input uses the registrar information, on restart, to
// determine where in each file to restart a harvester.
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
