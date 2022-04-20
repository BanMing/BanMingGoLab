package mail

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func SendReportEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	em := email.NewEmail()
	em.From = "ttttt@outlook.com"
	em.To = []string{"tttt@outlook.com.cn"}
	em.Subject = "Orca Report Email Test"
	em.Text = []byte("Reporter ID:" + r.Form["ReporterPlayerID"][0] + "\n" +
		"Reported ID:" + r.Form["ReportedPlayerID"][0] + "\n" +
		"State Reason:" + r.Form["StatedReason"][0] + "\n" +
		"Description:" + r.Form["Description"][0] + "\n")

	em.Attach(strings.NewReader(r.Form["ReplayFile"][0]), r.Form["ReplayFileName"][0]+".replay", "text")

	// smtp server configuration.
	smtpHost := "smtp.office365.com"

	tslconfig := &tls.Config{
		ServerName: smtpHost,
	}

	auth := LoginAuth(em.From, "Welcome2022!")
	err := em.SendWithStartTLS("smtp.office365.com:587", auth, tslconfig)

	if err != nil {
		log.Fatal(err)
		w.Write([]byte("0"))
	} else {
		w.Write([]byte("1"))
		log.Println("send successfully ... ")
	}
}

func main() {
	log.Println("start test smtp server....")
	http.HandleFunc("/report", SendReportEmail)
	http.ListenAndServe(":9999", nil)
}
