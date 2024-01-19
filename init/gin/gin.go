package gin

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
)

const (
	addr = "localhost"
	port = 8080
)

func Start(ctx context.Context) error {
	engine := gin.Default()

	// 使用无缓冲的通道来监听系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

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
	// 等待系统信号或者优雅停止信号
	<-quit

	// 创建一个 context，设置超时时间
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	// 关闭 HTTP 服务
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v\n", err)
	} else {
		return fmt.Errorf("server gracefully stopped")
	}
}
