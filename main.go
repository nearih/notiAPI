package main

import (
	"fmt"
	"net/http"
	"time"

	logs "log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func initAPI() {
	pathPrefix := "/noti"
	fmt.Println("starting server")
	logs.Println("start api")
	port := viper.GetString("port")
	if port == "" {
		port = ":8005"
	}
	h := Handler{}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.GET(pathPrefix+"/test", h.Test)
	e.POST(pathPrefix+"/listen", h.Listen)
	e.POST("/", h.Listen)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(port))
}

func main() {
	initAPI()
}

type Handler struct{}

func (h *Handler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}

func (h *Handler) Listen(c echo.Context) error {

	t := time.Now()

	json_map := make(map[string]interface{})
	if err := c.Bind(&json_map); err != nil {
		c.Logger().Error(t, err)
		return err
	}
	c.Logger().Info(fmt.Sprintf("[Noti] %v , Body :  %+v", t.Format("2006-01-02T15:04:05"), json_map))
	return c.JSON(http.StatusOK, json_map)
}
