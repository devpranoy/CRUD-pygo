import psycopg2
from flask import Flask, request
import json
app = Flask(__name__) #app initialisation

# CREATE
@app.route('/create', methods=['POST']) 
def create():
    name = request.json['name']
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql ="insert into data values('%s')"%name
    cur.execute(sql)
    conn.commit()
    conn.close()
    response = {'Message':'Success'}
    return json.dumps(response)

# READ
@app.route('/read', methods=['GET']) 
def read():
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql = "select * from data"
    cur.execute(sql)
    data = cur.fetchone()
    conn.close()
    response = {'Name':data[0]}
    return json.dumps(response)

# UPDATE
@app.route('/update', methods=['PUT']) 
def update():
    name = request.json['Name']
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql = "update data set name = '%s'"%name
    cur.execute(sql)
    conn.commit()
    conn.close()
    response = {'Message':'Success'}
    return json.dumps(response)

# DELETE
@app.route('/delete', methods=['DELETE']) 
def delete():
    name = request.json['Name']
    conn = psycopg2.connect(database="postgres", user = "postgres", password = "password", host = "127.0.0.1", port = "5432")
    cur = conn.cursor()
    sql = "delete from data where name='%s'"%name
    cur.execute(sql)
    conn.commit()
    conn.close()
    response = {'Message':'Success'}
    return json.dumps(response)


if __name__=='__main__':
	app.run(threaded=True,host="0.0.0.0",port=8000)
