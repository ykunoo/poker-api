package main

import (
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
	e := echo.New()

	e.POST("/evaluate", func(c echo.Context) error {
		req := new(HandRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		}

		results, errors := poker.EvaluateHands(req.Hands)

		return c.JSON(http.StatusOK, HandResponse{
			Results: results,
			Errors:  errors,
		})
	})

	e.Start(":8080")
}
