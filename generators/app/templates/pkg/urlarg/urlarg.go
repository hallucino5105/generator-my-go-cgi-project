package urlarg

import (
	"fmt"

	"github.com/labstack/echo"
)

type URLParam struct{}

func ParseArg(c echo.Context) (*URLParam, error) {
	param := URLParam{}
	if err := c.Bind(&param); err != nil {
		return nil, fmt.Errorf("argument parse error")
	}
	return &param, nil
}
