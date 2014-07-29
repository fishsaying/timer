package main 

import "fmt"
import "net/http"
import "strconv"
import "strings"
//import "msg"
//import "log"
//import "os"
import "github.com/drone/routes"

var logfile string = "run.log"

func main() {
	mux := routes.New()
	
    mux.Get("/notifications", index)
    mux.Post("/notifications", add)
    mux.Patch("/notifications", edit)
    mux.Del("/notifications", remove)
    
   // logfile, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
   // if err != nil {
   //     fmt.Println("Can not open log file", err)
   //     log.Fatal(err)
   // }
   // log.SetOutput(logfile)
   // log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

    http.Handle("/", mux)
    http.ListenAndServe(":8000", nil)
    
}

//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "rootHandler: %s\n", r.URL.Path)
//    fmt.Fprintf(w, "URL: %s\n", r.URL)
//    fmt.Fprintf(w, "Method: %s\n", r.Method)
//    fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI )
//    fmt.Fprintf(w, "Proto: %s\n", r.Proto)
//    fmt.Fprintf(w, "HOST: %s\n", r.Host) 
//}

func index(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    lastName := params.Get("last")
    firstName := params.Get("first")
    fmt.Fprintf(w, "you are %s %s", firstName, lastName)
}

func add(w http.ResponseWriter, r *http.Request) {
    untilRaw := r.FormValue("until")
    if untilRaw == "" {
    	http.Error(w, "`until` is empty", http.StatusInternalServerError)
    }
    until, err := strconv.Atoi(untilRaw)
    if err != nil || until <= 0 {
    	http.Error(w, "`until` is invalid", http.StatusInternalServerError)
    }
    message := r.FormValue("message")
    message = strings.Trim(message, " ")
    if message == "" {
    	http.Error(w, "`message` is empty", http.StatusInternalServerError)
    }
//    noti := &msg.Notification{}
    // fmt.Fprintf(w, "you are %s %s", first, last)
    
    fmt.Fprintf(w, "msg: %s will expire after %d", message, until)
}

func edit(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    lastName := params.Get(":last")
    firstName := params.Get(":first")
    fmt.Fprintf(w, "you are %s %s", firstName, lastName)
}

func remove(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    lastName := params.Get(":last")
    firstName := params.Get(":first")
    fmt.Fprintf(w, "you are %s %s", firstName, lastName)
}
