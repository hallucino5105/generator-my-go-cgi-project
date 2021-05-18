package main

import (
	"github.com/hallucino5105/glog"
	"github.com/hallucino5105/<%= projectNameSnakeCase %>/cmd/garg"
	"github.com/hallucino5105/<%= projectNameSnakeCase %>/pkg/<%= projectNameSnakeCase %>"
)

func init() {
	garg.ParseArg()
	glog.SetupLogger(glog.WithVerbose(garg.GlobalOptions.Verbose))
}

func main() {
	glog.Debug("start <%= projectNameSnakeCase %>")

	if err := <%= projectNameSnakeCase %>.Entry(); err != nil {
		glog.Fatal(err)
	}
}
