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
		buf := bytes.NewBuffer(make([]byte, 0, 5000)) // This stores data, so don't reuse this buf
		var args interface{}
		if static.createArgs != nil {
			args = static.createArgs()
		}
		if err := g.t.ExecuteTemplate(buf, static.templateName, args); err != nil {
			log.Fatal(err)
		}
		static.tmpl.compiledTemplate = buf.Bytes()
		static.tmpl.genEtag()
	}
}

type DynamicTemplate struct {
	group        *TemplateGroup
	templateName string
	contentType  string
}

func (g *TemplateGroup) NewDynamicTemplate(templateName, fileName string) *DynamicTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	return &DynamicTemplate{g, templateName, DefaultContentType}
}

func (tmpl *DynamicTemplate) GetReader(args interface{}) ReadSeekerCloser {
	tmpl.group.checkInit()
	buf := bufPool.Get().(pooledBuffer)
	if err := tmpl.group.t.ExecuteTemplate(buf, tmpl.templateName, args); err != nil {
		log.Panic(err)
	}
	return buf.getReader()
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
	etag             string
	contentType      string
}

func (g *TemplateGroup) NewStaticTemplate(templateName, fileName string, createArgs func() interface{}) *StaticTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	tmpl := &StaticTemplate{g, []byte{}, "", DefaultContentType}
	g.statics = append(g.statics, _prod_staticTemplateInfo{tmpl, templateName, createArgs})
	return tmpl
}

func (tmpl *StaticTemplate) GetReader() (ReadSeekerCloser, string) {
	tmpl.group.checkInit()
	return ByteReadSeekerCloser{bytes.NewReader(tmpl.compiledTemplate)}, tmpl.etag
}

func (tmpl *StaticTemplate) Execute(w io.Writer) {
	tmpl.group.checkInit()
	if _, err := w.Write(tmpl.compiledTemplate); err != nil {
		log.Panic(err)
	}
}

func (tmpl *StaticTemplate) GetEtag() string {
	return tmpl.etag
}

func (tmpl *StaticTemplate) genEtag() {
	tmpl.etag = genEtag(tmpl.compiledTemplate)
}
