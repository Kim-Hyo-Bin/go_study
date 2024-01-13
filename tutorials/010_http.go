package main

import (
	"fmt"
	"net/http"
)

func boardView(w http.ResponseWriter, r *http.Request) {
	pageIdx := r.URL.Query().Get("page_id")
	fmt.Fprintf(w, pageIdx)
}

func boards(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	fmt.Fprintf(w, page+" just page.")
}

func main() {
	port := 8080
	fmt.Printf("서버가 %d 포트에서 실행 중...\n", port)

	http.HandleFunc("/board/", boardView)
	http.HandleFunc("/board", boards)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("서버 시작 실패:", err)
	}
}
