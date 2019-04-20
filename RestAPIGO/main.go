package main

import (
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

type User struct {
	UserName   string 
	EmailId string 
	PhoneNo string 
	Password string
	DateTime string
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
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	//fmt.Println(user.EmailId)
	if(len(user.EmailId)>0){	
		//fmt.Fprintln(w,"Already exists")
		//code to update table
		update, err:=db.Query("UPDATE userData SET userName='"+_userName+"',phoneNo='"+_userPhonenumber+"',password='"+_userPassword+"',dateTime=NOW() where emailId='"+_userEmailid+"'");
		if err != nil{
			panic(err.Error())
		}
		defer update.Close()
		//fmt.Fprintln(w,"Data succesfully updated for "+user.EmailId)
		tpl.ExecuteTemplate(w,"Display.html","Data succesfully updated for "+user.EmailId)
	}else{
		//fmt.Fprintln(w,"Dosent exist")
		//code to insert into table
		insert, err := db.Query("INSERT INTO userData VALUES('"+_userName+"','"+_userEmailid+"','"+_userPhonenumber+"','"+_userPassword+"',NOW())")
		if err!= nil{
			panic(err.Error())
		}
		defer insert.Close()
		//fmt.Fprintln(w,"Data successfully Inserted")
		tpl.ExecuteTemplate(w,"Display.html","Data succesfully Inserted")
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
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.EmailId)>0){	
		/*fmt.Fprintln(w,"Username:"+user.UserName)
		fmt.Fprintln(w,"EmailId:"+user.EmailId)
		fmt.Fprintln(w,"PhoneNo:"+user.PhoneNo)
		fmt.Fprintln(w,"Password:"+user.Password)
		fmt.Fprintln(w,"Time:"+user.DateTime)*/
		tpl.ExecuteTemplate(w,"SearchResults.html",user)
	}else{
		tpl.ExecuteTemplate(w,"Display.html","Record Dosent Exist")
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
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.EmailId)>0){	
		delete,err :=db.Query("Delete from userData where emailId='"+_userEmailid+"'")
		if err!=nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w,"Display.html","Record found and deleted!")
		defer delete.Close()	
	}else{
		tpl.ExecuteTemplate(w,"Display.html","Record dosen't exist!")
	}
}
