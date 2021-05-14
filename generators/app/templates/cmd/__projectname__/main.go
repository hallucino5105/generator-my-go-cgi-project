package main

import (
	"github.com/hallucino5105/glog"
	"github.com/hallucino5105/my_go_cgi_project/cmd/garg"
	"github.com/hallucino5105/my_go_cgi_project/pkg/<%= project_name %>"
)

func init() {
	garg.ParseArg()
	glog.SetupLogger(glog.WithVerbose(garg.GlobalOptions.Verbose))
}

func main() {
	glog.Debug("start <%= project_name %>")

	if err := <%= project_name %>.Entry(); err != nil {
		glog.Fatal(err)
	}
}
