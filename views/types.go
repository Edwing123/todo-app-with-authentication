package views

// Represents page meta data, information that goes into the "head" element
type Meta struct {
	Title       string
	Description string
}

// Represents page related information, like meta data
type Page struct {
	Meta Meta
}

// Represents the data passed to every view when it's going to be executed
type ViewContext struct {
	Page Page
	Data interface{}
}
