package routers

import (
	"server/internal/controller"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	apiV1 := r.Group("/api/v1")

	apiV1Auth := apiV1.Group("/auth")
	{
		apiV1Auth.POST("/login", controller.HandlerLogin)
		apiV1Auth.POST("/register", controller.HandlerRegister)
		apiV1Auth.POST("/logout")
	}
	// apiV1Submission := apiV1.Group("/submission")
	// {
	// 	apiV1Submission.POST("/create",controller.HandlerCreateSubmission)
	// }
	apiV1.POST("/submission", controller.HandlerCreateSubmission)
	apiV1.GET("/submission", controller.HandlerGetSubmission)
	apiV1.DELETE("/submission", controller.HandlerDeleteSubmission)
	apiV1.PUT("/submission", controller.HandlerUpdateSubmission)
	apiV1.GET("/searchsubmission", controller.HandlerSearchSubmission)

	apiV1.POST("/reviewer", controller.HandlerCreateReviewer)
	apiV1.GET("/reviewer", controller.HandlerGetReviewer)
	apiV1.DELETE("/reviewer", controller.HandlerDeleteReviewer)

	apiV1.POST("/reviewersbm", controller.HandlerCreateReviewerSbm)
	apiV1.GET("/reviewersbm", controller.HandlerGetReviewerSbms)
	apiV1.DELETE("/reviewersbm", controller.HandlerDeleteReviewerSbm)

	apiV1.POST("/reviewcmt", controller.HandlerCreateReviewComment)
	// apiV1.GET("/reviewcmt", controller.HandlerGetReviewerSbms)
	// apiV1.DELETE("/reviewcmt", controller.HandlerDeleteReviewerSbm)

	return r
}
