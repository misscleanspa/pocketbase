package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	/* app.OnBeforeServe().Add(func(e *core.ServeEvent) error { */
	/* 	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false)) */
	/* 	e.Router.GET("/hello", func(c echo.Context) error { */
	/* 		return c.String(http.StatusOK, "Hello world!") */
	/* 	}, apis.ActivityLogger(app)) */
	/* 	return nil */
	/* }) */

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello world!")
		}, apis.ActivityLogger(app))

		return nil
	})

	app.OnRecordAuthRequest().Add(func(e *core.RecordAuthEvent) error {
		log.Println(e.HttpContext)
		log.Println(e.Record)
		log.Println(e.Token)
		log.Println(e.Meta)
		return nil
	})

	// fires only for "users" and "managers" auth collections
	app.OnRecordAuthRequest("users", "managers").Add(func(e *core.RecordAuthEvent) error {
		log.Println(e.HttpContext)
		log.Println(e.Record)
		log.Println(e.Token)
		log.Println(e.Meta)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
