package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "fmt"
)

var startTime = time.Now()

func main() {
    r := gin.Default()

    // Set the timezone 
    location, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        location = time.UTC 
    }

    r.GET("/health-check", func(c *gin.Context) {
        uptime := time.Since(startTime)

        days := int(uptime.Hours()) / 24
        hours := int(uptime.Hours()) % 24
        minutes := int(uptime.Minutes()) % 60
        seconds := int(uptime.Seconds()) % 60

        formattedUptime := fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
        currentTime := time.Now().In(location)    

        c.JSON(200, gin.H{
            "nama":      "Faith Mary Sani",
            "nrp":       "5025231103",
            "status":    "UP",
            "timestamp": currentTime.Format(time.RFC3339), 
            "uptime":    formattedUptime,           
        })
    })

    r.Run(":8080")
}
