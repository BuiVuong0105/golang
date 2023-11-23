package pattern

import (
	"fmt"
	"net/http"
)

type DB interface {
	connect(string) error
	save(string) error
	update(string) error
}

type Store struct {
}

// implicit Implement DB
func (s *Store) connect(value string) error {
	fmt.Println("Connect to database: ", value)
	return nil
}

// implicit Implement DB
func (s *Store) save(value string) error {
	fmt.Println("save to database: ", value)
	return nil
}

// implicit Implement DB
func (s *Store) update(value string) error {
	fmt.Println("update to database: ", value)
	return nil
}

type ExecuteFunction func(string)

// tham số là interface DB có khi gọi hàm có thể chấp nhận 1 con trỏ hoặc 1 giá trị của struct mà implement interface DB đó
// Trong Go, interface DB trong trường hợp của bạn không phải là một con trỏ, nhưng nó có thể giữ tham chiếu đến một giá trị cụ thể, bao gồm cả con trỏ.
func MyExecuteFunction(db DB) ExecuteFunction {
	return func(s string) {
		db.connect(s)
	}
}

func Execute(fn ExecuteFunction) {
	fn("http://localhost:9995")
}

type HttpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

func Handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func MakeHttpFunc(db DB, httpFunc HttpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := httpFunc(db, w, r); err != nil {
			fmt.Print("Handle Success")
		}
	}
}
