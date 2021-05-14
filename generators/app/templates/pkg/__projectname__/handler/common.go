package handler

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo"
)

type URLParam struct{}

func parseArg(c echo.Context) (*url.Values, error) {
	param := URLParam{}
	if err := c.Bind(&param); err != nil {
		return nil, fmt.Errorf("argument parse error")
	}

	v := url.Values{}

	return &v, nil
}
