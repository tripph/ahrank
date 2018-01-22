package main

import (
	"github.com/labstack/echo"
	"github.com/thilltbc/ahrank/backend/auctions"
	"github.com/thilltbc/ahrank/backend/realms"

	"os"
	"github.com/labstack/echo/middleware"
	"fmt"
)

func main() {
	e := echo.New()
	//auctions.INIT()
	realmList := realms.GetRealmNames()
	fmt.Printf("Realms: %v\n", realmList)
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

	e.GET("/auction-count-scores", func(c echo.Context) error {
		return c.JSON(200, auctions.GetAuctionCountRanking())
	})
	e.Logger.Fatal(e.Start(":"+port))
}