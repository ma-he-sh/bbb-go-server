package db

import (
	"fmt"

	env "github.com/devmarka/bbb-go-server/env"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	DBName  string = env.DBName()
	TBUsers string = "users"
	TBSetup string = "setup"
	TBEvent string = "events"
	Session *rethinkdb.Session
)

func init() {
	var err error
	Session, err = rethinkdb.Connect(rethinkdb.ConnectOpts{
		Address:  env.DBHost(),
		Database: env.DBName(),
		Username: env.DBUser(),
		Password: env.DBPass(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateTable() (rethinkdb.WriteResponse, error) {
	result, err := rethinkdb.DBCreate(DBName).RunWrite(Session)

	userkey := rethinkdb.TableCreateOpts{PrimaryKey: "email"}
	_, err = rethinkdb.DB(DBName).TableCreate(TBUsers, userkey).Run(Session)

	setupkey := rethinkdb.TableCreateOpts{PrimaryKey: "setup"}
	_, err = rethinkdb.DB(DBName).TableCreate(TBSetup, setupkey).Run(Session)

	eventkey := rethinkdb.TableCreateOpts{PrimaryKey: "event"}
	_, err = rethinkdb.DB(DBName).TableCreate(TBEvent, eventkey).Run(Session)

	if err != nil {
		return result, err
	}
	return result, err
}