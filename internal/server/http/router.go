package internalhttp

import (
	"github.com/gin-gonic/gin"

	"catalog/internal/app"
)

func NewGinRouter(app *app.App) *gin.Engine {
	router := gin.Default()

	v1Group := router.Group("/api/v1")

	v1Group.GET("/buildings", func(c *gin.Context) { GetBuildingsHandler(c, app) })
	v1Group.GET("/building/:id", func(c *gin.Context) { GetBuildingHandler(c, app) })
	v1Group.PUT("/building", func(c *gin.Context) { UpdateBuildingHandler(c, app) })
	v1Group.POST("/building", func(c *gin.Context) { CreateBuildingHandler(c, app) })
	v1Group.DELETE("/building/:id", func(c *gin.Context) { DeleteBuildingHandler(c, app) })

	v1Group.GET("/categories", func(c *gin.Context) { GetCategoriesHandler(c, app) })
	v1Group.GET("/category/:id", func(c *gin.Context) { GetCategoryHandler(c, app) })
	v1Group.PUT("/category", func(c *gin.Context) { UpdateCategoryHandler(c, app) })
	v1Group.POST("/category", func(c *gin.Context) { CreateCategoryHandler(c, app) })
	v1Group.DELETE("/category/:id", func(c *gin.Context) { DeleteCategoryHandler(c, app) })

	v1Group.GET("/organizations", func(c *gin.Context) { GetOrganizationsHandler(c, app) })
	v1Group.GET("/organization/:id", func(c *gin.Context) { GetOrganizationHandler(c, app) })
	v1Group.PUT("/organization", func(c *gin.Context) { UpdateOrganizationHandler(c, app) })
	v1Group.POST("/organization", func(c *gin.Context) { CreateOrganizationHandler(c, app) })
	v1Group.DELETE("/organization/:id", func(c *gin.Context) { DeleteOrganizationHandler(c, app) })

	v1Group.PUT("/building/:id/organizations", func(c *gin.Context) { UpdateBuilding2OrganizationsHandler(c, app) })
	v1Group.PUT("/category/:id/organizations", func(c *gin.Context) { UpdateCategory2OrganizationsHandler(c, app) })

	return router
}
