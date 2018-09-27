package main

import (
	
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vwa/util"
	"github.com/vwa/util/render"
	"github.com/vwa/helper/middleware"
	"github.com/vwa/modules/user"
	"github.com/vwa/modules/user/profile"
	"github.com/vwa/modules/product/komentar"

)


func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
	data := make(map[string]interface{})
	data["title"] = "Index"

	render.HTMLRender(w,r, "template.index", data)
}

func main(){
	mw := middleware.New()
	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.GET("/", mw.LoggingMiddleware(indexHandler))
	router.GET("/index", mw.LoggingMiddleware(indexHandler))

	user := user.New()
	user.SetRouter(router)
	
	komentar := komentar.New()
	komentar.SetRouter(router)

	profile := profile.New()
	profile.SetRouter(router)

	s := http.Server{
		Addr : ":8082",
		Handler : router,
	}

	fmt.Printf("Server running at port %s\n", s.Addr)
	fmt.Printf("Open this url %s on your browser to access VWA",util.Fullurl)
	fmt.Println("")
	s.ListenAndServe()
}