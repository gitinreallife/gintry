package view

import (
   "github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	files := v1.Group("/files")
	
	type File struct{
		Name string `uri:"name" binding:"required"`
	}

	files.POST("/", func(c *gin.Context) {
	     	// Controller code's goes here
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		n, err  := controller.Upload(file)
		if err != nil {
		   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		   return
		}
		c.JSON(http.StatusOK, gin.H{
		   "success": "Uploaded successfully",
		   "name": fmt.Sprintf("%s", n),
		})
	})

	
	files.GET("/:name/", func(c *gin.Context) {
		var file File
		if err := c.ShouldBindUri(&file); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		m, cn, err := controller.Download(f.Name)
		if err != nil {
		   c.JSON(http.StatusNotFound, gin.H{"error": err})
		   return
		}
		c.Header("Content-Disposition", "attachment; filename="+n)
		c.Data(http.StatusOK, m, cn)
	})

	
}
