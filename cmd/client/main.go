package main

import (
	"google.golang.org/grpc"
	"myapp/myapp"
	"context"
	"strconv"
	"flag"
	"log"

)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := myapp.NewMyAppClient(conn)
	res, err := c.Sum(context.Background(), &myapp.AddRequest{ X: int32(x), Y: int32(y) })
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}

