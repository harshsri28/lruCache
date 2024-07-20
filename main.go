package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
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

    lruCache := cache.NewLRUCache(10)
    router := gin.New()
    router.Use(gin.Logger())

    routes.InitCacheRoutes(router,lruCache)

    router.Run(":" + port)
}
