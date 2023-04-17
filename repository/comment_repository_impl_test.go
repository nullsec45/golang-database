package repository

import (
	"golang-database/database"

	"context"
	"fmt"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 23)

	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
