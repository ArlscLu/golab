package gosyntax

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func Callback(a int, f func(int, int) int) int {
	return f(a, 10)
}

func Add(a, b int) int {
	return a + b
}

func R() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	// fmt.Println(`can see the seasion`,season)
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}

func HtmlParse() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find link%v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
