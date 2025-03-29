package sdk

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/surya-apty/goro/logger"
	middleware "github.com/surya-apty/goro/middlewares"
	"go.uber.org/zap"
)

type HandlerFunc func(*Context)

type App struct {
	hostname   string
	router     *httprouter.Router
	middleware []HandlerFunc
	pool       sync.Pool
	prefix     string
	log        *logger.Logger
	env        string
	port       string
	startTime  string
}

func New() *App {
	// Initialize logger
	log := logger.InitLogger()
	r := httprouter.New()
	app := &App{
		router: r,
		pool: sync.Pool{
			New: func() any {
				return &Context{}
			},
		},
		log:       log,
		env:       "development",
		startTime: time.Now().Format("2006-01-02 15:04:05"),
		hostname:  "localhost",
		port:      "8080",
	}
	return app
}

func (a *App) Use(mw HandlerFunc) {
	a.middleware = append(a.middleware, mw)
}

func (a *App) Listen(addr string) error {
	// Print beautiful server start message
	fmt.Println()
	fmt.Println(string(bold) + string(cyan) + "🚀 Goro Server Started" + string(reset))
	fmt.Println(string(green) + "---------------------------" + string(reset))
	fmt.Printf("%s%-15s%s %s\n", blue, "Environment:", reset, a.env)
	fmt.Printf("%s%-15s%s http://%s:%s\n", blue, "Listening on:", reset, a.hostname, a.port)
	fmt.Printf("%s%-15s%s %s\n", blue, "Started at:", reset, a.startTime)
	fmt.Println(string(green) + "---------------------------" + string(reset))

	// Create server
	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: middleware.RequestLogger(a.log.Logger)(a.router), // Use the app's router as the handler
	}

	// Start server in a goroutine
	go func() {
		a.log.Info("Starting server...", zap.String("port", a.port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Fatal("server failed to start", zap.Error(err))
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		a.log.Error("server shutdown error", zap.Error(err))
		return err
	}

	a.log.Info("Server exited gracefully")
	return nil
}
