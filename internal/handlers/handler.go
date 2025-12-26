package handlers

import (
	"github.com/MerBerd/blog-app/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("api/", h.userIdentity)
	{
		articles := api.Group("/articles")
		{
			articles.POST("/", h.createArticle)
			articles.GET("/", h.getAllArticles)
			articles.GET("/:id", h.getArticleById)
			articles.PATCH("/:id", h.updateArticle)
			articles.DELETE("/:id", h.deleteArticle)

			comments := articles.Group("/:id/comments")
			{
				comments.GET("", h.getAllComment)
				comments.POST("", h.createComment)
			}

		}

		comments := api.Group("/comments")
		{
			comments.GET("/:id", h.getCommentById)
			comments.PATCH("/:id", h.changeComment)
			comments.DELETE("/:id", h.deleteComment)
		}
	}

	return router
}
