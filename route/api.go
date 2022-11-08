package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handler "github.com/joshuaetim/quiz/handler"
	infrastructure "github.com/joshuaetim/quiz/infrastructure"
	middleware "github.com/joshuaetim/quiz/middleware"
)

func RunAPI(address string) error {
	db := infrastructure.DB()
	userHandler := handler.NewUserHandler(db)
	staffHandler := handler.NewStaffHandler(db)
	visitorHandler := handler.NewVisitorHandler(db)
	dashboardHandler := handler.NewDashboardHandler(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to CSC Quiz v1")
	})
	apiRoutes := r.Group("/api")

	apiRoutes.GET("/checkauth", middleware.AuthorizeJWT(), handler.CheckAuth)

	userRoutes := apiRoutes.Group("/auth")
	userRoutes.POST("/register", userHandler.CreateUser)
	userRoutes.POST("/login", userHandler.SignInUser)

	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeJWT())
	userProtectedRoutes.GET("/:id", userHandler.GetUser)
	userProtectedRoutes.PUT("/", userHandler.UpdateUser)

	staffRoutes := apiRoutes.Group("/staff", middleware.AuthorizeJWT())
	staffRoutes.GET("/:id", staffHandler.GetUserStaff)
	staffRoutes.GET("/", staffHandler.GetAllStaffByUser)
	staffRoutes.POST("/", staffHandler.CreateUserStaff)
	staffRoutes.PUT("/:id", staffHandler.UpdateUserStaff)
	staffRoutes.DELETE("/:id", staffHandler.DeleteUserStaff)

	visitorRoutes := apiRoutes.Group("/visitors", middleware.AuthorizeJWT())
	visitorRoutes.GET("/", visitorHandler.GetAllUserVisitors)
	visitorRoutes.POST("/", visitorHandler.CreateUserVisitor)
	visitorRoutes.GET("/:id", visitorHandler.GetUserVisitor)
	visitorRoutes.GET("/staff/:staffID", visitorHandler.GetAllStaffVisitors)
	visitorRoutes.PUT("/:id", visitorHandler.UpdateUserVisitor)
	visitorRoutes.DELETE("/:id", visitorHandler.DeleteUserVisitor)

	dashboardRoutes := apiRoutes.Group("/dashboard", middleware.AuthorizeJWT())
	dashboardRoutes.GET("/users/count", dashboardHandler.GetUsersCount)

	return r.Run(address)
}
