package examples

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func HttpClientExample() {
	// create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	// start server
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// wait server
	time.Sleep(100 * time.Millisecond)

	// client access
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))

	// quit server
	ctx, _ := context.WithTimeout(context.Background(), 0)
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
