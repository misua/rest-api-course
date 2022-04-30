package main

import (
	"context"
	"fmt"

	"github.com/misua/go-rest-api-course/internal/comment"
	"github.com/misua/go-rest-api-course/internal/db"
)

func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to db")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	cmtService := comment.NewService(db)
	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "9fbf72e8-7d13-42a0-af26-6a2818036f37",
			Slug:   "manual-test",
			Author: "Charles",
			Body:   "Hello World",
		},
	)

	fmt.Println(
		cmtService.GetComment(
			context.Background(),
			"5fde3067-16b5-480f-82a2-77920beb1eda",
		))

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }
	// fmt.Println("successfully connected and pinged db")
	return nil

}

func main() {
	fmt.Println("hakeng coins.ph")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
