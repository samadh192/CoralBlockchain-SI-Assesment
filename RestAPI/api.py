from flask import Flask, render_template,make_response
from flask_restful import Resource, Api
from flask_restful import reqparse
from flaskext.mysql import MySQL

app = Flask(__name__)

mysql = MySQL()

# MySQL configurations
app.config['MYSQL_DATABASE_USER'] = 'dummyUser'
app.config['MYSQL_DATABASE_PASSWORD'] = 'dummyUser01'
app.config['MYSQL_DATABASE_DB'] = 'db_intern'
app.config['MYSQL_DATABASE_HOST'] = 'db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com'


mysql.init_app(app)
api = Api(app)

class FormPage(Resource):
    def __init__(self):
        pass
    def get(self):
        headers ={'Content-Type':'text/html'}
        return make_response(render_template('form.html'),200,headers)

class CreateUser(Resource):
    def post(self):
        try:
            # Parse the arguments
            parser = reqparse.RequestParser()
            parser.add_argument('username', type=str, help='USERNAME FOR USER')
            parser.add_argument('password', type=str, help='PASSWORD TO CREATE USER')
            parser.add_argument('emailid', type=str, help='EMAIL ID FOR USER')
            parser.add_argument('phonenumber', type=str, help='PHONE NUMBER FOR USER')
            args = parser.parse_args()

            _userName = args['username']
            _userEmail = args['emailid']
            _userPassword = args['password']
            _userPhoneNumber = args['phonenumber']

            conn = mysql.connect()
            cursor = conn.cursor()
            cursor.callproc('spCreateUser',(_userName,_userEmail,_userPhoneNumber,_userPassword))
            data = cursor.fetchall()
            if len(data) is 0:
                conn.commit()
                return {'StatusCode':'200','Message': 'Operation Success'}
            else:
                return {'StatusCode':'1000','Message': str(data[0])}
        except Exception as e:
            return {'error': str(e)}
    def get(self):
        return{'about':'getMethodCalled'}
class Search(Resource):
    def post(self):
        try:
            # Parse the arguments
            parser = reqparse.RequestParser()
            parser.add_argument('emailid', type=str, help='EMAIL ID FOR USER')
            args = parser.parse_args()

            _userEmail = args['emailid']
            print(_userEmail)
            conn = mysql.connect()
            #print("1")
            cursor = conn.cursor()
            #print("2")
            cursor.callproc('spSearch',(_userEmail,))
            #print("3")
            data = cursor.fetchall()
            #print(str(data[0][0]))
            if (len(data)>0):
                return {'StatusCode':'200','Username': str(data[0][0]),'Email':str(data[0][1]),'PhoneNumber':str(data[0][2]),'Password':str(data[0][3]),'Time':str(data[0][4])}
            else:
                return {'StatusCode':'1000','Message': "not found"}
        except Exception as e:
            return {'error': str(e)}
    def get(self):
        return{'about':'getMethodCalled'}
class Test(Resource):
    def post(self):
        return{'about':'submit button clicked!'}
api.add_resource(CreateUser, '/CreateUser')
api.add_resource(Search, '/Search')
api.add_resource(Test, '/Test')
api.add_resource(FormPage, '/')

if __name__ == '__main__':
    app.run(debug=True)
