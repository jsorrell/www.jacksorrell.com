package templates

import (
	"html/template"
	"strings"

	"github.com/jsorrell/www.jacksorrell.com/configloader"
	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/templates/templatemap"
)

type templateDef struct {
	title    string
	filename string
}

var templates = []templateDef{
	/* Contact */
	templateDef{"contactBox", "contact/contact-box.gohtml"},
	templateDef{"contactPage", "contact/contact-page.gohtml"},
	templateDef{"floatingContact", "contact/floating-contact.gohtml"},
	/* Resume */
	templateDef{"resume", "resume/resume.gohtml"},
	/* Error */
	templateDef{"error", "error/error.gohtml"},
	/* Includes */
	templateDef{"favicon", "includes/favicon.gohtml"},
}

var templateMap = templatemap.Create()

func init() {
	templateMap.SetTemplate("www", func() (*template.Template, error) {
		tmpl := template.New("").Funcs(template.FuncMap{
			"contactMaxLength": configloader.ContactMaxLength,
			"resumeGenerateStars": func(stars int) string {
				if stars < 0 {
					stars = 0
				}
				if stars > 5 {
					stars = 5
				}
				return strings.Repeat("★", stars) + strings.Repeat("☆", 5-stars)
			},
		})

		/* Contact */
		for _, t := range templates {
			if err := addTemplate(&tmpl, t.title, t.filename); err != nil {
				return nil, err
			}
		}
		return tmpl, nil
	})
}

// GetTemplate retrieves the template defined by the string id
func GetTemplate(id string) (*template.Template, error) {
	return templateMap.GetTemplate(id)
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
