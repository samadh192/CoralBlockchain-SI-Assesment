package main

import (
	"html/template"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

type User struct {
	userName   string 
	emailId string 
	phoneNo string 
	password string
	dateTime string
}

func init(){
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}

func main (){
	http.HandleFunc("/",index)
	http.HandleFunc("/CreateUser",CreateUser)
	http.HandleFunc("/Search",Search)
	http.HandleFunc("/Delete",Delete)
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"formmod.html",nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()

	//fmt.Println("Successfully connected")


	_userName := r.FormValue("username")
	_userPassword := r.FormValue("password")
	_userEmailid := r.FormValue("emailid")
	_userPhonenumber := r.FormValue("phonenumber")
	
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	var user User
	for results.Next(){
		err= results.Scan(&user.userName,&user.emailId,&user.phoneNo,&user.password,&user.dateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	//fmt.Println(user.emailId)
	if(len(user.emailId)>0){	
		//fmt.Fprintln(w,"Already exists")
		//code to update table
		update, err:=db.Query("UPDATE userData SET userName='"+_userName+"',phoneNo='"+_userPhonenumber+"',password='"+_userPassword+"',dateTime=NOW() where emailId='"+_userEmailid+"'");
		if err != nil{
			panic(err.Error())
		}
		defer update.Close()
		fmt.Fprintln(w,"Data succesfully update for "+user.emailId)
	}else{
		//fmt.Fprintln(w,"Dosent exist")
		//code to insert into table
		insert, err := db.Query("INSERT INTO userData VALUES('"+_userName+"','"+_userEmailid+"','"+_userPhonenumber+"','"+_userPassword+"',NOW())")
		if err!= nil{
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Fprintln(w,"Data successfully Inserted")
	}
	
	//fmt.Println("Successfully inserted")
}
func Search(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()

	_userEmailid := r.FormValue("emailid")
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	var user User
	for results.Next(){
		err= results.Scan(&user.userName,&user.emailId,&user.phoneNo,&user.password,&user.dateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.emailId)>0){	
		fmt.Fprintln(w,"Username:"+user.userName)
		fmt.Fprintln(w,"EmailId:"+user.emailId)
		fmt.Fprintln(w,"PhoneNo:"+user.phoneNo)
		fmt.Fprintln(w,"Password:"+user.password)
		fmt.Fprintln(w,"Time:"+user.dateTime)
	}else{
		fmt.Fprintln(w,"Record Dosent exist!")
	}
}
func Delete(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()

	_userEmailid := r.FormValue("emailid")
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	var user User
	for results.Next(){
		err= results.Scan(&user.userName,&user.emailId,&user.phoneNo,&user.password,&user.dateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.emailId)>0){	
		delete,err :=db.Query("Delete from userData where emailId='"+_userEmailid+"'")
		if err!=nil{
			panic(err.Error())
		}
		fmt.Fprintln(w,"Record found and deleted!")
		defer delete.Close()	
	}else{
		fmt.Fprintln(w,"Record Dosent exist!")
	}
}