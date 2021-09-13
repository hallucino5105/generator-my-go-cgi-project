package <%= projectNameSnakeCase %>

import (
	"fmt"
  "path"

	"github.com/hallucino5105/glog"
	"github.com/hallucino5105/<%= projectNameSnakeCase %>/cmd/garg"
	"github.com/hallucino5105/<%= projectNameSnakeCase %>/pkg/appconf"
	"github.com/hallucino5105/<%= projectNameSnakeCase %>/pkg/<%= projectNameSnakeCase %>/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func Entry() error {
	glog.SetupLogger(glog.WithVerbose(garg.GlobalOptions.Verbose))

	conf, err := appconf.LoadAppConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	startServer(conf)

	return nil
}

func startServer(conf *appconf.AppConfig) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

  e.GET(path.Join(conf.Serve.PublicPath, "/api/sample.json"), handler.HandlerSample())

	addr := fmt.Sprintf("%s:%s", conf.Serve.Host, conf.Serve.Port)
	glog.Debugf("serve address \"%s\"", addr)

	e.Logger.Fatal(e.Start(addr))
}

