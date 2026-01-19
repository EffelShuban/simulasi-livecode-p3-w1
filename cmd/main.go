package main

import (
	"context"
	"fmt"
	"net/http"
	"sim-livecodep3w1/config"
	"sim-livecodep3w1/handler"
	"sim-livecodep3w1/internal/repository"
	"sim-livecodep3w1/internal/service"
	"time"

	"github.com/labstack/echo"
	"github.com/robfig/cron/v3"
)

func main() {
	db := config.InitDB()
	
	repo := repository.NewMongodbGameRepository(db)

	svc := service.NewGameService(repo)

	handler := handler.NewGameHandler(svc)

	c := cron.New(cron.WithSeconds())

	e := echo.New()
	g := e.Group("/api")
	g.GET("/games", handler.Find)
	g.GET("/games/:id", handler.FindByID)
	g.POST("/games", handler.Create)
	g.PUT("/games/:id", handler.Update)
	g.DELETE("/games/:id", handler.Delete)
	_, err := c.AddFunc("0 * * * * *", func() { 
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := svc.UpdateVersion(ctx); err != nil {
			fmt.Printf("failed to update game version: %v\n", err)
		} else {
			fmt.Println("ran version update")
		}
	})
	if err != nil{
		panic(err)
	}
	c.Start()

	g.GET("/test", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{"message": "test success"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
