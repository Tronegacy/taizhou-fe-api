package router

import (
	"anla.io/taizhou-fe-api/app/jwt"
	"anla.io/taizhou-fe-api/handler"
	"github.com/kataras/iris"
)

// FileRouter is
func FileRouter(party iris.Party) {
	o := party.Party("/upload")
	{
	}
	a := o.Party("/a")
	a.Use(jwt.JwtHandler.Serve)
	{
		a.Post("/", handler.UploadFile)
	}

}
