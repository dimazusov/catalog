package http

import (
	"github.com/gin-gonic/gin"

	"catalog/internal/app"
	"catalog/internal/server/http/handlers/building"
	"catalog/internal/server/http/handlers/category"
	"catalog/internal/server/http/handlers/organization"
)

// @title Swagger API
// @version 1.0
// @description Catalog api
func NewGinRouter(app *app.App) *gin.Engine {
	router := gin.Default()

	v1Group := router.Group("/api/v1")

	v1Group.GET("/buildings", func(c *gin.Context) { building.GetBuildingsHandler(c, app) })
	v1Group.GET("/building/:id", func(c *gin.Context) { building.GetBuildingHandler(c, app) })
	v1Group.PUT("/building", func(c *gin.Context) { building.UpdateBuildingHandler(c, app) })
	v1Group.POST("/building", func(c *gin.Context) { building.CreateBuildingHandler(c, app) })
	v1Group.DELETE("/building/:id", func(c *gin.Context) { building.DeleteBuildingHandler(c, app) })

	v1Group.GET("/categories", func(c *gin.Context) { category.GetCategoriesHandler(c, app) })
	v1Group.GET("/category/:id", func(c *gin.Context) { category.GetCategoryHandler(c, app) })
	v1Group.PUT("/category", func(c *gin.Context) { category.UpdateCategoryHandler(c, app) })
	v1Group.POST("/category", func(c *gin.Context) { category.CreateCategoryHandler(c, app) })
	v1Group.DELETE("/category/:id", func(c *gin.Context) { category.DeleteCategoryHandler(c, app) })

	v1Group.GET("/organizations", func(c *gin.Context) { organization.GetOrganizationsHandler(c, app) })
	v1Group.GET("/organization/:id", func(c *gin.Context) { organization.GetOrganizationHandler(c, app) })
	v1Group.PUT("/organization", func(c *gin.Context) { organization.UpdateOrganizationHandler(c, app) })
	v1Group.POST("/organization", func(c *gin.Context) { organization.CreateOrganizationHandler(c, app) })
	v1Group.DELETE("/organization/:id", func(c *gin.Context) { organization.DeleteOrganizationHandler(c, app) })

	v1Group.PUT("/building/:id/organizations", func(c *gin.Context) { building.UpdateBuilding2OrganizationsHandler(c, app) })
	v1Group.PUT("/category/:id/organizations", func(c *gin.Context) { category.UpdateCategory2OrganizationsHandler(c, app) })

	return router
}
