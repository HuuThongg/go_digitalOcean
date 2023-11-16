package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err:= ctx.Err(); err != nil {
				fmt.Printf("Doanother err: %s\n", err)
			}
			fmt.Printf("DoAnother finished")
			return	
		case num := <-printch:
			fmt.Printf("DoAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey","myValue")
	doSomething(ctx)
}
