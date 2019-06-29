package service

import (
	"net/http"
	"strconv"

	"github.com/Painti/finalexam/database"
	"github.com/Painti/finalexam/model"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetCustomerHandler(c *gin.Context, conn *database.DBConnection) {
	customers, err := model.GetAllCustomer(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func GetCustomerByIDHandler(c *gin.Context, conn *database.DBConnection) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	customer := model.Customer{ID: id}
	err = customer.GetData(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func CreateCustomerHandler(c *gin.Context, conn *database.DBConnection) {
	var customer model.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	err = customer.Create(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomerHandler(c *gin.Context, conn *database.DBConnection) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	var customer model.Customer
	err = c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	customer.ID = id
	err = customer.Save(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomerHandler(c *gin.Context, conn *database.DBConnection) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	customer := model.Customer{ID: id}
	err = customer.Delete(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
