package controllers

import (
	"database/sql"
	"fmt"
	"github.com/lichunchengPG/go-pratice/goblog/pkg/logger"
	"github.com/lichunchengPG/go-pratice/goblog/pkg/route"
	"github.com/lichunchengPG/go-pratice/goblog/pkg/types"
	"html/template"
	"net/http"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariables("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示文章
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"Name2URL" : route.Name2URL,
				"Int64ToString" : types.Int64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		tmpl.Execute(w, article)
	}
}