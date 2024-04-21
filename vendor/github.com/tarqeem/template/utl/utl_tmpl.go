package utl

import (
	"embed"
	"github.com/tarqeem/template/utl/fs"
	"io"
	"log"
	"text/template"
)

var Views embed.FS
var TemplateFuncs template.FuncMap

type TemplateExecutor interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type DebugTemplateExecutor struct {
	Glob []string
}

func (e DebugTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t, err := GetTemplates()
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(wr, name, data)
}

type ReleaseTemplateExecutor struct {
	Template *template.Template
}

func (e ReleaseTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return e.Template.ExecuteTemplate(wr, name, data)
}

func GetTemplates() (*template.Template, error) {
	files, err := fs.GetFSFilesRecursively(&Views, "pages")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	ts := template.Must(template.New("").Funcs(TemplateFuncs).ParseFiles(files...))

	return ts, err

}
