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

}

func TestIsInTable(t *testing.T) {

}

func TestInsertHashtag(t *testing.T) {

}

func TestSelectRandomHashtag(t *testing.T) {

}

func TestPurgeDB(t *testing.T) {

}
