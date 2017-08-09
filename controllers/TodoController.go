package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"todo/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		showError(w, "异常", "非法请求，服务器无法响应")
	} else {
		if r.URL.Path == "/" {
			todos, err := models.QueryAll()
			if err != nil {
				showError(w, "异常", "查询异常")
				return
			}
			t, err := template.ParseFiles("views/index.html")
			if err != nil {
				showError(w, "异常", "页面渲染异常")
				return
			}
			data := make(map[string][]models.Todo)
			data["TodoList"] = todos
			t.Execute(w, data)
		} else {
			// 404页面，路由不到的都会到这里
			showError(w, "404", "页面不存在")
		}
	}
}

func NewTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, "异常", "非法请求")
	} else {
		title := r.FormValue("title")
		id, err := models.InsertTodo(title)

		if err != nil || id <= 0 {
			showError(w, "异常", "插入数据异常")
			return
		}
		// 重定向到主界面
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// 没有return，没有效果，重定向不过去
		return
	}
}

func FinishTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		showError(w, "异常", "非法请求")
	} else {
		// 获取表单参数，也可以这么写
		// r.ParseForm()
		// id := r.Form["id"]
		id := r.FormValue("id")
		finish := r.FormValue("finish")
		// FormValue取到的数据都为string类型，将id转为int64类型
		// strconv.ParseInt(id, 10, 64) 10意思为10进制，64意思为64位
		intId, _ := strconv.ParseInt(id, 10, 64)
		boolFinish, _ := strconv.ParseBool(finish)

		_, err := models.FinishTodo(intId, !boolFinish)
		if err != nil {
			showError(w, "异常", "完成Todo失败")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		showError(w, "异常", "非法请求")
	} else {
		id := r.FormValue("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		_, err := models.DeleteTodo(intId)
		if err != nil {
			showError(w, "异常", "删除失败")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 显示edit页面
		// 本可以将title内容提交至此，但url将会异常难看，还是根据id查询吧
		id := r.FormValue("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		title, err := models.GetTodoTitle(intId)
		if err != nil {
			showError(w, "异常", "查询Todo内容失败")
			return
		}
		t, _ := template.ParseFiles("views/edit.html")
		data := make(map[string]string)
		data["Id"] = id
		data["Title"] = title
		t.Execute(w, data)

	} else if r.Method == "POST" {
		// edit后的数据post提交至此处
		id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		title := r.FormValue("title")
		res, err := models.EditTodo(id, title)
		if err != nil || res <= 0 {
			showError(w, "异常", "修改失败")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
}

// 错误处理
func showError(w http.ResponseWriter, title string, message string) {
	t, _ := template.ParseFiles("views/error.html")
	data := make(map[string]string)
	data["title"] = title
	data["message"] = message
	t.Execute(w, data)
}
//Oss