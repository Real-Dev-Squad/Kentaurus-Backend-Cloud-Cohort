package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	//"io/ioutil"
	"github.com/gorilla/mux"
)

type Word struct{
	Word string 
}

//word:="Abhay"
var word string
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome Don")
}

func getWord(w http.ResponseWriter, r *http.Request){
	//word:=""
	wordv:=Word{Word:word}
	//fmt.Fprintf(wordv.Word)
	json.NewEncoder(w).Encode(wordv)
}

func postWord(w http.ResponseWriter,r *http.Request){
	//reqBody, _ := ioutil.ReadAll(r.Body)
	vars:=mux.Vars(r)
	word=vars["word"]
	//fmt.Println(word)
	wordv:=Word{Word:word}
	//fmt.Fprintf(wordv.Word)
	json.NewEncoder(w).Encode(wordv)
}

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/get",getWord)
	myRouter.HandleFunc("/post/{word}",postWord).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8090",myRouter))
}

func main(){
	word="1st"
	handleRequest()
}
