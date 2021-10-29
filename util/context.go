package util

import "github.com/labstack/echo"

type TypeErrorResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(c echo.Context, code int, error string) error {
	return c.JSON(code, TypeErrorResponse{
		Error: error,
	})
}
