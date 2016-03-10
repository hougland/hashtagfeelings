package main

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestOpenDBIfClosed(t *testing.T) {
	var err error
	db, _, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := OpenDBIfClosed()

	if result != db {
		t.Errorf("OpenDBIfClosed() returned a different database")
	}
}

func TestViewRows(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var hashtag1, hashtag2 Hashtag
	hashtag1.Name = "happy"
	hashtag2.Name = "sad"

	var hashtags []Hashtag
	hashtags = append(hashtags, hashtag1)
	hashtags = append(hashtags, hashtag2)

	rows := sqlmock.NewRows([]string{"id", "hashtag", "sentiment", "created"}).AddRow(1, "happy", "positive", "2016-03-08 23:57:51.645176+00").AddRow(2, "sad", "negative", "2016-03-06 23:57:51.645176+00")

	mock.ExpectQuery("^SELECT (.+) FROM hashtags$").WillReturnRows(rows)

	result := ViewRows()
	if result[0].Name != hashtags[0].Name && result[1].Name != hashtags[1].Name {
		t.Fatalf("ViewRows() failed. Expected: %v. Got: %v", hashtags, result)
	}
}

func TestIsInTable(t *testing.T) {
	// var err error
	// var mock sqlmock.Sqlmock
	// db, mock, err = sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()
	//
	// var trend1, trend2 anaconda.Trend
	// trend1.Name = "happy"
	// trend2.Name = "neutral"
	//
	// row1 := sqlmock.NewRows([]string{"id"}).AddRow(1)
	// // // _ = sqlmock.NewRows([]string{"id"})
	// //
	// mock.ExpectQuery("^SELECT id FROM hashtags WHERE hashtag = (.+)$").WillReturnRows(row1)
	// // // mock.ExpectQuery("^SELECT id FROM hashtags WHERE hashtag = 'neutral'").WillReturnRows(row2)
	//
	// result := IsInTable(trend1)
	// if result != true {
	// 	t.Fatalf("IsInTable failing")
	// }
}

func TestInsertHashtag(t *testing.T) {

}

func TestSelectRandomHashtag(t *testing.T) {

}

func TestPurgeDB(t *testing.T) {

}
