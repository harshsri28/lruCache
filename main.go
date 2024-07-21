package main

import (
    "log"
    "os"
    "strconv"
    "time"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors" 
    "github.com/harshsri28/apica/module"
    "github.com/harshsri28/apica/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT environment variable not set")
    }

    cacheCapacityStr := os.Getenv("CACHE_CAPACITY")
    if cacheCapacityStr == "" {
        log.Fatal("CACHE_CAPACITY environment variable not set")
    }

    cacheCapacity, err := strconv.Atoi(cacheCapacityStr)
    if err != nil {
        log.Fatalf("Invalid CACHE_CAPACITY value: %v", err)
    }

    lruCache := cache.NewLRUCache(cacheCapacity)
    router := gin.New()
    router.Use(gin.Logger())

    // Configure CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Log incoming requests
    router.Use(func(c *gin.Context) {
        log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
        c.Next()
    })

    routes.InitCacheRoutes(router, lruCache)

    router.Run(":" + port)
}
