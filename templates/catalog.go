package templates

import "html/template"

//GetCatalogTemplate GetCatalogTemplate
func GetCatalogTemplate() (*template.Template, error) {
	return template.New("").Parse(catalogTemplate)
}

const catalogTemplate = `
// Code generated by go generate; DO NOT EDIT.
package {{.PackageName}}

import (
	"github.com/docker/oscalkit/types/oscal/catalog"
)

var ApplicableControls = []catalog.Catalog{
{{range .Catalogs}}
	catalog.Catalog{
		Title: "{{ .Title }}",
		Groups: []catalog.Group{
			{{range .Groups}}
				catalog.Group{
					Id:  "{{.Title}}",
					Controls: []catalog.Control{
						{{range .Controls}}
							catalog.Control{
								Id: 	"{{.Id}}",
								Class: 	"{{.Class}}",
								Title:	"{{.Title}}",
								Subcontrols: []catalog.Subcontrol{
									{{range .Subcontrols}}
									catalog.Subcontrol{
										Id: 	"{{.Id}}",
										Class: 	"{{.Class}}",
										Title:	"{{.Title}}",		
									},
									{{end}}
								},
							},
						{{end}}
					},
				},
			{{end}}
		},
	},
{{end}}
}
`
