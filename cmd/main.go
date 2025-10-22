// @title Mela API
// @version 1.0
// @description This is the Swapper system backend API built with Go and Gin.
// @termsOfService http://swagger.io/terms/

// @contact.name Mikiyas Daniel
// @contact.url http://github.com/yourusername
// @contact.email mikiyas@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
package main

import (
	"context"
	"fmt"
	"log"

	serverUser "auth-sso/cmd/server/user"
	"auth-sso/internals/infrastructure/config"
	userRouteHandler "auth-sso/internals/infrastructure/handler/user"

	_ "auth-sso/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



func main() {
    // Load config
    cfg, err := config.LoadConfig("../config/config.yaml")
    if err != nil {
        log.Fatal("failed to load config:", err)
    }

    // Build DSN
    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.DBName,
    )
 

    // Init pgx pool
    ctx := context.Background()
    dbPool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        log.Fatalf("unable to connect to database: %v", err)
    }
    defer dbPool.Close()

    // Init user service and handler
    userHandler := serverUser.SetupUser(dbPool);

    // Set up Gin router
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Group all API under /api
    api := r.Group("/api")
    


    // Register domain-specific routes
    userRouteHandler.RegisterRoutes(api, userHandler)

    // Start server
    fmt.Println("🚀 Server is running at http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal("failed to run server:", err)
    }
}
