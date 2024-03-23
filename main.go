package main

import (
	"bufio"
	"educationlsp/rpc"
	"log"
	"os"
)

func main() {
    loger := getLoger("/home/alexjamison/work/lsp/log.txt")
    loger.Println("Hey, I started!")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)
    for scanner.Scan() {
        msg := scanner.Text()
        handlMessage(loger, msg)
    }
}
func handlMessage(logger *log.Logger,msg any) {
    logger.Println(msg)
}

func getLoger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("hey, you didnt give me a good file")
    }
    return log.New(logfile, "[educationlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
