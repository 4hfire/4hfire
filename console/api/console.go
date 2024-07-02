package main

import (
	"4hfire/api/internal/handler"
	"4hfire/resource"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"io/fs"
	"net/http"
	"os"

	"4hfire/api/internal/config"
	"4hfire/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/console-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	webHandler := WrapHandler(http.FileServer(http.FS(getFS(c.RestConf.Mode == service.DevMode))))

	// 前端服务
	server.AddRoutes(
		[]rest.Route{
			{
				// 前端页面
				Method:  http.MethodGet,
				Path:    "/",
				Handler: webHandler,
			},
			{
				// 静态资源
				Method:  http.MethodGet,
				Path:    "/assets/:filepath",
				Handler: webHandler,
			},
			{
				// logo
				Method:  http.MethodGet,
				Path:    "/favicon.ico",
				Handler: webHandler,
			},
		})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}

func getFS(useOS bool) fs.FS {
	if useOS {
		return os.DirFS("../web/dist")
	}

	fsys, err := fs.Sub(resource.Resource, "dist")
	if err != nil {
		panic(err)
	}

	return fsys
}

func WrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", `public, max-age=31536000`)
		fmt.Println(r.URL)
		h.ServeHTTP(w, r)
	}
}
