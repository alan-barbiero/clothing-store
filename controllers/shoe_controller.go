package controllers

import (
	"clothing-store/config"
	"clothing-store/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateShoe(c *gin.Context) {
	var shoe models.Shoe
	if err := c.ShouldBindJSON(&shoe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO shoes (name, size, quantity, price)
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := config.GetDB().QueryRow(
		context.Background(),
		query,
		shoe.Name, shoe.Size, shoe.Quantity, shoe.Price,
	).Scan(&shoe.ID, &shoe.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not insert shoe"})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

// READ
func GetAllShoes(c *gin.Context) {
	query := `SELECT id, name, size, quantity, price, created_at FROM shoes`
	rows, err := config.GetDB().Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch shoes"})
		return
	}
	defer rows.Close()

	var shoes []models.Shoe
	for rows.Next() {
		var s models.Shoe
		err := rows.Scan(&s.ID, &s.Name, &s.Size, &s.Quantity, &s.Price, &s.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning row"})
			return
		}
		shoes = append(shoes, s)
	}

	c.JSON(http.StatusOK, shoes)
}

// UPDATE
func UpdateShoe(c *gin.Context) {
	id := c.Param("id")
	var shoe models.Shoe

	if err := c.ShouldBindJSON(&shoe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE shoes SET name=$1, size=$2, quantity=$4, price=$5 WHERE id=$6`

	cmdTag, err := config.GetDB().Exec(context.Background(), query,
		shoe.Name, shoe.Size, shoe.Quantity, shoe.Price, id,
	)

	if err != nil || cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update shoe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "shoe updated successfully"})
}

// DELETE
func DeleteShoe(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM shoes WHERE id=$1`

	cmdTag, err := config.GetDB().Exec(context.Background(), query, id)
	if err != nil || cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete shoe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "shoe deleted successfully"})
}
