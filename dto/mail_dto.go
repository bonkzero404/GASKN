package dto

type Mail struct {
	From         string
	To           []string
	Subject      string
	BodyParam    any
	TemplateHtml string
	Attachment   string
}
