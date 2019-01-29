package plugin

func init() {
	Register("echo", &Echo{})
}

type Echo struct{}

func (*Echo) Exe(cmd string, args []string) ([]string, error) {
	return args, nil
}
