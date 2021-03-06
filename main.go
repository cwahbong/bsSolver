package main

import (
	"flag"
	"fmt"
	"github.com/cwahbong/bsSolver/bs"
	"log"
	"net/http"
)

type serverArgs struct {
	Port           uint
	StaticFilePath string
}

func main() {
	var args serverArgs
	flag.UintVar(&args.Port, "port", 80, "Port.")
	flag.StringVar(&args.StaticFilePath, "static-path", "./static/app/", "Static file path.")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(args.StaticFilePath)))
	http.Handle("/j", bs.RpcServer())
	server := http.Server{
		Addr: fmt.Sprintf(":%d", args.Port),
	}
	log.Fatal(server.ListenAndServe())
}
