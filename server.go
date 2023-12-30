package main

import (
	"echo-templ-htmx-test/controllers"
	"echo-templ-htmx-test/controllers/clicked"
	"embed"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
)

//go:embed assets
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "assets")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()

	assetHandler := http.FileServer(getFileSystem())

	e.GET("/", controllers.Root)
	e.GET("/clicked", clicked.ClickedGet)
	e.POST("/clicked", clicked.ClickedPost)
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))
	e.Logger.Fatal(e.Start(":1323"))
}
