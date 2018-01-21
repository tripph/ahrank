package main

import (
	"github.com/labstack/echo"
	"github.com/thilltbc/ahrank/backend/auctions"
	"os"
)

func main() {
	e := echo.New()
	auctions.INIT()
	port := os.Getenv("PORT")
	if port == "" {
		port = "1232"
	}

	e.Static("/", "backend/static_assets")
	e.File("/", "backend/static_assets/index.html")

	e.GET("/auction-count-scores", func(c echo.Context) error {
		return c.JSON(200, auctions.GetAuctionCountRanking())
	})
	e.Logger.Fatal(e.Start(":"+port))
}