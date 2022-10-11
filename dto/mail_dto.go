package dto

type Mail struct {
	From         string
	To           []string
	Subject      string
	BodyParam    map[string]interface{}
	TemplateHtml string
	Attachment   string
}
