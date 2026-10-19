package controllers

import (
	"file-system/database"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 1. Get the uploaded file from the request
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please select a file",
		})
		return
	}

	// 2. Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	// 3. Read the actual file contents
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	// 4. Get file metadata
	filename := filepath.Base(fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")
	fileSize := int64(len(fileBytes))

	// 5. Insert the file into PostgreSQL
	var fileID int64

	err = database.DB.QueryRow(
		`INSERT INTO uploaded_files
        (filename, content_type, size_bytes, file_data)
        VALUES ($1, $2, $3, $4)
        RETURNING id`,
		filename,
		contentType,
		fileSize,
		fileBytes,
	).Scan(&fileID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file to database",
		})
		return
	}

	// 6. Return the result to the client
	c.JSON(http.StatusOK, gin.H{
		"message":      "File uploaded to PostgreSQL successfully",
		"id":           fileID,
		"filename":     filename,
		"content_type": contentType,
		"size":         fileSize,
	})
}

func GetUpload(c *gin.Context) {
	rows, err := database.DB.Query(
		`SELECT id, filename, content_type, size_bytes
		FROM uploaded_files`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get the file",
		})
		return
	}
	defer rows.Close()

	var files []gin.H

	for rows.Next() {
		var id int64
		var filename, contentType string
		var sizeBytes int64

		if err := rows.Scan(&id, &filename, &contentType, &sizeBytes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read files"})
			return
		}

		files = append(files, gin.H{
			"id":           id,
			"filename":     filename,
			"content_type": contentType,
			"size_bytes":   sizeBytes,
		})
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read files",
		})
		return
	}

	c.JSON(http.StatusOK, files)
}
