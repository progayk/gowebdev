package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// check cookie
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name: "session",
			Value: "",
		}
	}
	// process form submission
	if r.Method == http.MethodPost {
		v := r.FormValue("mail")
		h := genMAC([]byte(v))
		c.Value = v + "|" + h
	}

	http.SetCookie(w, c)

	io.WriteString(w, `
	<html>
	<body>
	<form method="POST">
	<input type="text" name="mail" placeholder="mail">
	<input type="submit">
	</form>
	<a href="/authenticate">Validate This `+c.Value+`</a>
	</body>
	</html>
	`)
}

func auth(w http.ResponseWriter, r *http.Request) {
	// check cookie
	c, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// parse hashed key and mail from cookie
	mail := strings.Split(c.Value, "|")[0]
	mac := strings.Split(c.Value, "|")[1]
	fmt.Println(mail, mac)

	// validate hashed value
	if valid := checkMAC([]byte(mail), []byte(mac)); !valid {
		log.Println("cookie is not valid", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	fmt.Fprintf(w, "msg recieved: %s\nhashed value: %s", mail, mac)
}

// WARNING!
// private key shouldn't be constant
// TODO: search how to generate a random value for each unique session
var privateKey = []byte("privatekey")

func genMAC(msg []byte) string {
	mac := hmac.New(sha256.New, privateKey)
	mac.Write(msg)
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func checkMAC(msg, msgMAC []byte) bool {
	expectedMAC := genMAC(msg)
	return hmac.Equal(msgMAC, []byte(expectedMAC))
}