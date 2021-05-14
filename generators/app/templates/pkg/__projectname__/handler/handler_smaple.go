package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandlerSample() echo.HandlerFunc {
	return func(c echo.Context) error {
		arg, err := parseArg(c)
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
