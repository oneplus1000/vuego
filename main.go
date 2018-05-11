package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("43232467687834343432767"))

const sessionName = "vuego-session"

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./client/dist"))))
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/is_logined", isLogined)
	http.ListenAndServe(":3000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	err := readJSONFromBody(r, &req)
	if err != nil {
		setupXSS(w)
		fmt.Fprintf(w, "{ \"result\" : \"FAIL\" }")
		return
	}

	if req.Username != "one" {

		session, err := store.Get(r, sessionName)
		if err != nil {
			setupXSS(w)
			fmt.Fprintf(w, "{ \"result\" : \"FAIL\" }")
			return
		}
		session.Values["logined"] = ""
		session.Save(r, w)

		setupXSS(w)
		fmt.Fprintf(w, "{ \"result\" : \"FAIL\" }")
		return
	}

	session, err := store.Get(r, sessionName)
	if err != nil {
		setupXSS(w)
		fmt.Fprintf(w, "{ \"result\" : \"FAIL\" }")
		return
	}
	session.Values["logined"] = "true"
	session.Save(r, w)

	setupXSS(w)
	fmt.Fprintf(w, "{ \"result\" : \"OK\" , \"userid\" : \"#oneplus\" }")

}

func isLogined(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		setupXSS(w)
		fmt.Fprintf(w, "{ \"result\" : \"FAIL\" }")
		return
	}
	log.Printf("%s", session.Values["logined"])
	if session.Values["logined"] != "true" {
		setupXSS(w)
		fmt.Fprintf(w, "{ \"logined\" : false }")
		return
	}
	setupXSS(w)
	fmt.Fprintf(w, "{ \"logined\" : true }")
	return
}

func setupXSS(w http.ResponseWriter) {
	/*
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	*/
}

//ReadStringFromBody read string from http body
func readStringFromBody(r *http.Request) (string, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//ReadJSONFromBody read json from http body
func readJSONFromBody(r *http.Request, obj interface{}) error {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	return nil
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

type loginReq struct {
	Username string
	Password string
}
