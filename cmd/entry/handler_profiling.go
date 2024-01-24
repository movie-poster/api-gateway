package entry

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
)

func HandlerProfiling(router *echo.Echo, port string) {
	router.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))
	http.ListenAndServe(":"+port, nil)
}
