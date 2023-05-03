package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// 网络编程：Go 语言RESTful API 服务

// Gin
func main() {

	//http.HandleFunc("/users", handleUsers)
	//http.ListenAndServe(":8080", nil)

	r := gin.Default()
	r.GET("/users", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.DELETE("/users/:id", deleteUser)
	r.PATCH("/users/:id", updateUserName)
	r.Run(":8080")

}

//func handleUsers(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "ID:1,Name:张三")
//	fmt.Fprintln(w, "ID:2,Name:李四")
//	fmt.Fprintln(w, "ID:3,Name:王五")
//}

//func handleUsers(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintln(w, "ID:1,Name:张三")
//		fmt.Fprintln(w, "ID:2,Name:李四")
//		fmt.Fprintln(w, "ID:3,Name:王五")
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprintln(w, "not found")
//	}
//}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}

}

func listUser(c *gin.Context) {
	c.JSON(200, users)
}

func getUser(c *gin.Context) {

	id := c.Param("id")
	var user User
	found := false

	//类似于数据库的SQL查询
	for _, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}
	if found {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})

	}

}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{ID: len(users) + 1, Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	} else {
		c.JSON(http.StatusOK, gin.H{

			"message": "请输入用户名称",
		})

	}

}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	i := -1
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
	}

	if i >= 0 {
		users = append(users[:i], users[i+1:]...)
		c.JSON(http.StatusNoContent, "")
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}

}

func updateUserName(c *gin.Context) {
	id := c.Param("id")
	i := -1
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
	}

	if i >= 0 {
		users[i].Name = c.DefaultPostForm("name", users[i].Name)
		c.JSON(http.StatusOK, users[i])
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

// 用户
type User struct {
	ID   int
	Name string
}
