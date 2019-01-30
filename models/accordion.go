package models

// Option model for html
type Option struct {
	OptionName string
}

// Panel model for html
type Panel struct {
	PanelName string
	Options   []Option
}

// Accordion model for html
type Accordion struct {
	Panels []Panel
}
