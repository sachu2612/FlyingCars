package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

type User struct {
    gorm.Model
    FirstName    string
    LastName     string
    MobileNumber string
    PassKey      string
}

type Car struct {
    Model             string
    DateOfManufacture string
    LastServicedDate  string
    UniqueID          string
    LastUsedDate      string
}
type dbConfig struct {
    driverName     string
    dataSourceName string
}

type AuthService struct {
    db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{db: db}
}
type authHandler struct {
	db     *sql.DB
	config *config
}
func getConfig() *dbConfig {
    // Read database configuration from a file or environment variable
    // and return it as a struct
    return &dbConfig{
        driverName:     "postgres",
        dataSourceName: "user=postgres password=postgres dbname=car_db host=localhost port=5432 sslmode=disable",
    }
}

func InitDB() (*gorm.DB, error) {
    // Load database configuration from a file or environment variable
    dbConfig := getConfig()

    // Create a new database connection
    db, err := gorm.Open(dbConfig.driverName, dbConfig.dataSourceName)
    if err != nil {
        return nil, err
    }

    // Ping the database to check if the connection is successful
    err = db.DB().Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
authHandler := NewAuthHandler(authService)
func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err.Error())
    }

    // Initialize database connection
    db, err := InitDB()
    if err != nil {
        log.Fatalf("Error initializing database connection: %s", err.Error())
    }
    defer db.Close()

    // Initialize auth service
    authService := NewAuthService(db)

    // Initialize gin router
    router := gin.Default()

    // Set up GraphQL handler for auth service
    authHandler := NewAuthHandler(authService)
    router.POST("/auth", authHandler.GraphQLHandler())

    // Start server
    srv := &http.Server{
        Addr:    ":3000",
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error starting server: %s", err.Error())
        }
    }()

    // Wait for interrupt signal to gracefully shut down the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    // Set up context with timeout for shutting down server gracefully
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Error shutting down server: %s", err.Error())
    }

    log.Println("Server shut down successfully.")
}
