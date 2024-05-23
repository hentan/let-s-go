package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	filePtr, err := os.OpenFile("mylogs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
	l := log.New(filePtr, "", log.LstdFlags)
	mux.HandleFunc("/hello", hello(l))
	authHandlerWithLogging := authHandler(l)
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: authHandlerWithLogging(
			logMiddleware(l)(mux),
		),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(l *log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		msg := "Hello, Go!"
		l.Println("resp:", msg)
		fmt.Fprint(res, msg)
	}
}

func authHandler(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			xId := r.Header.Get("x-my-app-id")
			if xId != "my_secret" {
				l.Println("Неавторизованный доступ: url:", r.URL, "header x-my-app-id:", xId)
				http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)

				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
