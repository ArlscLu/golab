package main

import (
	"bytes"
	"flag"
	"html/template"
	"os"
	"testing"
)

func TestHttpServer(t *testing.T) {

}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
