package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	fmt.Print("server start\n")
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":9901", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	p, _ := makePage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	p, err := makePage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	_ = p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

//函数接收一个处理的方法，去处理某个问题
//函数接收一个处理的方法，对这个方法进行调整
//* 方法config/use
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

type Page struct {
	Title string
	Body  []byte
}

//使用
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//构造
func makePage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
