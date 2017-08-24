package main

import (
	"log"
	"net/http"
	"todo/controllers"
)

func main() {
	// 静态资源服务
	http.Handle("/public/", http.FileServer(http.Dir("./")))

	// 路由
	http.HandleFunc("/new", controllers.NewTodo)
	http.HandleFunc("/edit", controllers.EditTodo)
	http.HandleFunc("/finish", controllers.FinishTodo)
	http.HandleFunc("/delete", controllers.DeleteTodo)
	http.HandleFunc("/oss", controllers.HelloOss)
	http.HandleFunc("/", controllers.Index)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServer error:", err)
	}
}
