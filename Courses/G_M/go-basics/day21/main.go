// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler started")
// 	time.Sleep(5 * time.Second)
// 	fmt.Fprintln(w, "hello")
// 	fmt.Println("handler finished")
// }

// // step 2
// func main() {
// 	// http.HandleFunc("/hello", helloHandler)
// 	// fmt.Println("server started on :8080")
// 	// http.ListenAndServe(":8080", nil)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/hello", helloHandler)

// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}
// 	// запускаем сервер
// 	go func() {
// 		fmt.Println("server started on :8080")
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			fmt.Println("server error:", err)
// 		}
// 	}()

// 	// ловим сигналы ОС
// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

// 	<-stop
// 	fmt.Println("\nshutdown signal received")

// 	// graceful shutdown
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	if err := server.Shutdown(ctx); err != nil {
// 		fmt.Println("shutdown error:", err)
// 	}

// 	fmt.Println("server gracefully stopped")
// }

// step 1

// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("handler started")
// 		w.Write([]byte("hello\n"))
// 		fmt.Println("handler finished")

// 		fmt.Fprintln(w, "hello world")
// 	})
// 	fmt.Println("server started")
// 	http.ListenAndServe(":8080", nil)
// }
//
//
//
//
//

// ##################
// 		Step 1
//   HTTP с нуля
// ##################

// package main

// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello"))
// }

// func main() {
// 	http.HandleFunc("/hello", hello)
// 	http.ListenAndServe(":8080", nil)
// }

// 1️⃣ Кто такой клиент
// это  curl, браузер или postman
// 2️⃣ Кто такой сервер
// это моя программа написанная на go
// 3️⃣ Что делает handler
// это функция хэндлер вызываемая в ответ на http запрос клиента
// 4️⃣ Почему ListenAndServe не завершается
// ждет пока не произойдет неисправимая ошибка или не будет отправлен сигнал завершения
// 5️⃣ Что происходит, если открыть /hello в браузере
// не получится получить доступ к файлу, но если указать http://localhost:8080/hello,
// то выведется надпись Hello

// ##################
// 		Step 2
// ServeMux и HandlerFunc
// ##################

// 1️⃣ Что такое ServeMux
// это встроиный в го иструмент для сопоставления http запроса(путь) и функции для ответа
// 2️⃣ Что делает HandleFunc
// регистрирует маршрут /hello
// связывает его с функцией hello
// кладёт это в DefaultServeMux
// 3️⃣ Почему nil передаётся в ListenAndServe
// для указания что используется DefaultServeMux

// ##################
// 		Step 3
// HTTP graceful shutdown
// ##################

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// запкскаем сервер
	go func() {
		fmt.Println("server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("serever error:", err)
		}
	}()

	// ждем сигнал завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nshutdown started")

	// ждем 5 секунд на завершение запроса

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	fmt.Println("server stoped")
}

// 1️⃣ Зачем http.Server, если есть ListenAndServe
// собственый srv нужен для передачи остановки сигнала из вне
// 2️⃣ Почему запуск сервера в goroutine
// что бы не блокировался main
// 3️⃣ Что делает Shutdown(ctx)
// передает контекст для корректной остановки сервера
