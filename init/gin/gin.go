package gin

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	addr = "localhost"
	port = 8080
)

func Start(ctx context.Context) error {
	engine := gin.Default()

	// 创建一个 HTTP 服务
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, port),
		Handler: engine,
	}

	// 启动 HTTP 服务（在 Goroutine 中）
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("shutting down gracefully, press Ctrl+C again to force")
	}

	// 创建一个 context，设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v\n", err)
	} else {
		return fmt.Errorf("server gracefully stopped")
	}
}
