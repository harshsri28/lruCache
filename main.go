package main

import (
    "log"
    "os"
    "strconv"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/harshsri28/apica/module"
    "github.com/harshsri28/apica/routes"
    "github.com/harshsri28/apica/helper"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using environment variables directly.")
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "3001"
    }

    cacheCapacityStr := os.Getenv("CACHE_CAPACITY")
    if cacheCapacityStr == "" {
        cacheCapacityStr = "10"
    }

    cacheCapacity, err := strconv.Atoi(cacheCapacityStr)
    if err != nil {
        log.Fatalf("Invalid CACHE_CAPACITY value: %v", err)
    }

    lruCache := cache.NewLRUCache(cacheCapacity)
    router := gin.New()
    router.Use(gin.Logger())

    // Configured CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Logging incoming requests
    router.Use(func(c *gin.Context) {
        log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
        c.Next()
    })

    routes.InitCacheRoutes(router, lruCache)

    // Added WebSocket routes
    router.GET("/ws", websocket.HandleConnections)
    go websocket.HandleMessages()

    router.Run(":" + port)
}
