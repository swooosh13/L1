package main

import "fmt"

type client struct {}

func (c *client) connectToDb(db databaseWP) {
	fmt.Println("Connecting to db start")
	db.ConnectWithPort()
}

type databaseWP interface {
	ConnectWithPort()
}

type mysql struct {}

func (d *mysql) ConnectWithPort() {
	fmt.Println("connected with mysql db on port 5432")
}


type postgres struct {}

func (d *postgres) Connect() {
	fmt.Println("connect with postgres")
}

// реализует тот же интерфейс который ожидает клиент
// (databaseWP)
type postgresAdapter struct {
	postgresDb *postgres
}

func (p *postgresAdapter) ConnectWithPort() {
	fmt.Println("adapting postgres connection with port 5432")
	p.postgresDb.Connect()
}

func main() {
	cl := client{}

	mySql := &mysql{}			// удовлетворяет условию для клиента
	pgSql := &postgres{}	// неудовлетворяет условию для клиента
	pgAdapter := &postgresAdapter{
		postgresDb: pgSql,
	}

	cl.connectToDb(mySql)
	//cl.connectToDb(postgreSql) // не удовлетворит
	cl.connectToDb(pgAdapter) 	 // удовлетворит
}
