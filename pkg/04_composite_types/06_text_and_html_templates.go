package gosamplecode

import (
	"html/template"
	"log"
	"os"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F30 function
func F30() {
	const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age: 	{{.CreatedAt | daysAgo}} days
{{end}}
`

	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": library.DaysAgo}).
		Parse(templ)

	if nil != err {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, F29()); nil != err {
		log.Fatal(err)
	}
}

// F31 function
func F31() {
	var templ = `<!DOCTYPE html>
<html>

<body>

    <h1>{{.TotalCount}} issues</h1>
    <table>
        <tr style='text-align: left'>
            <th>#</th>
            <th>State</th>
            <th>User</th>
            <th>Title</th>
        </tr>
        {{range .Items}}
        <tr>
            <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
            <td>{{.State}}</td>
            <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
            <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
        </tr>
        {{end}}
    </table>

</body>

</html>
`
	var issueList = template.Must(template.New("issuelist").
		Parse(templ))

	if err := issueList.Execute(os.Stdout, F29()); nil != err {
		log.Fatal(err)
	}
}

// F32 function
func F32() {
	const templ = `<p>A: {{.A}}</p>
<p>B: {{.B}}</p>
`
	t := template.Must(template.New("escape").
		Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}

	data.A = "<b>Hello, World!</b>"
	data.B = "<b>Hello, World!</b>"

	if err := t.Execute(os.Stdout, data); nil != err {
		log.Fatal(err)
	}
}
