package main

import (
	"context"
	"fmt"

	"github.com/misua/go-rest-api-course/internal/db"
)

func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatbase()
	if err != nil {
		fmt.Println("failed to connect to db")
		return err
	}
	if err :=db.MigrateDB();if err != nil{
		fmt.Println("failed to migrate db")
		return err
	}


	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }
	// fmt.Println("successfully connected and pinged db")
	return nil

}

func main() {
	fmt.Println("go Rest API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
