package main

//Importing dependencies
import (
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//parser for templates
var tpl *template.Template

//Structure to represent user data
type User struct {
	UserName   string 
	EmailId string 
	PhoneNo string 
	Password string
	DateTime string
}

//init() initialises tpl to recognize all html files in the templates folder
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.html"))
}

//Hash function to calculate Hash of password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//main() function
func main (){
	//Fucntion definitions and url handlers
	http.HandleFunc("/",index)
	http.HandleFunc("/CreateUser",CreateUser)
	http.HandleFunc("/Search",Search)
	http.HandleFunc("/Delete",Delete)
	http.ListenAndServe(":8080",nil)
}

//Function to handle localhost:8080/
//Serves the main form which is 'formmod.html'
func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"formmod.html",nil)
}

//Function to handle localhost:8080/CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request){
	//Redirect to index if method is not POST
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
    //MySQL Configurations
    dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    //Establish Communication to MySQL Database
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()

	//Accessing Form Values
	_userName := r.FormValue("username")
	_userPassword := r.FormValue("password")
	hash,_ := HashPassword(_userPassword)
	_userEmailid := r.FormValue("emailid")
	_userPhonenumber := r.FormValue("phonenumber")
	
	//Excecuting Search Query
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	
	var user User
	
	//Storing results of query in user(structure User)
	for results.Next(){
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	//If emailId already exists
	if(len(user.EmailId)>0){	
		//Query to update table
		update, err:=db.Query("UPDATE userData SET userName='"+_userName+"',phoneNo='"+_userPhonenumber+"',password='"+hash+"',dateTime=NOW() where emailId='"+_userEmailid+"'");
		if err != nil{
			panic(err.Error())
		}
		defer update.Close()
		//Serving webpage "Display.html" to display message
		tpl.ExecuteTemplate(w,"Display.html","Data succesfully updated for "+user.EmailId)
	}else{
		//If emailid dosent exist in database
		//Query to insert into database from form
		insert, err := db.Query("INSERT INTO userData VALUES('"+_userName+"','"+_userEmailid+"','"+_userPhonenumber+"','"+_userPassword+"',NOW())")
		if err!= nil{
			panic(err.Error())
		}
		defer insert.Close()
		//Serve page 'Display.html' to display message Successfully Inserted
		tpl.ExecuteTemplate(w,"Display.html","Data succesfully Inserted")
	}
}
//Function to handle localhost:8080/Search
//Returns the row corresponding to given enail id if it exists in database else return null
func Search(w http.ResponseWriter,r *http.Request){
	//Redirect to index if method id not POST
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
    //MySQL configurations	
    dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    //Establish communication to database
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()
	
	//Get email id from form
	_userEmailid := r.FormValue("emailid")
	
	//Query to search database for existance of given email id
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	
	var user User
	
	//Storing results of query in user(structure User)
	for results.Next(){
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.EmailId)>0){	
		//If search successful print details on 'SearchResults.html'
		tpl.ExecuteTemplate(w,"SearchResults.html",user)
	}else{
		//If search unsuccessful print message on 'Display.html'
		tpl.ExecuteTemplate(w,"Display.html","Record Dosent Exist")
	}
}

//function to handle localhost:8080/Delete
func Delete(w http.ResponseWriter,r *http.Request){
	//Redirect to index if method id not POST
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
    //MySQL configurations
    dbDriver := "mysql"
    dbUser := "dummyUser"
    dbPass := "dummyUser01"
    dbName := "db_intern"
    //Establish communication to database
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err!= nil{
		panic(err.Error())
	}

	defer db.Close()

	//Get email id from form
	_userEmailid := r.FormValue("emailid")
	//Query to search database for existance of given email id in database
	results,err :=db.Query("Select * from userData where emailId='"+_userEmailid+"'");
	if err!=nil{
		panic(err.Error())
	}
	var user User
	//Storing results of query in user(structure User)
	for results.Next(){
		err= results.Scan(&user.UserName,&user.EmailId,&user.PhoneNo,&user.Password,&user.DateTime)
		if err!= nil{
			panic(err.Error())
		}
	}
	if(len(user.EmailId)>0){
		//If search successful delete the row corresponding to email id
		delete,err :=db.Query("Delete from userData where emailId='"+_userEmailid+"'")
		if err!=nil{
			panic(err.Error())
		}
		//Display message on 'Display.html'
		tpl.ExecuteTemplate(w,"Display.html","Record found and deleted!")
		defer delete.Close()	
	}else{
		//If search unsuccessful display message on 'Display.html'
		tpl.ExecuteTemplate(w,"Display.html","Record dosen't exist!")
	}
}
