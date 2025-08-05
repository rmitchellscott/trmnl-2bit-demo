package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var imagePaths []string

func findPNGImages(root string) ([]string, error) {
	var images []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".png" {
			images = append(images, path)
		}
		return nil
	})
	return images, err
}

func setupImages() error {
	images, err := findPNGImages("images")
	if err != nil {
		return err
	}

	log.Printf("Found %d PNG images", len(images))

	imagePaths = images

	if len(imagePaths) == 0 {
		return fmt.Errorf("no PNG images found in images/ directory")
	}

	for _, img := range imagePaths {
		log.Printf("Found image: %s", img)
	}

	return nil
}

func getRandomImage(c *gin.Context) {
	if len(imagePaths) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no images available"})
		return
	}

	refreshRate := 300
	if refreshParam := c.Query("refresh"); refreshParam != "" {
		if parsed, err := strconv.Atoi(refreshParam); err == nil && parsed > 0 {
			refreshRate = parsed
		}
	}

	randomIndex := rand.Intn(len(imagePaths))
	imagePath := imagePaths[randomIndex]

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	} else if c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	} else if c.Request.URL.Scheme != "" {
		scheme = c.Request.URL.Scheme
	}

	url := fmt.Sprintf("%s://%s/%s", scheme, c.Request.Host, imagePath)

	c.JSON(http.StatusOK, gin.H{
		"filename":     imagePath,
		"url":          url,
		"refresh_rate": refreshRate,
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := setupImages(); err != nil {
		log.Fatalf("Failed to setup images: %v", err)
	}

	r := gin.Default()

	r.GET("/", getRandomImage)
	r.Static("/images", "./images")

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}