package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *sql.DB
)

type TodoList struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed int8   `json:"completed"`
}

//	func getAlbums(c *gin.Context) {
//		c.IndentedJSON(http.StatusOK, albums)
//	}
func getTodoLists(c *gin.Context) {
	var todoList []TodoList
	rows, err := db.Query("SELECT id, title, completed FROM TodoList")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item TodoList
		if err := rows.Scan(&item.ID, &item.Title, &item.Completed); err != nil {
			panic(err.Error())
		}

		fmt.Printf("User ID: %d, Title: %s, Completed:%d\n", item.ID, item.Title, item.Completed)
		todoList = append(todoList, item)
	}
	fmt.Println(todoList)
	c.IndentedJSON(http.StatusOK, todoList)
}
func insertData(c *gin.Context) {
	var todoList TodoList
	if err := c.BindJSON(&todoList); err != nil {
		return
	}
	fmt.Println(todoList.Completed)
	fmt.Println(todoList.Title)
	query := "INSERT INTO TodoList (title, completed) VALUES (?, ?)"
	_, err := db.Exec(query, todoList.Title, todoList.Completed)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data inserted successfully!")
	c.IndentedJSON(http.StatusOK, todoList)
}
func editData(c *gin.Context) {
	var todoList TodoList
	fmt.Println(c.Accepted)

	if err := c.BindJSON(&todoList); err != nil {
		return
	}
	fmt.Println(todoList.ID)
	fmt.Println(todoList.Title)
	fmt.Println(todoList.Completed)
	query := "UPDATE TodoList SET title = ?, completed = ? WHERE id = ?"
	_, err := db.Exec(query, todoList.Title, todoList.Completed, todoList.ID)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data updated successfully!")
	c.IndentedJSON(http.StatusOK, todoList)
}
func deleteData(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	query := "DELETE FROM TodoList WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	fmt.Println(newAlbum.ID)
// 	if newAlbum.ID == "" {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please provide Id"})
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "golang:Messi8189@@tcp(mysql-golang.alwaysdata.net:3306)/golang_sql")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	fmt.Println("Connected to MySQL successfully!")
	return db

}

func Create(db *sql.DB) {

	_, err := db.Exec(`
        CREATE TABLE TodoList (
			id INT PRIMARY KEY AUTO_INCREMENT,  
			title VARCHAR(40) NOT NULL,  
			completed BOOLEAN  
        )
    `)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Table 'Reminders' created successfully!")
}

func main() {
	// Connect()
	db = Connect()
	router := gin.Default()
	router.GET("/todoList", getTodoLists)
	router.POST("/insertData", insertData)
	router.POST("/updateData", editData)
	router.GET("/deleteData/:id", deleteData)
	// router.Run("192.168.1.4:8000")
	router.Run("localhost:8080")
	// router.Run("golang.alwaysdata.net")
}
