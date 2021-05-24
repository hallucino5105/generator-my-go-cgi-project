package handler

import (
	"net/http"

	"github.com/hallucino5105/<%= projectNameSnakeCase %>/pkg/urlarg"
	"github.com/labstack/echo"
)

func HandlerSample() echo.HandlerFunc {
	return func(c echo.Context) error {
		arg, err := urlarg.ParseArg(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		} else if arg == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "invalid arguments")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"res": "api sample",
		})
	}
}
