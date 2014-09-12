package deploy

import (
	"github.com/drone/drone/pkg/build/buildfile"
	"github.com/drone/drone/pkg/plugin/condition"
)

type Bash struct {
	Script  []string `yaml:"script,omitempty"`
	Command string   `yaml:"command,omitempty"`

	Condition *condition.Condition `yaml:"when,omitempty"`
}

func (g *Bash) Write(f *buildfile.Buildfile) {
	g.Script = append(g.Script, g.Command)

	for _, cmd := range g.Script {
		f.WriteCmd(cmd)
	}
}

func (g *Bash) GetCondition() *condition.Condition {
	return g.Condition
}
