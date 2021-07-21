package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func SelectData(db *sql.DB) (interface{}, error) {
	query := `select id, s_name from student`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, query)
	}
	var id string
	var name string
	studentList := make([]Student, 10)
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("id: %s, name: %s\n", id, name)
		studentList = append(studentList, Student{Id: id, Name: name})
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		rows.Close()
		return nil, err
	}

	return studentList, err
}

type Student struct {
	Id   string
	Name string
}

func daoMain() (interface{}, error) {
	db, _ := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/school")

	defer db.Close()
	// .....Logic层处理
	// 开始操作Dao层
	StudentList, err := SelectData(db)
	if err != nil {
		return nil, errors.Wrap(err, "SelectData error")
	}
	return StudentList, err
	// 将dao层错误返回抛像上层调用者
	// 原因: web开发上下文一般会有before_request 和 after_request, 为保证日志完整，dao层错误一般只网上抛, 在成功完成一次请求处理才进行log
}
