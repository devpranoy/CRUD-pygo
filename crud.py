import psycopg2
from flask import Flask, request
import requests
import json
app = Flask(__name__) #app initialisation

# CREATE
@app.route('/create', methods=['POST']) 
def create():
    name = request.json['name']
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql ="""insert into data values('%s')"""%name
    cur.execute(sql)
    conn.commit()
    conn.close()
    response = {'message':'Success'}
    return json.dumps(response)

# READ
@app.route('/read', methods=['GET']) 
def read():
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql = "select * from data"
    cur.execute(sql)
    data = cur.fetchone()
    conn.commit()
    conn.close()
    response = {'name':data[0]}
    return json.dumps(response)

# UPDATE
@app.route('/update', methods=['GET','POST']) 
def update():
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    
    conn.commit()
    conn.close()

# DELETE
@app.route('/delete', methods=['GET','POST']) 
def delete():
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    
    conn.commit()
    conn.close()


if __name__=='__main__':
	app.run(debug=True,host="0.0.0.0",port=8000)
