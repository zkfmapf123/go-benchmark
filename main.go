package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/zkfmapf123/go-buffer/src"
)

var (
	APP_NAME = "go buffer"
	VERSION  = "1.0.0"
	PORT     = "3000"
)

func main() {
	// pprof 서버 시작
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  APP_NAME,
		AppName:       fmt.Sprintf("%s-%s", APP_NAME, VERSION),
	})

	app.Use(logger.New())

	// q := src.NewQueue(100)
	q := src.NewGoodQueue(100, 10)
	q.Close()

	// ping
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	// job
	app.Post("/job", func(c *fiber.Ctx) error {
		uid := uuid.New()
		job := src.Job{
			Idx: uid.String(),
		}

		q.Producer(job)

		return c.SendStatus(200)
	})

	// 메모리 상태 확인
	app.Get("/memstats", func(c *fiber.Ctx) error {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		return c.JSON(fiber.Map{
			"alloc":       m.Alloc,
			"total_alloc": m.TotalAlloc,
			"sys":         m.Sys,
			"num_gc":      m.NumGC,
		})
	})

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}
