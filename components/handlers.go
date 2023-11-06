package components

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLinks(c *gin.Context) {
	rows, err := db.Query("SELECT description, url FROM links")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var links []LinkMetadata
	for rows.Next() {
		var description, url string
		err := rows.Scan(&description, &url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		links = append(links, LinkMetadata{Description: description, Address: link{URL: url}})
	}

	c.JSON(http.StatusOK, links)
}

func postLinks(c *gin.Context) {
	var newLink LinkMetadata
	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	_, err := db.Exec("INSERT INTO links (description, url) VALUES (:description, :url)", sql.Named("description", newLink.Description), sql.Named("url", newLink.Address.URL))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, newLink)
}

func getLinksHTML(c *gin.Context) {
	rows, err := db.Query("SELECT description, url FROM links")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var links []LinkMetadata
	for rows.Next() {
		var description, url string
		err := rows.Scan(&description, &url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		links = append(links, LinkMetadata{Description: description, Address: link{URL: url}})
	}

	data := struct {
		Links []LinkMetadata
	}{
		Links: links,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "text/html")
	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}

func refreshLinksHTML(c *gin.Context) {
	getLinksHTML(c)
}

func getLinksByID(c *gin.Context) {
	id := c.Param("id")

	rows, err := db.Query("SELECT description, url FROM links WHERE description = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var linkData LinkMetadata
	if rows.Next() {
		var description, url string
		err := rows.Scan(&description, &url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		linkData = LinkMetadata{Description: description, Address: link{URL: url}}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, linkData)
}
