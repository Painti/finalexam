package main

import (
	"fmt"

	"github.com/Painti/finalexam/database"
	"github.com/Painti/finalexam/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.New("customers")
	if err != nil {
		fmt.Println("DB Connection Fail")
		return
	}
	err = db.CreateCustomerTable()
	if err != nil {
		fmt.Println("Create customer table fail")
		return
	}
	defer db.Close()
	r := setupRouter(db)
	fmt.Println("Server started")
	r.Run(":2019")
}

func setupRouter(db *database.DBConnection) *gin.Engine {
	r := gin.Default()

	r.Use(service.AuthMiddleware)

	r.POST("/customers", func(c *gin.Context) {
		service.CreateCustomerHandler(c, db)
	})
	r.GET("/customers/:id", func(c *gin.Context) {
		service.GetCustomerByIDHandler(c, db)
	})
	r.GET("/customers", func(c *gin.Context) {
		service.GetCustomerHandler(c, db)
	})
	r.PUT("/customers/:id", func(c *gin.Context) {
		service.UpdateCustomerHandler(c, db)
	})
	r.DELETE("/customers/:id", func(c *gin.Context) {
		service.DeleteCustomerHandler(c, db)
	})

	return r
}
