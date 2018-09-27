package setup

import(
	//"log"
	//"vwa/util/database"
	"net/http"
	"github.com/vwa/helper/middleware"
	"github.com/vwa/util/render"
	"github.com/julienschmidt/httprouter"
)

type Self struct{}

func New() *Self {
	return &Self{}
}
func (self *Self) SetRouter(r *httprouter.Router) {
	/* register all router */

	mw := middleware.New() //implement middleware

	r.GET("/setup", mw.LoggingMiddleware(mw.CapturePanic(SetupHandler)))

}

func SetupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
