package main

import (
	"fmt"
	"net/http"
	"time"
)

//cookie内容
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

func cookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration.Year() += 1
	cookie := http.Cookie{Name: "username", Value: "banming", Expires: expiration}
	http.SetCookie(w, &cookie)

	for _, cookie := range r.Cookies() {
		fmt.Println(cookie)
		//fmt.Printf(w,cookie.Name)
	}

}

func main() {

}
