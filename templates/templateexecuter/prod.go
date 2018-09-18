// +build !dev

package templateexecuter

import (
	"bytes"
	"html/template"

	"github.com/jsorrell/www.jacksorrell.com/log"

	"io"
)

type _prod_staticTemplateInfo struct {
	tmpl         *StaticTemplate
	templateName string
	createArgs   func() interface{}
}

type TemplateGroup struct {
	name    string
	t       *template.Template
	statics []_prod_staticTemplateInfo
	inited  bool
}

func NewTemplateGroup(name string, funcs template.FuncMap) *TemplateGroup {
	return &TemplateGroup{name, template.New("").Funcs(funcs), make([]_prod_staticTemplateInfo, 0), false}
}

func (g *TemplateGroup) NewIncludedTemplate(templateName, fileName string) {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
}

func (g *TemplateGroup) Init() {
	if g.inited {
		log.Warn("TemplateGroup: Multiple calls to init on " + g.name + ".")
		return
	}
	g.inited = true
	for _, static := range g.statics {
		buf := bytes.NewBuffer(make([]byte, 0, 5000))
		var args interface{}
		if static.createArgs != nil {
			args = static.createArgs()
		}
		if err := g.t.ExecuteTemplate(buf, static.templateName, args); err != nil {
			log.Fatal(err)
		}
		static.tmpl.compiledTemplate = buf.Bytes()
	}
}

type DynamicTemplate struct {
	group        *TemplateGroup
	templateName string
}

func (g *TemplateGroup) NewDynamicTemplate(templateName, fileName string) *DynamicTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	return &DynamicTemplate{g, templateName}
}

func (tmpl *DynamicTemplate) GetReadCloser(args interface{}) io.ReadCloser {
	tmpl.group.checkInit()
	buf := bufPool.Get().(ReadCloser)
	if err := tmpl.group.t.ExecuteTemplate(buf, tmpl.templateName, args); err != nil {
		log.Panic(err)
	}
	return buf
}

func (tmpl *DynamicTemplate) Execute(w io.Writer, args interface{}) {
	tmpl.group.checkInit()
	if err := tmpl.group.t.ExecuteTemplate(w, tmpl.templateName, args); err != nil {
		log.Panic(err)
	}
}

type StaticTemplate struct {
	group            *TemplateGroup
	compiledTemplate []byte
}

func (g *TemplateGroup) NewStaticTemplate(templateName, fileName string, createArgs func() interface{}) *StaticTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	tmpl := &StaticTemplate{g, []byte{}}
	g.statics = append(g.statics, _prod_staticTemplateInfo{tmpl, templateName, createArgs})
	return tmpl
}

type ByteReadCloser struct {
	io.Reader
}

func (ByteReadCloser) Close() error {
	return nil
}

func (tmpl *StaticTemplate) GetReadCloser() io.ReadCloser {
	tmpl.group.checkInit()
	return ByteReadCloser{bytes.NewReader(tmpl.compiledTemplate)}
}

func (tmpl *StaticTemplate) Execute(w io.Writer) {
	tmpl.group.checkInit()
	if _, err := w.Write(tmpl.compiledTemplate); err != nil {
		log.Panic(err)
	}
}
