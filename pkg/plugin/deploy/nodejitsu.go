package deploy

import (
	"github.com/drone/drone/pkg/build/buildfile"
	"github.com/drone/drone/pkg/plugin/condition"
)

type Nodejitsu struct {
	App   string `yaml:"app,omitempty"`
	User  string `yaml:"user,omitempty"`
	Token string `yaml:"token,omitempty"`

	Condition *condition.Condition `yaml:"when,omitempty"`
}

func (n *Nodejitsu) Write(f *buildfile.Buildfile) {
	f.WriteEnv("username", n.User)
	f.WriteEnv("apiToken", n.Token)

	// Install the jitsu command line interface then
	// deploy the configured app.
	f.WriteCmdSilent("[ -f /usr/bin/sudo ] || npm install -g jitsu")
	f.WriteCmdSilent("[ -f /usr/bin/sudo ] && sudo npm install -g jitsu")
	f.WriteCmd("jitsu deploy")
}

func (n *Nodejitsu) GetCondition() *condition.Condition {
	return n.Condition
}
