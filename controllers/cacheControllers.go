package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/harshsri28/apica/module"
    "net/http"
    "time"
    "log"
)

// Cache handler instance
var lruCache *cache.LRUCache

// InitializeCache sets up the cache and the expiration routine
func InitializeCache(cacheInstance *cache.LRUCache) {
    lruCache = cacheInstance
    lruCache.StartExpirationRoutine()
}


// return all the cache list
func GetAllCache(c *gin.Context) {
    allItems := lruCache.GetAll()
    c.JSON(http.StatusOK, gin.H{"items": allItems})
}

// GetCache retrieves an item from the cache
func GetCache(c *gin.Context) {
    key := c.Param("key")
    if value, found := lruCache.Get(key); found {
        c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
    }
}

// SetCache adds or updates an item in the cache
func SetCache(c *gin.Context) {
    var requestBody struct {
        Key       string `json:"key" binding:"required"`
        Value     string `json:"value" binding:"required"`
        Duration  int    `json:"duration" binding:"required"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }

    duration := time.Duration(requestBody.Duration) * time.Second
    lruCache.Set(requestBody.Key, requestBody.Value, duration)
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// DeleteCache removes an item from the cache
func DeleteCache(c *gin.Context) {
    key := c.Param("key")
    lruCache.Delete(key)
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
