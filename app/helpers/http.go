package helpers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HTTPNotFound 404 code helper
func HTTPNotFound(c echo.Context) error {
	return c.JSONBlob(http.StatusNotFound, []byte(`{"message": "Not Found"}`))
}

// InternalServerError 500 code helper
func InternalServerError(c echo.Context, err error) error {
	c.Logger().Error(err)
	return c.JSON(http.StatusInternalServerError, err)
}
