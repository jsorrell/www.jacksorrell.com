package templates

import (
	"html/template"

	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/utils/io"
)

var bufPool = io.CreateByteBufPool(5000)

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
	return data.ReadFileToString(templateFile)
}
