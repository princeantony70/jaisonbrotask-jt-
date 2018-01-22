package main

import (

          "net/http"
          "html/template"
          "database/sql"
          "log"
          _ "github.com/go-sql-driver/mysql"
        )



var db *sql.DB
var err error
var tpl *template.Template




type question struct {
QID int
Question string
Answer1 string
Answer2 string
Answer3 string
}


func checkErr(err error) {
    if err != nil {
   log.Fatalln(err)
        }
    }



func init(){
  db, err = sql.Open("mysql", "root:nfn@/shift")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
}



func main(){
//  defer db.close()
   http.HandleFunc("/",index)
   http.HandleFunc("/takeone", takeone)
   http.HandleFunc("/Createone", Createone)
   http.HandleFunc("/thankyou", thankyou)
   log.Println("Server is up on  port")
   log.Fatalln(http.ListenAndServe(":7016", nil))
}





func index(w http.ResponseWriter, req *http.Request){

  tpl,err :=template.ParseFiles("index.html")

  if err != nil{
		log.Fatalln("error parsing template index",err)
	}

	err =tpl.ExecuteTemplate(w,"index.html",nil)
	if err !=nil{
		log.Fatalln("error executing template index ",err)
	}
	checkErr(err)
}


func takeone(w http.ResponseWriter, req *http.Request){
  tpl,err :=template.ParseFiles("takeone.html")

  if err != nil{
    log.Fatalln("error parsing template takeone",err)
  }

  err =tpl.ExecuteTemplate(w,"takeone.html",nil)
  if err !=nil{
    log.Fatalln("error executing template takeone" ,err)
  }
  checkErr(err)
  }






func Createone(w http.ResponseWriter, req *http.Request){
if req.Method == http.MethodPost{
qus := question{}
qus.Question = req.FormValue("question")
qus.Answer1  = req.FormValue("answer1")
qus.Answer2  = req.FormValue("answer2")
qus.Answer3  = req.FormValue("answer3")
checkErr(err)
_,err = db.Exec(
  "INSERT INTO  questions (question, answer1, answer2,answer3) VALUES (?, ?, ?, ?)",
  qus.Question,
  qus.Answer1,
  qus.Answer2,
  qus.Answer3,
)
checkErr(err)
http.Redirect(w, req, "/thankyou", http.StatusSeeOther)
return
}
http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}






func thankyou(w http.ResponseWriter, req *http.Request){
tpl,err :=template.ParseFiles("thankyou.html")

if err != nil{
  log.Fatalln("error parsing template thankyou",err)
}

err =tpl.ExecuteTemplate(w,"thankyou.html",nil)
if err !=nil{
  log.Fatalln("error executing template thank you",err)
}
checkErr(err)
}
