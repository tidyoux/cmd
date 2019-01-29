package plugin

import "sort"

func init() {
	Register("help", &Help{})
}

type Help struct{}

func (*Help) Exe(cmd string, args []string) ([]string, error) {
	result := make([]string, 0, len(plugins)+2)
	result = append(result, `example: echo this is a test | md5 | base64`)
	result = append(result, "valid cmd:")
	for name := range plugins {
		result = append(result, " "+name)
	}
	sort.Strings(result[2:])
	return result, nil
}
