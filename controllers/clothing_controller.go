package controllers

import (
	"bytes"
	"clothing-store/config"
	"clothing-store/models"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClothing(c *gin.Context) {
	var clothing models.Clothing

	// Log the raw request body
	body, err := c.GetRawData()
	if err != nil {
		fmt.Printf("Error reading request body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not read request body"})
		return
	}
	fmt.Printf("Raw request body: %s\n", string(body))

	// Reset the request body so it can be read again
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&clothing); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		fmt.Printf("Request body: %s\n", string(body))
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid JSON: %v", err)})
		return
	}

	fmt.Printf("Attempting to insert clothing: %+v\n", clothing)
	fmt.Printf("Price value type: %T, value: %v\n", clothing.Price, clothing.Price)

	// Get database connection
	db := config.GetDB()
	if db == nil {
		fmt.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
		return
	}

	// Test database connection
	err = db.Ping(context.Background())
	if err != nil {
		fmt.Printf("Database ping failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
		return
	}

	query := `INSERT INTO clothing (name, size, category, quantity, price)
			  VALUES($1, $2, $3, $4, $5) RETURNING id, created_at`

	fmt.Printf("Executing query: %s\n", query)
	fmt.Printf("With values: name=%s, size=%s, category=%s, quantity=%d, price=%f\n",
		clothing.Name, clothing.Size, clothing.Category, clothing.Quantity, clothing.Price)

	err = db.QueryRow(
		context.Background(),
		query,
		clothing.Name, clothing.Size, clothing.Category, clothing.Quantity, clothing.Price,
	).Scan(&clothing.ID, &clothing.CreatedAt)

	if err != nil {
		fmt.Printf("Error inserting clothing: %v\n", err)
		fmt.Printf("SQL State: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not insert clothing: %v", err)})
		return
	}

	fmt.Printf("Successfully inserted clothing with ID: %d\n", clothing.ID)
	c.JSON(http.StatusOK, clothing)
}

func GetAllClothing(c *gin.Context) {
	query := `SELECT id, name, size, category, quantity, price, created_at FROM clothing`
	rows, err := config.GetDB().Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch clothing"})
		return
	}
	defer rows.Close()

	var clothes []models.Clothing
	for rows.Next() {
		var item models.Clothing
		err := rows.Scan(&item.ID, &item.Name, &item.Size, &item.Category, &item.Quantity, &item.Price, &item.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning row"})
			return
		}
		clothes = append(clothes, item)
	}

	c.JSON(http.StatusOK, clothes)
}

func UpdateClothing(c *gin.Context) {
	id := c.Param("id")
	var clothing models.Clothing

	if err := c.ShouldBindJSON(&clothing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE clothing SET name=$1, size=$2, category=$3, quantity=$4, price=$5 WHERE id=$6`

	cmdTag, err := config.GetDB().Exec(context.Background(), query,
		clothing.Name, clothing.Size, clothing.Category, clothing.Quantity, clothing.Price, id)

	if err != nil || cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update clothing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "clothing updated successfully"})
}

func DeleteClothing(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM clothing WHERE id=$1`

	cmdTag, err := config.GetDB().Exec(context.Background(), query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete clothing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "clothing deleted successfully"})
}
