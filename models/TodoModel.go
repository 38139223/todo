package models

import (
	"database/sql"
//	_ "github.com/mattn/go-mysql"
)

// 为开发方便，使用sqlite数据库
type Todo struct {
	Id     int64
	Title  string
	Finish bool
}


func initConnPool() (*sql.DB,error) {
	Db := &MySQLClient{Host:"localhost",User:"office",Pwd:"baidong",DB:"todo",Port:3306,MaxOpen:300,MaxIdle:200}
	Db.Init()
	return Db.pool,nil
}
func InsertTodo(title string) (int64, error) {
	// 数据库path是相对路径（相对于main.go）
/*	db, err := sql.Open("mysql", "office:baidong@/todo")
	// 函数代码执行完后关闭数据库，这是个好习惯，我爱defer
	defer db.Close()
	if err != nil {
		return -1, err
	}*/
	//
/*	Db := &MySQLClient{Host:"localhost",User:"office",Pwd:"baidong",DB:"todo",Port:3306,MaxOpen:300,MaxIdle:200}
	Db.Init()*/
	db,err :=initConnPool()
	stmt, err := db.Prepare("INSERT INTO todo(title, finish) VALUES(?, ?)")
	defer stmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(title, false)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func QueryAll() ([]Todo, error) {
	db,err :=initConnPool()
	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var todos []Todo

	for rows.Next() {
		var id int64
		var title string
		var finish bool
		err = rows.Scan(&id, &title, &finish)
		if err != nil {
			return nil, err
		}
		todo := Todo{id, title, finish}
		todos = append(todos, todo)
	}
	return todos, nil
}

func FinishTodo(todoId int64, finish bool) (int64, error) {
	db,err :=initConnPool()
	stmt, err := db.Prepare("UPDATE todo SET finish=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return 0, nil
	}

	res, err := stmt.Exec(finish, todoId)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}

func DeleteTodo(todoId int64) (int64, error) {
	db,err := initConnPool()
	stmt, err := db.Prepare("DELETE FROM todo WHERE id=?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(todoId)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}

func GetTodoTitle(todoId int64) (string, error) {
	db,_ := initConnPool()
	// 只查询一行数据
	row := db.QueryRow("SELECT title FROM todo WHERE id=?", todoId)
	var title string
	e := row.Scan(&title)
	if e != nil {
		return "", e
	}
	return title, nil
}

func EditTodo(id int64, title string) (int64, error) {
	db,err := initConnPool()
	stmt, err := db.Prepare("UPDATE todo SET title=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return 0, nil
	}

	res, err := stmt.Exec(title, id)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}