package main

import "net/http"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("hello, world"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
