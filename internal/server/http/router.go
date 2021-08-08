package internalhttp

import (
	"github.com/gin-gonic/gin"

	"catalog/internal/app"
)

func NewGinRouter(app *app.App) *gin.Engine {
	router := gin.Default()

	router.GET("/buildings", func(c *gin.Context) {GetBuildingsHandler(c, app)})
	router.GET("/building/:id", func(c *gin.Context) {GetBuildingHandler(c, app)})
	router.PUT("/building", func(c *gin.Context) {UpdateBuildingHandler(c, app)})
	router.POST("/building", func(c *gin.Context) {CreateBuildingHandler(c, app)})
	router.DELETE("/building/:id", func(c *gin.Context) {DeleteBuildingHandler(c, app)})

	router.GET("/categories", func(c *gin.Context) {GetCategoriesHandler(c, app)})
	router.GET("/category/:id", func(c *gin.Context) {GetCategoryHandler(c, app)})
	router.PUT("/category", func(c *gin.Context) {UpdateCategoryHandler(c, app)})
	router.POST("/category", func(c *gin.Context) {CreateCategoryHandler(c, app)})
	router.DELETE("/category/:id", func(c *gin.Context) {DeleteCategoryHandler(c, app)})

	router.GET("/organizations", func(c *gin.Context) {GetOrganizationsHandler(c, app)})
	router.GET("/organization/:id", func(c *gin.Context) {GetOrganizationHandler(c, app)})
	router.PUT("/organization", func(c *gin.Context) {UpdateOrganizationHandler(c, app)})
	router.POST("/organization", func(c *gin.Context) {CreateOrganizationHandler(c, app)})
	router.DELETE("/organization/:id", func(c *gin.Context) {DeleteOrganizationHandler(c, app)})

	return router
}
