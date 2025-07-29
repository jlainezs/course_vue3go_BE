package data

import "testing"

func Test_Ping(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("failed to ping database")
	}
}

func TestBook_GetAll(t *testing.T) {
	all, err := models.Book.GetAll()
	if err != nil {
		t.Error("failed to get all books", err)
	}

	if len(all) != 1 {
		t.Error("failed to get the correct number of books")
	}
}

func TestBook_GetOneById(t *testing.T) {
	book, err := models.Book.GetOneById(1)
	if err != nil {
		t.Error("failed to get book by id", err)
	}

	if book.Title != "My Book" {
		t.Errorf("expected title to be My Book, get %s", book.Title)
	}
}

func TestBook_GetOneBySlug(t *testing.T) {
	book, err := models.Book.GetOneBySlug("my-book")
	if err != nil {
		t.Error("failed to get book by slug", err)
	}

	if book.Title != "My Book" {
		t.Errorf("expected title to be My Book, get %s", book.Title)
	}

	_, err = models.Book.GetOneBySlug("bad-slug")
	if err == nil {
		t.Error("expected error when getting book by bad slug")
	}
}
