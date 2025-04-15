package controllers

import (
	"clothing-store/config"
	"clothing-store/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClothing(c *gin.Context) {
	var clothing models.Clothing
	if err := c.ShouldBindJSON(&clothing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO clothing (name, size, category, quantity, price)
			  VALUES($1, $2, $3, $4, $5) RETURNING id, created_at`

	err := config.GetDB().QueryRow(
		context.Background(),
		query,
		clothing.Name, clothing.Size, clothing.Category, clothing.Quantity, clothing.Price,
	).Scan(&clothing.ID, &clothing.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not insert clothing"})
		return
	}

	c.JSON(http.StatusOK, clothing)
}

func GetAllClothing(c *gin.Context) {
	query := `SELECT id, name, size, category, quantity, price, create_at FROM clothing`
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

	if err := c.ShouldBindJSON(clothing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE clothing SET name=$1, size=$2, color=$3, category=$4, quantity=$5, price=$6 WHERE id=$7`

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
