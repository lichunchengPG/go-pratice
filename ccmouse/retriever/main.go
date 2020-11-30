package main

import (
	"fmt"
	"lichuncheng/retriever/mock"
	"lichuncheng/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents" : "another faked imooc.com",
	})

	return s.Get(url)
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf("> Type:%T Value:%v\n", r, r)
	fmt.Print("> Type Switch:")
	switch v:= r.(type) {
	case *mock.Retrivever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}

	fmt.Println()
}
func main() {
	var r Retriever
	mockRetriever := mock.Retrivever{Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	inspect(r)

	r = &real.Retriever{UserAgent: "Mozilla/5.0", Timeout: time.Minute}

	inspect(r)

	// Type Assertion
	if mockRetriever, ok := r.(*mock.Retrivever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}


	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}