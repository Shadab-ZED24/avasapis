package db

import "gopkg.in/mgo.v2"

type DBConnection struct {
	session *mgo.Session
}

func NewConnection(host string) (conn *DBConnection) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

func (conn *DBConnection) Use(dbName string, collectionName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(collectionName)
}

func (conn *DBConnection) Close(){
	conn.session.Close()
	return
}