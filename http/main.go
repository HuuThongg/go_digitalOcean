// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// )

// const keyServerAddr = "serverAddr"

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
// 	hasFirst := r.URL.Query().Has("first")
// 	first := r.URL.Query().Get("first")
// 	hasSecond := r.URL.Query().Has("second")
	
// 	second := r.URL.Query().Get("second")

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Printf("could not read body: %s\n", err)
// 	}

// 	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
// 		ctx.Value(keyServerAddr),
// 		hasFirst, first,
// 		hasSecond, second,
// 		body)
// 	io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
// 	myName := r.PostFormValue("myName")
// 	if myName == "" {
// 		w.Header().Set("x-missing-field", "myName")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
// }
// func main() {
//     ctx, cancelCtx := context.WithCancel(context.Background())

//     muxOne := http.NewServeMux()
//     muxOne.HandleFunc("/", getRoot)
//     muxOne.HandleFunc("/hello", getRoot)


//     serverOne := &http.Server{
//         Addr:    ":3333",
//         Handler: muxOne,
//         BaseContext: func(l net.Listener) context.Context {
//             ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
//             return ctx
//         },
//     }

//     muxTwo := http.NewServeMux()
//     muxTwo.HandleFunc("/hello", getHello)
//     muxTwo.HandleFunc("/", getRoot)

//     serverTwo := &http.Server{
//         Addr:    ":4444",
//         Handler: muxTwo,
//         BaseContext: func(l net.Listener) context.Context {
//             ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
//             return ctx
//         },
//     }

//     go func() {
//         err := serverOne.ListenAndServe()
//         if errors.Is(err, http.ErrServerClosed) {
//             fmt.Printf("server one closed\n")
//         } else if err != nil {
//             fmt.Printf("error listening for server one: %s\n", err)
//         }
//         cancelCtx()
//     }()
//     go func() {
//         err := serverTwo.ListenAndServe()
//         if errors.Is(err, http.ErrServerClosed) {
//             fmt.Printf("server two closed\n")
//         } else if err != nil {
//             fmt.Printf("error listening for server two: %s\n", err)
//         }
//         cancelCtx()
//     }()

//     <-ctx.Done()
// }
