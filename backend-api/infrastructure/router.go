package infrastructure

import (
	"fmt"
	"github.com/labstack/echo/middleware"
	"log"
	"os"

	"github.com/labstack/echo"

	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/handler"
)

var E *echo.Echo

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	e := echo.New()

	e.Use(middleware.CORS())

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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("API_SERVER_PORT"))))

}
