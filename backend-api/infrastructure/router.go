package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tetsuzawa/vue-go-pay-tutorial/backend-api/handler"
	"log"
	"net/http"
)

var E *echo.Echo

/*
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

	E = e

}

*/

func NewRouter() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())

	//e.Use(middleware.CORS())

	// e.DefaultHTTPErrorHandler(controllers.APIError)//TODO
	e.GET("/", func(c echo.Context) error {
		e.Logger.Printf("root\n")
		err := c.JSON(http.StatusOK, "hellollo")
		if err != nil {
			e.Logger.Print(err)
			log.Println(err)
			return err
		}
		return nil
	})

	e.GET("/api/v1/items", func(c echo.Context) error {
		e.Logger.Printf("items\n")
		log.Println("getlists")
		err := handler.GetLists(c)
		if err != nil {
			e.Logger.Print(err)
			return err
		}
		return nil
	})
	e.GET("/api/v1/items/:id", func(c echo.Context) error {
		e.Logger.Printf("id\n")
		log.Println("getitem")
		err := handler.GetItem(c)
		if err != nil {
			e.Logger.Print(err)
			return err
		}
		return nil
	})
	e.POST("/api/v1/charge/items/:id", func(c echo.Context) error {
		e.Logger.Printf("charge\n")
		handler.Charge(c)
		return nil
	})

	//e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%s", os.Getenv("API_SERVER_PORT"))))

	return e
}
