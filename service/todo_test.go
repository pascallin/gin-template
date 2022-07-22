package service

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestShouldCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// combine sqlmock connection to gorm
	mockMysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("error was not expected while mockMysqlDB stats: %s", err)
	}

	// store orignal function to another variable
	original := getMysqlDb
	// replace original function by mock function
	getMysqlDb = func() *gorm.DB { return mockMysqlDB }
	// restore package variable after test finished
	defer func() { getMysqlDb = original }()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `todos`").WithArgs("test", "test", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = CreateTodo("test", "test"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
