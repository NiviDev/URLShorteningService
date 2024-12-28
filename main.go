package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/teris-io/shortid"
)

//var db = make(map[string]string)

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Load HTML Templates
	r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Load the Main Page
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title":   "URL Shortening Service",
			"message": "Welcome",
		})
	})

	// Create Short URL
	r.POST("/shorten", func(ctx *gin.Context) {
		// Struct that contains the url and the shortened code
		var requestBody struct {
			URL string `form:"url" binding:"required,url"`
		}

		if err := ctx.ShouldBind(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		// Short id generator
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if handleError(ctx, 500, err) {
			return
		}

		shortCode, err := sid.Generate()
		if handleError(ctx, 500, err) {
			return
		}

		url := Shorten{
			URL:         requestBody.URL,
			ShortCode:   shortCode,
			AccessCount: 0,
		}

		if err := db.Create(&url).Error; err != nil {
			ctx.JSON(500, gin.H{"msg": "Failed to create URL", "error": err.Error()})
			return
		}

		ctx.JSON(201, gin.H{
			"id":        url.ID,
			"url":       url.URL,
			"shortCode": url.ShortCode,
			"createdAt": url.CreatedAt.Format(time.RFC3339),
			"updatedAt": url.UpdatedAt.Format(time.RFC3339),
		})

	})

	r.GET("/shorten/:shortCode", func(ctx *gin.Context) {
		shortCode := ctx.Param("shortCode")
		var url Shorten

		if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(404, gin.H{"msg": "Short URL not found"})
			} else {
				ctx.JSON(500, gin.H{"msg": "Failed to retrieve URL", "error": err.Error()})
			}
			return
		}

		ctx.JSON(200, gin.H{
			"id":        url.ID,
			"url":       url.URL,
			"shortCode": url.ShortCode,
			"createdAt": url.CreatedAt.Format(time.RFC3339),
			"updatedAt": url.UpdatedAt.Format(time.RFC3339),
		})

	})

	return r
}

func main() {
	// Connect to te DB
	dsn := "nivi:urlpassword@tcp(127.0.0.1:3306)/url_shortening_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Shorten{})
	r := setupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
