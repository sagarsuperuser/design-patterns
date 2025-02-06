package main

import "fmt"

type Database interface {
	Connect()
	Query(query string)
	Close()
}

// MySQl - A legacy MySQL database driver with different methods.
type MySQl struct{}

func (m *MySQl) MySQLConnect() {
	fmt.Println("Connecting to MySQL Database...")
}

func (m *MySQl) MySQLQuery(sql string) {
	fmt.Println("Executing MySQL query: ", sql)
}

func (m *MySQl) MySQlClose() {
	fmt.Println("Closing MySQL database connection.")
}

// MySQLAdapter - Adapts MySQL to match the database interface.
type MySQLAdapter struct {
	mySQL *MySQl
}

func (m *MySQLAdapter) Connect() {
	m.mySQL.MySQLConnect()
}

func (m *MySQLAdapter) Query(query string) {
	m.mySQL.MySQLQuery(query)
}

func (m *MySQLAdapter) Close() {
	m.mySQL.MySQlClose()
}

// MongoDB -  A legacy MongoDB database driver with different API.
type MongoDB struct{}

func (m *MongoDB) MongoConnect() {
	fmt.Println("Connecting to MongoDB Database...")
}

func (m *MongoDB) MongoQuery(query string) {
	fmt.Println("Executing MongoDB query: ", query)
}

func (m *MongoDB) MongoClose() {
	fmt.Println("Closing MongoDB database connection.")
}

// MongoDBAdapter -  Adapts MongoDB to match the database interface.
type MongoDBAdapter struct {
	mongo *MongoDB
}

func (m *MongoDBAdapter) Connect() {
	m.mongo.MongoConnect()
}

func (m *MongoDBAdapter) Query(query string) {
	m.mongo.MongoQuery(query)
}

func (m *MongoDBAdapter) Close() {
	m.mongo.MongoClose()
}

func useDataBase(db Database) {
	db.Connect()
	db.Query("SELECT * FROM users")
	db.Close()
}

func main() {

	// Using MySQl Adapter
	fmt.Println("\nUsing MySQL Adapter: ")
	sqlAdapter := &MySQLAdapter{mySQL: &MySQl{}}
	useDataBase(sqlAdapter)

	// Using Mongo Adapter
	fmt.Println("\nUsing Mongo Adapter: ")
	mongoAdapter := &MongoDBAdapter{mongo: &MongoDB{}}
	useDataBase(mongoAdapter)

}
