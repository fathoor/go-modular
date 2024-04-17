package templates

var ModelTmpl = `package model

type {{.Name}}Request struct {
}

type {{.Name}}Response struct {
}

type {{.Name}}PageResponse struct {
	Page int
	Size int
	Total int
	{{.Name}} []{{.Name}}Response
}
`
