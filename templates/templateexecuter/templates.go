package templateexecuter

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"sync"

	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/log"
)

var bufPool *sync.Pool

func init() {
	bufPool = &sync.Pool{
		New: func() interface{} {
			return ReadCloser{new(bytes.Buffer)}
		},
	}
}

func (g *TemplateGroup) checkInit() {
	if !g.inited {
		g.Init()
		log.Warn("Template group " + g.name + " not inited. Initing...")
	}
}

type ReadCloser struct {
	*bytes.Buffer
}

func (r ReadCloser) Close() error {
	r.Reset()
	bufPool.Put(r)
	return nil
}

func addTemplate(tmpl **template.Template, templateName string, templateFile string) error {
	templateString, err := readTemplateString(templateFile)
	if err != nil {
		return err
	}
	*tmpl, err = (*tmpl).New(templateName).Parse(templateString)
	return err
}

func readTemplateString(filename string) (string, error) {
	templateFile, err := data.Templates.Open(filename)
	if err != nil {
		return "", err
	}
	defer templateFile.Close()
	return data.ReadFileToString(templateFile)
}

func serveHTTP(w http.ResponseWriter, req *http.Request, r io.Reader) {
	if req.Method != "HEAD" {
		_, err := io.Copy(w, r)
		if err != nil {
			log.Info(err)
		}
	} else {
		w.WriteHeader(200)
	}
}

func (tmpl *DynamicTemplate) ServeHTTP(w http.ResponseWriter, req *http.Request, args interface{}) {
	r := tmpl.GetReadCloser(args)
	defer r.Close()
	serveHTTP(w, req, r)
}

func (tmpl *StaticTemplate) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := tmpl.GetReadCloser()
	defer r.Close()
	serveHTTP(w, req, r)
}
