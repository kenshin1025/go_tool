package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	//Todo: チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			//Todo: チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
		//Todo: チャネルを閉じる
		close(ch)
	}()
	//Todo: チャネルを返す
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
