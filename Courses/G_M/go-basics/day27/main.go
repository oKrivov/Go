package main

import (
	"fmt"
	"net/http"
	"time"
)

// ##################
//		semaphore
// ##################

// максимум 2 одновременных запроса
// var sem = make(chan struct{}, 2)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	// пытаемся занять слот
// 	select {
// 	case sem <- struct{}{}:
// 		// слот получен
// 		defer func() { <-sem }() // освободим в конце
// 	default:
// 		// слотов нет
// 		http.Error(w, "server busy", http.StatusTooManyRequests)
// 		return
// 	}

// 	fmt.Println("request started")
// 	time.Sleep(3 * time.Second)
// 	fmt.Fprintln(w, "done")
// 	fmt.Println("request finished")
// }

// func main() {
// 	http.HandleFunc("/hello", hello)

// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

//########################
// отмена запроса через context.Context
//########################

var sem = make(chan struct{}, 2)

func hello(w http.ResponseWriter, r *http.Request) {
	// попытка занять слот
	select {
	case sem <- struct{}{}:
		defer func() {
			<-sem
		}()
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
		return
	}

	fmt.Println("requst starteed")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "Congratulations, the job is done")
		fmt.Println("request finished")
	case <-r.Context().Done():
		fmt.Println("request canceled by client")
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", nil)
}
