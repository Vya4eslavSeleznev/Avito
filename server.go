package avito

import (
	"context"
	"net/http"
	"time"
)

type Server struct { //структура сервера для запуска http
	httpServer *http.Server
}

// Run у сервера два метода запуск (Run) и остановка работы
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe() //возвращаем метод стандартного http сервера
	// он запускает бесконечный цикл for и слушает все запросы для обработки
}

// Shutdown необходим для выхода из приложения
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
