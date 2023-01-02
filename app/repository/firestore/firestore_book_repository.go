package repository

import (
	"context"
	"errors"
	"library/entity"
	"library/repository"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type rep struct {
	client *firestore.Client
}

func NewFireStoreBookRepository(project_id string) repository.BookRepository {

	client, _ := firestore.NewClient(context.Background(), project_id)

	return &rep{client}
}

func (r *rep) SelectById(id string) (*entity.Book, error) {

	doc, err := r.client.Collection("Books").Doc(id).Get(context.Background())

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, errors.New("Book not found")
		}
		return nil, err
	}

	var book entity.Book

	err = doc.DataTo(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *rep) Insert(b *entity.Book) (*entity.Book, error) {

	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	_, err := r.client.Collection("Books").Doc(b.ID).Create(context.Background(), b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (r *rep) SelectAll() ([]entity.Book, error) {

	docs, err := r.client.Collection("Books").Documents(context.Background()).GetAll()

	if err != nil {
		return nil, err
	}

	var output []entity.Book
	var book entity.Book

	for _, doc := range docs {

		err = doc.DataTo(&book)
		if err != nil {
			continue
		}

		book.ID = doc.Ref.ID
		output = append(output, book)
	}

	return output, nil
}

func (r *rep) Delete(id string) error {

	_, err := r.client.Collection("Books").Doc(id).Delete(context.Background())

	if err != nil && status.Code(err) == codes.NotFound {
		return errors.New("Book not found")
	}

	return err
}

func (r *rep) Update(b *entity.Book) error {

	_, err := r.client.Collection("Books").Doc(b.ID).Set(context.Background(), b)

	if err != nil && status.Code(err) == codes.NotFound {
		return errors.New("Book not found")
	}

	return err
}
