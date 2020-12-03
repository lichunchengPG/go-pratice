package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)


var router = mux.NewRouter()

type ArticleFormData struct {
	Title, Body string
	URL *url.URL
	Errors map[string]string
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w,"<h1>Hello 欢迎欢迎</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"此博客是用以记录编程笔记，如您有反馈或建议，请联系" +
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID: " + id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	//if err := r.ParseForm(); err != nil {
	//	fmt.Fprint(w, "请提供正确的数据")
	//	return
	//}

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := make(map[string]string)

	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if title == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}


	//fmt.Fprintf(w, "POST PostForm: %v <br>", r.PostForm)
	//fmt.Fprintf(w, "POST Form: %v <br>", r.Form)
	//fmt.Fprintf(w, "title 的值为: %v", title)

	//fmt.Fprintf(w, "r.Form 中 title 的值为: %v <br>", r.FormValue("title"))
	//fmt.Fprintf(w, "r.PostForm 中 title 的值为: %v <br>", r.PostFormValue("title"))
	//fmt.Fprintf(w, "r.Form 中 test 的值为: %v <br>", r.FormValue("test"))
	//fmt.Fprintf(w, "r.PostForm 中 test 的值为: %v <br>", r.PostFormValue("test"))

	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过!<br>")
		fmt.Fprintf(w, "title 的值为: %v <br>", title)
		fmt.Fprintf(w, "title 的长度为: %v <br>", utf8.RuneCountInString(title))
		fmt.Fprintf(w, "body 的值为: %v <br>", body)
		fmt.Fprintf(w, "body 的长度为: %v <br>", utf8.RuneCountInString(body))
	} else {
//		html := `
//<!DOCTYPE html>
//<html lang="en">
//<head>
//    <title>创建文章 —— 我的技术博客</title>
//    <style type="text/css">.error {color: red;}</style>
//</head>
//<body>
//    <form action="{{ .URL }}" method="post">
//        <p><input type="text" name="title" value="{{ .Title }}"></p>
//        {{ with .Errors.title }}
//        <p class="error">{{ . }}</p>
//        {{ end }}
//        <p><textarea name="body" cols="30" rows="10">{{ .Body }}</textarea></p>
//        {{ with .Errors.body }}
//        <p class="error">{{ . }}</p>
//        {{ end }}
//        <p><button type="submit">提交</button></p>
//    </form>
//</body>
//</html>
//`
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticleFormData {
			Title: title,
			Body: body,
			URL: storeURL,
			Errors: errors,
		}
		//tmpl, err := template.New("create-form").Parse(html)
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w,data)
	}

}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	storeURL, _ := router.Get("articles.store").URL()
	data := ArticleFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}

func forceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1、设置表头
		w.Header().Set("Content-Type","text/html; charset=utf-8")
		// 2、继续处理请求
		h.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router.HandleFunc("/", defaultHandler).Name("home")
	router.HandleFunc("/about", aboutHandler).Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	// 自定义路由页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	//  通过命名路由获取 URL 实例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)

	articleURL, _ := router.Get("articles.show").URL("id", "23")

	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}