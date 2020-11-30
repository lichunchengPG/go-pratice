package mock

import "fmt"

type Retrivever struct {
	Contents string
}

func (r *Retrivever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

func (r *Retrivever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]

	return "ok"
}

func (r *Retrivever) Get(url string) string  {
	return r.Contents
}