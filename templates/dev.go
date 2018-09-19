// +build dev

package templates

import (
	"html/template"

	"github.com/jsorrell/www.jacksorrell.com/log"
	myio "github.com/jsorrell/www.jacksorrell.com/utils/io"
	"github.com/jsorrell/www.jacksorrell.com/web/pages"

	"io"
)

type _dev_templateDef struct {
	templateName string
	fileName     string
}

type _dev_templateDefs []_dev_templateDef

type TemplateGroup struct {
	name   string
	funcs  template.FuncMap
	defs   _dev_templateDefs
	inited bool
}

func NewTemplateGroup(name string, funcs template.FuncMap) *TemplateGroup {
	return &TemplateGroup{name, funcs, _dev_templateDefs{}, false}
}

func (d *_dev_templateDefs) _dev_addTemplateDef(templateName, fileName string) {
	*d = append(*d, _dev_templateDef{templateName, fileName})
}

func (g *TemplateGroup) _dev_parseAllTemplates() (*template.Template, error) {
	tmpl := template.New("").Funcs(g.funcs)
	for _, def := range g.defs {
		if err := addTemplate(&tmpl, def.templateName, def.fileName); err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}

func (g *TemplateGroup) NewIncludedTemplate(templateName, fileName string) {
	g.defs._dev_addTemplateDef(templateName, fileName)
}

func (g *TemplateGroup) Init() {
	g.inited = true
}

type DynamicTemplate struct {
	group        *TemplateGroup
	templateName string
}

func (g *TemplateGroup) NewDynamicTemplate(templateName, fileName string) *DynamicTemplate {
	g.defs._dev_addTemplateDef(templateName, fileName)
	return &DynamicTemplate{g, templateName}
}

func (tmpl *DynamicTemplate) GetReader(args interface{}) myio.ReadSeekerCloser {
	tmpl.group.checkInit()
	buf := bufPool.Get()
	err := tmpl._dev_execute(buf, args)
	if err != nil {
		buf.Close()
		log.Panic(err)
	}
	return buf.GetReader()
}

func (tmpl *DynamicTemplate) Execute(w io.Writer, args interface{}) {
	tmpl.group.checkInit()
	err := tmpl._dev_execute(w, args)
	if err != nil {
		log.Panic(err)
	}
}

func (tmpl *DynamicTemplate) _dev_execute(w io.Writer, args interface{}) error {
	t, err := tmpl.group._dev_parseAllTemplates()
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, tmpl.templateName, args)
}

type StaticTemplate struct {
	group        *TemplateGroup
	templateName string
	createArgs   func() interface{}
}

func (g *TemplateGroup) NewStaticTemplate(templateName, fileName string, createArgs func() interface{}) *StaticTemplate {
	g.defs._dev_addTemplateDef(templateName, fileName)
	return &StaticTemplate{g, templateName, createArgs}
}

func (tmpl *StaticTemplate) GetReader() (myio.ReadSeekerCloser, string) {
	tmpl.group.checkInit()

	buf := bufPool.Get()
	err := tmpl._dev_execute(buf)
	if err != nil {
		buf.Close()
		log.Panic(err)
	}
	return buf.GetReader(), pages.GenEtag(buf.Bytes())
}

func (tmpl *StaticTemplate) Execute(w io.Writer) {
	tmpl.group.checkInit()
	err := tmpl._dev_execute(w)
	if err != nil {
		log.Panic(err)
	}
}

func (tmpl *StaticTemplate) _dev_execute(w io.Writer) error {
	t, err := tmpl.group._dev_parseAllTemplates()
	if err != nil {
		return err
	}
	var args interface{}
	if tmpl.createArgs != nil {
		args = tmpl.createArgs()
	}
	return t.ExecuteTemplate(w, tmpl.templateName, args)
}
