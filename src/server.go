package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("foo")
	http.HandleFunc("/test", testController)
	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.ListenAndServe(":8888", nil)
}

func testController(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "bar")
}
