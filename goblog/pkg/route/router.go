package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router路由对象
var Router *mux.Router


// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	}

	return url.String()
}


func GetRouteVariables(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}