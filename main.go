package main

import (
    "time"
    "github.com/gin-gonic/gin"
)

var startTime = time.Now()

func main() {
    r := gin.Default()

    // Set the timezone 
    location, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        location = time.UTC 
    }

    r.GET("/health", func(c *gin.Context) {
        uptime := time.Since(startTime).Seconds() 
        currentTime := time.Now().In(location)    

        c.JSON(200, gin.H{
            "nama":      "Faith Mary Sani",
            "nrp":       "5025231103",
            "status":    "UP",
            "timestamp": currentTime.Format(time.RFC3339), 
            "uptime":    uptime,                           // Uptime in seconds
        })
    })

    r.Run(":8080")
}
