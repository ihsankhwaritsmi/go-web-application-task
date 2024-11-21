package main

import (
	"context"
	"os"
)

func main() {

	button(false).Render(context.Background(), os.Stdout)

	// component := Hello("walalal")

	// http.Handle("/", templ.Handler(component))
	// fmt.Println("Listening on http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}
