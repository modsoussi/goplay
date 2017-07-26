package main

import (
	"fmt"
	"net/http"
	"strings"

	"errors"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/kidstuff/mongostore"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var store *mongostore.MongoStore
var sess *mgo.Session

// GetHandler handles get request
func GetHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "ssid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("TOK")
	if err != nil {
		http.Error(w, `{"status":"error","message":"unauthorized request"}`, http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("modsoussi"), nil
	})

	if token.Valid {
		fmt.Fprintf(w, `{"name":"%s"}`, session.Values["name"])
		return
	}

	http.Error(w, `{"status":"error","message":"unauthorized request"}`, http.StatusUnauthorized)
}

// PostHandler handles post request
func PostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "ssid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var name string
	r.ParseForm()

	for k, v := range r.Form {
		if k == "name" {
			name = strings.Join(v, "")
		}
	}

	cookie, err := r.Cookie("TOK")
	if err != nil {
		http.Error(w, `{"status":"error","message":"unauthorized request"}`, http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("modsoussi"), nil
	})

	if token.Valid {
		session.Values["name"] = name
		session.Save(r, w)

		fmt.Fprint(w, `{"status":"success"}`)
		return
	}

	http.Error(w, `{"status":"error","message":"unauthorized request"}`, http.StatusUnauthorized)
}

// EndHandler ends a session
func EndHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "ssid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("TOK")
	if err != nil {
		http.Error(w, `{"status":"error","message":"unauthorized request"}`, http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("modsoussi"), nil
	})

	if token.Valid {
		s := getMongoSession()
		defer s.Close()

		invalidJWT := s.DB("local").C("invalidJWT")

		count, err := invalidJWT.Find(bson.M{"token": token.Raw}).Count()
		if err != nil || count == 0 {
			invalidJWT.Insert(bson.M{"token": token.Raw})
		}
	}

	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Fprint(w, `{"status":"success"}`)
}

// RegisterHandler handles registering a user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var username, password string
	r.ParseForm()

	for k, v := range r.Form {
		if k == "username" {
			username = strings.Join(v, "")
		} else if k == "password" {
			password = strings.Join(v, "")
		}
	}

	err := register(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, `{"status":"success"}`)
}

// LoginHandler handles logging a user in
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var username, password string
	r.ParseForm()

	for k, v := range r.Form {
		if k == "username" {
			username = strings.Join(v, "")
		} else if k == "password" {
			password = strings.Join(v, "")
		}
	}

	if username == "" || password == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := authenticate(username, password)
	if err != nil {
		fmt.Fprintf(w, `{"status":"error", "message":"%s"}`, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       "http://modsoussi.tech",
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"modsoussi": "sandra"})

	tokenString, err := token.SignedString([]byte("modsoussi"))
	if err != nil {
		fmt.Println(err)
	}

	session, err := store.Get(r, "ssid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Save(r, w)

	cookie := &http.Cookie{
		Name:    "TOK",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
		Domain:  "localhost"}

	http.SetCookie(w, cookie)

	fmt.Fprint(w, `{"status":"success"}`)
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/end", EndHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))

	defer sess.Close()
}

// Helper Methods

func init() {
	sess = getMongoSession()
	store = mongostore.NewMongoStore(sess.DB("local").C("sessions"), 3600*24*7, true, []byte("nsansansa"))
}

func getMongoSession() *mgo.Session {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	return s
}

func register(username, password string) error {
	s := getMongoSession()
	defer s.Close()

	users := s.DB("local").C("users")

	err := users.Insert(bson.M{"username": username, "password": password})
	if err != nil {
		return err
	}

	return nil
}

func authenticate(username, password string) error {
	s := getMongoSession()
	defer s.Close()

	users := s.DB("local").C("users")

	count, err := users.Find(bson.M{"username": username, "password": password}).Count()
	if err != nil || count == 0 {
		return errors.New("user not found")
	}

	return nil
}
