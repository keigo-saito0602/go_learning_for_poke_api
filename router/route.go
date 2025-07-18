package router

import (
	"net/http"

	"github.com/keigo-saito0602/go_learning_for_poke_api/auth"
	"github.com/keigo-saito0602/go_learning_for_poke_api/interface/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/keigo-saito0602/go_learning_for_poke_api/docs"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handlers) {
	ApplyCORS(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User
	e.GET("/users", h.User.ListUsers)
	e.GET("/users/:id", h.User.GetUser)
	e.POST("/users", h.User.CreateUser)
	e.PUT("/users/:id", h.User.UpdateUser)
	e.DELETE("/users/:id", h.User.DeleteUser)

	// Login
	e.POST("/login", h.User.Login)

	// Memo
	e.GET("/memos", h.Memo.ListMemos)
	e.GET("/memos/:id", h.Memo.GetMemo)

	RegisterAuthRoutes(e, h)
}

func RegisterAuthRoutes(e *echo.Echo, h *handler.Handlers) {
	authGroup := e.Group("/auth")
	authGroup.Use(auth.JWTMiddleware)
	authGroup.GET("/me", func(c echo.Context) error {
		userID := c.Get("userID").(uint64)
		return c.JSON(http.StatusOK, map[string]any{
			"user_id": userID,
		})
	})

	adminGroup := authGroup.Group("/admin")
	adminGroup.Use(auth.AdminOnlyMiddleware)
	adminGroup.GET("/settings", h.User.AdminOnlySetting)

	// 認証機能が必要なAPIの場合以下に追加
	authGroup.POST("/memos", h.Memo.CreateMemo)
	authGroup.PUT("/memos/:id", h.Memo.UpdateMemo)
	authGroup.DELETE("/memos/:id", h.Memo.DeleteMemo)
}
