package templates

import (
	"html/template"
	"io/ioutil"

	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/log"
	myio "github.com/jsorrell/www.jacksorrell.com/utils/io"
)

var bufPool = myio.CreateByteBufPool(5000)

/* Template Group */

// TemplateGroup contains a group of templates that can include eachother.
type TemplateGroup struct {
	*templateGroup
	name   string
	inited bool
}

// NewTemplateGroup returns a new TemplateGroup
func NewTemplateGroup(name string, funcs template.FuncMap) *TemplateGroup {
	return &TemplateGroup{newTemplateGroup(funcs), name, false}
}

// Init initializes the template. Call this once when all templates have been added.
func (g *TemplateGroup) Init() {
	if g.inited {
		log.Warn("TemplateGroup: Multiple calls to init on " + g.name + ".")
		return
	}
	g.inited = true
	g.init()
}

/* Static Template */

// StaticTemplate defines a template that take no arguments to compile.
type StaticTemplate struct {
	*staticTemplate
	group *TemplateGroup
}

// NewStaticTemplate returns a new StaticTemplate
func (g *TemplateGroup) NewStaticTemplate(templateName, fileName string, argCreator func() interface{}) *StaticTemplate {
	return &StaticTemplate{g.newStaticTemplate(templateName, fileName, argCreator), g}
}

// GetReader returns a io.Reader for the template. Call close() when done reading.
// Also returns an etag. TODO remove this.
func (tmpl *StaticTemplate) GetReader() (myio.ReadSeekerCloser, string) {
	tmpl.group.checkInit()
	return tmpl.getReader()
}

/* Dynamic Template */

// DynamicTemplate defines a template that takes an argument to compile.
type DynamicTemplate struct {
	*dynamicTemplate
	group        *TemplateGroup
	templateName string
}

// NewDynamicTemplate returns a new DynamicTemplate
func (g *TemplateGroup) NewDynamicTemplate(templateName, fileName string) *DynamicTemplate {
	return &DynamicTemplate{g.newDynamicTemplate(templateName, fileName), g, templateName}
}

// GetReader returns an io.Reader for the template. Call close() when done reading.
func (tmpl *DynamicTemplate) GetReader(arg interface{}) myio.ReadSeekerCloser {
	tmpl.group.checkInit()
	buf := bufPool.Get()
	if err := tmpl.group.getTemplate().ExecuteTemplate(buf, tmpl.templateName, arg); err != nil {
		log.Panic(err)
	}
	return buf.GetReader()
}

/* Included Template */

// NewIncludedTemplate adds a template to the group that can be referenced by other templates.
func (g *TemplateGroup) NewIncludedTemplate(templateName, fileName string) {
	g.newIncludedTemplate(templateName, fileName)
}

/* Helpers */

func createArg(f func() interface{}) interface{} {
	if f == nil {
		return nil
	}
	return f()
}

func (g *TemplateGroup) checkInit() {
	if !g.inited {
		g.Init()
		log.Warn("Template group " + g.name + " not inited. Initing...")
	}
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

	d, err := ioutil.ReadAll(templateFile)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
