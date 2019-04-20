#Importing required Libraries
from flask import Flask, render_template,make_response
from flask_restful import Resource, Api
from flask_restful import reqparse
from flaskext.mysql import MySQL

#Initializing Flask object
app = Flask(__name__)

#Initiaizing mysql object
mysql = MySQL()

# MySQL configurations
app.config['MYSQL_DATABASE_USER'] = 'dummyUser'
app.config['MYSQL_DATABASE_PASSWORD'] = 'dummyUser01'
app.config['MYSQL_DATABASE_DB'] = 'db_intern'
app.config['MYSQL_DATABASE_HOST'] = 'db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com'


mysql.init_app(app)
api = Api(app)

#Class for main Page localhost:5000/
class FormPage(Resource):
    def __init__(self):
        pass
    #Function to render webpage 'formmod.html' when requested through a GET Request
    def get(self):
        headers ={'Content-Type':'text/html'}
        return make_response(render_template('formmod.html'),200,headers)

#Class corresponding to localhost:5000/CreateUser
#Takes values from 'formmod.html' and enters into database if the emailId is new or updates database if emailId already exists
class CreateUser(Resource):
    #Function to resond to POST request on localhost:5000/CreateUser
    def post(self):
        try:
            # Parse the arguments
            parser = reqparse.RequestParser()
            parser.add_argument('username', type=str, help='USERNAME FOR USER')
            parser.add_argument('password', type=str, help='PASSWORD TO CREATE USER')
            parser.add_argument('emailid', type=str, help='EMAIL ID FOR USER')
            parser.add_argument('phonenumber', type=str, help='PHONE NUMBER FOR USER')
            args = parser.parse_args()

            #Accessing individual arguments
            _userName = args['username']
            _userEmail = args['emailid']
            _userPassword = args['password']
            _userPhoneNumber = args['phonenumber']

            #Establishing communication with the MySQL database
            conn = mysql.connect()
            cursor = conn.cursor()
            #Calling procedure spCreateUser on inputs:UserName,emailID,password and phoneNumber
            #spCreateUser returns the row where the given emailId exists or creates a row with the given inputs and returns null 
            cursor.callproc('spCreateUser',(_userName,_userEmail,_userPhoneNumber,_userPassword))
            data = cursor.fetchall()
            if len(data) is 0:
                conn.commit()
                return {'StatusCode':'200','Message': 'Operation Success'}
            else:
                return {'StatusCode':'1000','Message': str(data[0])}
        except Exception as e:
            return {'error': str(e)}
    #Functionto respond to GET requests sent to localhost:5000/CreateUser
    def get(self):
        return{'about':'getMethodCalled'}
#Class to return search results based on the given email Id if it exists on database or return null if it dosen't
#Takes input from the search division of 'formmod.html'
class Search(Resource):
    #Functionto respond to POST requests sent to localhost:5000/Search
    def post(self):
        try:
            # Parse the arguments
            parser = reqparse.RequestParser()
            parser.add_argument('emailid', type=str, help='EMAIL ID FOR USER')
            args = parser.parse_args()
            
            #Access 'emailId' from form 
            _userEmail = args['emailid']
            
            #Establish communication to MySQL database
            conn = mysql.connect()
            cursor = conn.cursor()
            
            #Calling procedure spSearch on input:emailID
            #spSearch returns the row where the given emailId exists or returns null
            cursor.callproc('spSearch',(_userEmail,))
            #data stores the row
            data = cursor.fetchall()
            if (len(data)>0):
                return {'StatusCode':'200','Username': str(data[0][0]),'Email':str(data[0][1]),'PhoneNumber':str(data[0][2]),'Password':str(data[0][3]),'Time':str(data[0][4])}
            else:
                return {'StatusCode':'1000','Message': "not found"}
        except Exception as e:
            return {'error': str(e)}
    #Functionto respond to GET requests sent to localhost:5000/Search
    def get(self):
        return{'about':'getMethodCalled'}
#Class to delete row based on the given emailId if it exists or return null
#Takes input from the delete division of 'formmod.html'
class Delete(Resource):
    def post(self):
        try:
            # Parse the arguments
            parser = reqparse.RequestParser()
            parser.add_argument('emailid', type=str, help='EMAIL ID FOR USER')
            args = parser.parse_args()
            
            #Access 'emailId' from form 
            _userEmail = args['emailid']
            
            #Establish communication to MySQL database
            conn = mysql.connect()
            cursor = conn.cursor()
            
            #Calling procedure spSearch on input:emailID
            #spSearch returns the row where the given emailId exists or returns null
            cursor.callproc('spSearch',(_userEmail,))
            data = cursor.fetchall()
            if (len(data)>0):
                #Calling procedure spDelete on input:emailID
                #spDelete deletes the row where the given emailId exists or returns null
                cursor.callproc('spDelete',(_userEmail,))
                conn.commit()
                return {'Success':'Record Found and Deleted'}
            else:
                return{'about':'Record Not Found'}
        except Exception as e:
            return {'error': str(e)}
    #Function to respond to GET requests sent to localhost:5000/Delete
    def get(self):
        return{'about':'getMethodCalled'}
#Adding resources to api
api.add_resource(CreateUser, '/CreateUser')
api.add_resource(Search, '/Search')
api.add_resource(Delete,'/Delete')
api.add_resource(FormPage, '/')

#Run api
if __name__ == '__main__':
    app.run(debug=True)
