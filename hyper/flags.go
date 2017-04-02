package hyper

import (
	"sort"

	"github.com/hyperhq/hypercli/cli"
)

type byName []cli.Command

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

var dockerCommands []cli.Command

func init() {
	for _, cmd := range cli.DockerCommands {
		dockerCommands = append(dockerCommands, cmd)
	}
	sort.Sort(byName(dockerCommands))
}
