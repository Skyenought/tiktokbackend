{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}req {{.request}}{{if .stream}},stream {{
.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	// TODO: add your logic here and delete this line

	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}
