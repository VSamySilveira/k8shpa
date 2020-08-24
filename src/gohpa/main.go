package main

import (
	"fmt"
	"math"
	"net/http"
)

func raiz() string {
	x := 0.0001

	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "<b>Code.education Rocks!</b>"
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>"+raiz()+"<br>This is the new version using CI and CD!</body></html>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}
