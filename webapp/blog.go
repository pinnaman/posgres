package main

import (
	"ajay/models"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

const staticURL string = "/static/"
const staticRoot string = "static/"
const postPath string = "posts"

const tmplSrc = `<table>
<tr align="center">
{{ range . }}
		<td style="color: rgba(55, 248, 255, 1);">{{.Id}}</td>
		<td style="color: rgba(200, 200, 55, 1);">{{.Created}}</td>
{{ end }}
</tr>
</table>`

// compile
var tmplTab = template.Must(template.New("tmpl").Parse(tmplSrc))

//Compile templates on start
var templates = template.Must(template.ParseFiles("templates/header.html",
	"templates/footer.html",
	"templates/main.html",
	"templates/about.html",
	"templates/test_jscript.html",
	"templates/events.html",
),
)

//A Page structure
type Page struct {
	Title  string
	Static string
}

type Post struct {
	Title   string
	Date    string
	Summary string
	Body    string
	File    string
}

func logRequest(req *http.Request) {
	now := time.Now()
	log.Printf("%s - %s [%s] \"%s %s %s\" ",
		req.RemoteAddr,
		"",
		now.Format("02/Jan/2006:15:04:05 -0700"),
		req.Method,
		req.URL.RequestURI(),
		req.Proto)
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

// dynamic route
func javascript(w http.ResponseWriter, req *http.Request) {
	//js_templ.Execute(w, req.FormValue("s"))
	display(w, "test_jscript", &Page{Title: "Javascript", Static: staticRoot})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About", Static: staticRoot})
}

//The handlers.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	display(w, "main", &Page{Title: "Home", Static: staticRoot})

	posts := getPosts()
	fmt.Println(len(posts))
	//context := Context{Title: "Home"}
	//render(w, "index", context)
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		posts := getPosts()
		t := template.New("index.html")
		t, _ = t.ParseFiles("index.html")
		t.Execute(w, posts)
	} else {
		f := "static/posts/" + r.URL.Path[1:] + ".md"
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		post := Post{title, date, summary, body, r.URL.Path[1:]}
		t := template.New("post.html")
		t, _ = t.ParseFiles("post.html")
		t.Execute(w, post)
	}
}

// Need this to get the Jquery library in the DOM page...
func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(staticURL):]
	if len(static_file) != 0 {
		f, err := http.Dir(staticRoot).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func getPosts() []Post {
	a := []Post{}
	files, _ := filepath.Glob("static/posts/*")
	for _, f := range files {
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		a = append(a, Post{title, date, summary, body, file})
	}
	return a
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	evts, err := models.AllEvents()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Pass events to template...
	display(w, "events", &Page{Title: "Event", Static: staticRoot})

	for _, evt := range evts {
		//fmt.Println(evt)
		fmt.Fprintf(w, "%s, %s, %s, %s, %s, %s\n", evt.Id, evt.Created, evt.Start, evt.End, evt.Title, evt.Completed)
	}
	fmt.Fprintf(w, "%s\n", "Display Table Here...")

	//display(w, tmplSrc, evts)
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/events", eventHandler)
	http.HandleFunc("/javascript", javascript)
	http.HandleFunc(staticURL, StaticHandler)

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
