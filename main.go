package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// custom data type
type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// value type for this map is template.Template
// key type for this map is string
// map[keyType]*valueType
var templates = make(map[string]*template.Template, 3)

// You can define and initialize variables directly in code files, but the most useful language features can be done only inside functions.
func loadTemplates() {
	// Go’s concise syntax, which can be used only within functions.
	// This syntax specifies the name, followed by a colon (:), the assignment operator (=), and then a value.
	// This is a fixed-length Array
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	// For loop:  The type of the first variable is always int, which is a built-in Go data type for representing integers.
	// The range keyword is used with the for keyword to enumerate arrays, slices, and maps.
	for index, name := range templateNames {
		//The ParseFiles function returns two results, a pointer to a template.Template value and an error, which is the built-in data type for representing errors in Go.
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil { // The Go null value is nil, for some strange reason.
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			//Go provides a function named panic that can be called when an unrecoverable error happens.
			panic(err)
		}
	}
}

// create "array slice" (variable length array)
// "make" initializes a new slice with args (data type, initial size, initial capacity)
// will be resized automatically as items are added, once it reaches initial capacity
// [] denote a slice, * denotes a pointer
// "A slice of pointers to instances of the Rsvp struct"
var responses = make([]*Rsvp, 0, 10)

// The second argument is a pointer to an instance of the Request struct, defined in the net/http package,
// which describes the request being processed. The first argument is an example of an interface,
// which is why it isn’t defined as a pointer.
func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

// ResponseWriter can be used by any code that knows how to write data using the Writer interface.
func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

// https://learning.oreilly.com/library/view/pro-go-the/9781484273555/html/512642_1_En_1_Chapter.xhtml
// The main entry point for a GO Application.  To run:  "go run ."
// This compiles and executes in one step.  Useful during development.
func main() { //THE OPENING BRACE HAS TO BE ON THE SAME LINE AS THE FUNCTION DECLARATION, OR IT WON'T COMPILE.
	loadTemplates()
	//The net/http package defines the HandleFunc function, which is used to specify a URL path and the handler that will receive matching requests.
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	// Create an HTTP Server that listens on port 5000
	// The second argument is nil, which tells the server that requests should be processed using the functions registered with the HandleFunc function.
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
