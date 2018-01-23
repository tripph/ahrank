package main

import (
	"github.com/labstack/echo"
	"github.com/thilltbc/ahrank/backend/auctions"

	"os"
	"github.com/labstack/echo/middleware"

	"github.com/thilltbc/ahrank/backend/realms"
)

func main() {
	e := echo.New()
	// go auctions.INIT()
	port := os.Getenv("PORT")
	if port == "" {
		port = "1232"
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Static("/", "backend/static_assets")
	e.File("/", "backend/static_assets/index.html")

	e.GET("/api/auction-count-scores", func(c echo.Context) error {
		return c.JSON(200, auctions.GetAuctionCountRanking())
	})
	e.GET("/api/realms/list", func(c echo.Context) error {
		return c.JSON(200, realms.GetRealmList())
	})
	e.Logger.Fatal(e.Start(":"+port))
}