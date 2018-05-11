package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./client/dist"))))
	http.HandleFunc("/api/login", login)
	http.ListenAndServe(":3000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	setupXSS(w)
	fmt.Fprintf(w, "{ \"result\" : \"OK\" , \"userid\" : \"#oneplus\" }")
}

func setupXSS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

/*
func (b *baseCtrl) setupRespHeader(ctx *gin.Context) {
	if b.ngDev {
		//เฉพาะเวลา develop เท่านั้น
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
*/
