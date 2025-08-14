package main

import (
	"context"
	"fmt"
	"log"

	serverUser "auth-sso/cmd/server/user"
	"auth-sso/internals/infrastructure/config"
	userRouteHandler "auth-sso/internals/infrastructure/handler/user"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
