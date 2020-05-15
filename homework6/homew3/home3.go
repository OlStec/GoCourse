package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	nameValue := req.FormValue("name")

	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello `+nameValue+`!</title>
	</head>
	<body>
	     Hello `+nameValue+`!
	</body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
