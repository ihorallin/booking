package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntgMap   map[string]int
	FloatgMap map[string]float32
	Date      map[string]interface{}
	CSFRToken string
	Flash     string
	Warning   string
	Error     string
}
