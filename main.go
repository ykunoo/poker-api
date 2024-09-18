package main

import (
	"log"
	"net/http"
	"poker-api/poker"

	"github.com/labstack/echo/v4"
)

type HandRequest struct {
	Hands []string `json:"hands"`
}

type HandResponse struct {
	Results []poker.HandResult `json:"results"`
	Errors  []poker.HandError  `json:"errors"`
}

func main() {

	poker.InitDB()

	e := echo.New()

	e.POST("/evaluate", func(c echo.Context) error {
		req := new(HandRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		}

		results, errors := poker.EvaluateHands(req.Hands)

		// データベースに結果を保存
		if err := poker.SaveResultsToDB(results, errors); err != nil {
			log.Printf("Failed to save results to DB: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save results to database"})
		}

		return c.JSON(http.StatusOK, HandResponse{
			Results: results,
			Errors:  errors,
		})
	})

	e.Start(":8080")
}
