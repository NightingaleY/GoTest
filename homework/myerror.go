package homework

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

//import "errors"

func mockError() error {
	//报了某个错误
	return nil
}

var RowsNotFound = errors.New("Rows Not Found")
var myquery = ""

func myFunc() error {
	err := myDao(myquery)
	if errors.As(err, RowsNotFound) {
		fmt.Println("Results Not found")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "System error!")
	}
	return nil
}

func myDao(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		return errors.Wrapf(RowsNotFound, fmt.Sprintf("rows not found, query is : %s", query))
	}
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("another error,query is : %s", query))
	}
	return nil
}
