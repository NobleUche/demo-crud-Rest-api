 package main
import (
"fmt"
"github.com/gorilla/mux"
"net/http"
"log"
"encoding/json"
)

// structv declaration

type Course struct{
	Id string
	Name string
	Level string
	Hod string
}
// struct variable

var courses []Course

// api functions

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello world")
}
func getcourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
		if r.Method=="POST"{
			req:=r.Body
			json.NewDecoder(req).Decode(&courses)
		}else{
				fmt.Println("These are the various courses:",courses)
				json.NewEncoder(w).Encode(courses)
}
}
func getcourse(w http.ResponseWriter, r *http.Request){
	v:=mux.Vars(r)
	if r.Method=="GET"{
		for _, items:=range courses{
			if v["courseid"]==items.Id{
				fmt.Println("These are the courses:",items)
				json.NewEncoder(w).Encode(items)
			}
		}
	}else if r.Method=="PUT"{
		// do whatever
		for _, items:=range courses{
			if v["courseid"]==items.Id{
				fmt.Println("These are the courses:",items)
				json.NewDecoder(r.Body).Decode(&courses)// we do this cos we have to decode the retrieved json b4 we perfom PUT
				//update the json to the database i.e struct
				json.NewEncoder(w).Encode(items)// we do this cos we are done updating and so we can send it back to the server as json
			}
		}	
	}else if r.Method=="DELETE"{
		for _, items:=range courses{
			if v["courseid"]==items.Id{
				// devise delete method
				data:=items[v["courseid"]]
				courses=delete(courses,data)
			}
		}
		//do whatever you like 
	}
}

// main function
func main(){
	courses=append(courses,Course{Id:"1",Name:"BIO",Level:"100",Hod:"GKP"})
	courses=append(courses,Course{Id:"2",Name:"MedSurg",Level:"500",Hod:"BPL"})
	courses=append(courses,Course{Id:"3",Name:"VET",Level:"300",Hod:"KGB"})
	router:=mux.NewRouter()
	router.HandleFunc("/",home)
	router.HandleFunc("api/courses",getcourses).Methods("GET","POST")
	router.HandleFunc("api/course/{courseid}",getcourse).Methods("GET","PUT","DELETE")
	server:=http.Server{
		Addr:":8000",
		Handler:router,
	}
	log.Fatal(server.ListenAndServe())
}
