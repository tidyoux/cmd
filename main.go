package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidyoux/cmd/plugin"
)

var (
	port = flag.Uint("p", 6666, "service port")
)

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/cmd", run)
	if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
		panic(err)
	}
}

type Args struct {
	Args string `json:"args"`
}

func run(c *gin.Context) {
	var args Args
	if err := c.ShouldBindJSON(&args); err != nil {
		response(c, []string{"invalid args"})
		return
	}

	if data, err := exe(args.Args); err != nil {
		response(c, []string{err.Error()})
	} else {
		response(c, data)
	}
}

func exe(args string) ([]string, error) {
	cmds := strings.Split(strings.TrimSpace(args), "|")
	if len(cmds) == 0 {
		return nil, nil
	}

	var (
		result []string
		err    error
	)
	for _, s := range cmds {
		cmdinfo := trimSpace(strings.Split(s, " "))
		if len(cmdinfo) == 0 {
			continue
		}

		cmd := cmdinfo[0]
		p, ok := plugin.Find(cmd)
		if !ok {
			return nil, fmt.Errorf("invalid cmd: %s", cmd)
		}

		result, err = p.Exe(cmd, append(cmdinfo[1:], result...))
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func response(c *gin.Context, data []string) {
	c.String(http.StatusOK, strings.Join(data, "\n"))
}
