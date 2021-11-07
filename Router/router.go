package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", ServerHome)
	return muxRouter
}

func ServerHome(w http.ResponseWriter, r *http.Request) {

}
