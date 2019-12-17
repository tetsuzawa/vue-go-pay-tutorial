package infrastructure

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/handler"
)

var E *echo.Echo

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	e := echo.New()

	// e.DefaultHTTPErrorHandler(controllers.APIError)//TODO

	e.GET("/app/v1/items", func(c echo.Context) error {
		handler.GetLists(c)
		return nil
	})
	e.POST("/app/v1/items/:id", func(c echo.Context) error {
		handler.GetItem(c)
		return nil
	})
	e.POST("/api/v1/charge/items/:id", func(c echo.Context) error {
		handler.Charge(c)
		return nil
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Cfg.Web.Port)))

}
