// +build dev

package templates

import (
	"html/template"

	"github.com/jsorrell/www.jacksorrell.com/log"
	myio "github.com/jsorrell/www.jacksorrell.com/utils/io"
	"github.com/jsorrell/www.jacksorrell.com/web/pages"
)

// In Development, we don't care much about time.
// We must re-read the templates every time they are requested.

/* Template Definitions */

type _templateDef struct {
	templateName string
	fileName     string
}

type _templateDefs []_templateDef

func (d *_templateDefs) _addTemplateDef(templateName, fileName string) {
	*d = append(*d, _templateDef{templateName, fileName})
}

/* Template Group */

type templateGroup struct {
	funcs template.FuncMap
	defs  _templateDefs
}

func newTemplateGroup(funcs template.FuncMap) *templateGroup {
	return &templateGroup{funcs, _templateDefs{}}
}

func (g *TemplateGroup) getTemplate() *template.Template {
	tmpl := template.New("").Funcs(g.funcs)
	for _, def := range g.defs {
		if err := addTemplate(&tmpl, def.templateName, def.fileName); err != nil {
			log.Panic(err)
		}
	}
	return tmpl
}

func (g *TemplateGroup) init() {}

/* Static Template */

type staticTemplate struct {
	templateName string
	argCreator   func() interface{}
}

func (g *TemplateGroup) newStaticTemplate(templateName, fileName string, createArgs func() interface{}) *staticTemplate {
	g.defs._addTemplateDef(templateName, fileName)
	return &staticTemplate{templateName, createArgs}
}

func (tmpl *StaticTemplate) getReader() (myio.ReadSeekerCloser, string) {
	buf := bufPool.Get()
	t := tmpl.group.getTemplate()
	if err := t.ExecuteTemplate(buf, tmpl.templateName, createArg(tmpl.argCreator)); err != nil {
		buf.Close()
		log.Panic(err)
	}
	return buf.GetReader(), pages.GenEtag(buf.Bytes())
}

/* Dynamic Template */

type dynamicTemplate struct{}

func (g *TemplateGroup) newDynamicTemplate(templateName, fileName string) *dynamicTemplate {
	g.defs._addTemplateDef(templateName, fileName)
	return &dynamicTemplate{}
}

/* Included Template */

func (g *TemplateGroup) newIncludedTemplate(templateName, fileName string) {
	g.defs._addTemplateDef(templateName, fileName)
}
