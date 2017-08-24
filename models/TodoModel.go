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
	Img_url string
}


func initConnPool() (*sql.DB,error) {
	Db := &MySQLClient{Host:"localhost",User:"office",Pwd:"baidong",DB:"todo",Port:3306,MaxOpen:300,MaxIdle:200}
	Db.Init()
	return Db.pool,nil
}
func InsertTodo(title ,img_url string) (int64, error) {
	db,err :=initConnPool()
	stmt, err := db.Prepare("INSERT INTO todo(title,img_url,finish) VALUES(?,?,?)")
	defer stmt.Close()
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(title,img_url,false)
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
		var img_url  string
		err = rows.Scan(&id, &title, &finish,&img_url)
		if err != nil {
			return nil, err
		}
		todo := Todo{id, title, finish,img_url}
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

func GetTodoTitle(todoId int64) (string,string, error) {
	db,_ := initConnPool()
	// 只查询一行数据
	var title string
	var img_url string
	err := db.QueryRow("SELECT title,img_url FROM todo WHERE id=?", todoId).Scan(&title,&img_url)
	if err != nil {
		return "","", err
	}
	return title,img_url, nil
}

func EditTodo(title string,img_url string,id int64) (int64, error) {
	db,err := initConnPool()
	stmt, err := db.Prepare("UPDATE todo SET title=?,img_url=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return 0, nil
	}

	res, err := stmt.Exec(title,img_url,id)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return affect, nil
}