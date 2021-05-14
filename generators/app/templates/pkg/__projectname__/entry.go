package <%= project_name %>

import (
	"fmt"

	"github.com/hallucino5105/glog"
	"github.com/hallucino5105/<%= project_name %>/cmd/garg"
	"github.com/hallucino5105/<%= project_name %>/pkg/<%= project_name %>/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func Entry() error {
	glog.SetupLogger(glog.WithVerbose(garg.GlobalOptions.Verbose))

	conf, err := LoadAppConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	startServer(conf)

	return nil
}

func startServer(conf *AppConfig) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/sample.json", handler.HandlerSample())

	addr := fmt.Sprintf("%s:%s", conf.Serve.Host, conf.Serve.Port)
	glog.Debugf("serve address \"%s\"", addr)

	e.Logger.Fatal(e.Start(addr))
}

