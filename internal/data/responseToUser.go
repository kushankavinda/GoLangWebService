package data

type ResponseToUser struct {
	Title         string
	HtmlVersion   string
	HeadLines     string
	NumberOfLinks int
	IsLoginForm   bool
}

type InternalAndExternalLinks struct {
	NumberOfLinks             int
	NumberOfInaccessibleLinks int
}
