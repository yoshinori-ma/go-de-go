package main

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type Task struct {
	gorm.Model
	Title  string
	StoryPoint uint
	Story Story
	StoryID uint
}
type Story struct {
	gorm.Model
	Title  string
}

func main() {
	// r := gin.Default()
	// api := r.Group("/api")
	// v4 := api.Group("/v4")
	// v4.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Task{}, &Story{})

	story := Story{Title: "Story 1"}
	result := db.Create(&story)
	fmt.Println(story.ID)
	fmt.Println(result)
	task := Task{Title: "Task 1", StoryPoint: 1, StoryID: story.ID}
	result = db.Create(&task)

	var tasks []Task
	db.Find(&tasks)

	db.Delete(&tasks)

	for _, task := range tasks {
		fmt.Println(task.ID)
	}
	
	// db.Create(&Task{Title: "Task 1", StoryPoint: 1, StoryID: 1})
	// // Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
