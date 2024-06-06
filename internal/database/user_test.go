package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	username := "testuser"
	email := "testuser@example.com"
	hashedPassword := "hashedpassword"

	mock.ExpectExec(insertUserQuery).WithArgs(username, email, hashedPassword).WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := CreateUser(db, username, email, hashedPassword)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if id != 1 {
		t.Errorf("expected id to be 1, but got %d", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestCreateUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	username := "testuser"
	email := "testuser@example.com"
	hashedPassword := "hashedpassword"

	mock.ExpectExec(insertUserQuery).
		WithArgs(username, email, hashedPassword).
		WillReturnError(fmt.Errorf("insert failed"))

	_, err = CreateUser(db, username, email, hashedPassword)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "testuser@example.com"

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
		AddRow(1, "testuser", email, "hashedpassword")

	mock.ExpectQuery(getUserQuery).
		WithArgs(email).
		WillReturnRows(rows)

	user, err := GetUserByEmail(db, email)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if user == nil {
		t.Fatalf("expected a user, but got nil")
	}

	if user.Id != 1 || user.Username != "testuser" || user.Email != email || user.Password != "hashedpassword" {
		t.Errorf("unexpected user data: %+v", user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "nonexistentuser@example.com"

	mock.ExpectQuery(getUserQuery).
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	user, err := GetUserByEmail(db, email)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if user != nil {
		t.Errorf("expected no user, but got %+v", user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestEmailExists(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "testuser@example.com"

	// Test case: email exists
	mock.ExpectQuery(emailExistsQuery).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := EmailExists(db, email)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if !exists {
		t.Errorf("expected email to exist, but it does not")
	}

	// Test case: email does not exist
	mock.ExpectQuery(emailExistsQuery).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	exists, err = EmailExists(db, email)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if exists {
		t.Errorf("expected email to not exist, but it does")
	}

	// Test case: query error
	mock.ExpectQuery(emailExistsQuery).
		WithArgs(email).
		WillReturnError(sql.ErrConnDone)

	_, err = EmailExists(db, email)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
