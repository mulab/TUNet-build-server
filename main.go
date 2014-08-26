// tunet_build_server project main.go
package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

var __username string
var __password string
var __property string
var __antpath string
var __sdkdir string
var store *sessions.CookieStore //= sessions.NewCookieStore([]byte("!QAZse4RFVgy7UJMko0"))
var session_name = "user_session"
var mutex = &sync.Mutex{}

type Configuration struct {
	Username string
	Password string
	Property string
	Ant      string
	Sdk      string
	Secret   string
	Port     string
}

func main() {
	configfile, err := os.Open("conf.json")
	ErrorHandler(err)
	decoder := json.NewDecoder(configfile)
	var config Configuration
	err = decoder.Decode(&config)
	ErrorHandler(err)
	__username = config.Username
	__password = config.Password
	__property = config.Property
	__antpath = config.Ant
	__sdkdir = config.Sdk
	store = sessions.NewCookieStore([]byte(config.Secret))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/build", BuildHandler)
	http.Handle("/action", websocket.Handler(ActionHandler))
	http.HandleFunc("/download/", DownloadHandler)
	http.ListenAndServe(config.Port, context.ClearHandler(http.DefaultServeMux))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[9:])
	http.ServeFile(w, r, "TUNet/aTUNet/target"+r.URL.Path[9:])
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	ErrorHandler(err)
	err = t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	form := r.PostFormValue
	username := form("username")
	password := form("password")
	fmt.Println("/login", "username", username, "password", password)
	if username == __username && password == __password {
		session, _ := store.Get(r, session_name)
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/build", 301)
	} else {
		fmt.Fprintf(w, "login failed")
	}
}

func BuildHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, session_name)
	if session.Values["username"] != __username {
		http.Redirect(w, r, "/", 302)
	}
	t, err := template.ParseFiles("build.html")
	ErrorHandler(err)
	err = t.Execute(w, nil)
}

type Wsout struct {
	ws *websocket.Conn
}

func (wsout Wsout) Write(p []byte) (n int, err error) {
	mutex.Lock()
	defer func() { mutex.Unlock() }()
	return wsout.ws.Write(p)
}

type Wserr struct {
	ws *websocket.Conn
}

func (wserr Wserr) Write(p []byte) (n int, err error) {
	mutex.Lock()
	defer func() { mutex.Unlock() }()
	return wserr.ws.Write(p)
}

func ActionHandler(ws *websocket.Conn) {
	c := make(chan string)
	end := make(chan string)
	r := ws.Request()
	stdout := Wsout{ws}
	stderr := Wserr{ws}
	session, _ := store.Get(r, session_name)
	if session.Values["username"] != __username {
		ws.Close()
	}
	go func() {
		var err error
		ws.Write([]byte("building...")) //call ant main
		err = os.Chdir("TUNet")
		if err != nil {
			c <- err.Error()
			c <- "Cloning repo"
			clone := exec.Command("git", "clone", "--depth", "1", "--single-branch", "--branch", "master", "git@github.com:mulab/TUNet.git")
			clone.Stdout = stdout
			clone.Stderr = stderr
			err = clone.Run()
			if err != nil {
				end <- err.Error()
				return
			}
			err = os.Chdir("TUNet")
		}
		c <- "Running git pull..."
		reset := exec.Command("git", "reset", "HEAD", "--hard")
		reset.Stdout = stdout
		reset.Stderr = stderr
		err = reset.Run()
		if err != nil {
			end <- err.Error()
			return
		}
		cmd := exec.Command("git", "pull", "-f", "--no-tags")
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		err = cmd.Run()
		if err != nil {
			c <- err.Error()
		}
		c <- "Building..."
		err = os.Chdir("aTUNet")
		if err != nil {
			end <- err.Error()
			return
		}
		cp := exec.Command("cp", "-f", __property, "./")
		cp.Stdout = stdout
		cp.Stderr = stderr
		err = cp.Run()
		if err != nil {
			end <- err.Error()
			return
		}
		ant := exec.Command(__antpath, "main", "-q", "-DAsIs=true", "-Dsdk.dir="+__sdkdir)
		ant.Stdout = stdout
		ant.Stderr = stderr
		err = ant.Run()
		if err != nil {
			end <- err.Error()
			return
		}
		end <- "done"
	}()
	for {
		select {
		case msg := <-c:
			ws.Write([]byte(msg))
		case msg := <-end:
			ws.Write([]byte(msg))
			ws.Close()
		}
	}
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}
