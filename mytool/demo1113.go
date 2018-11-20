package mytool

import (
	"net/http"

	"github.com/labstack/echo"
)

// Main1113 echo test
func Main1113() {
	e := echo.New()
	e.GET("/", baseHello)

	e.Logger.Fatal(e.Start(":8887"))
}

func baseHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Allen!!!")
}
