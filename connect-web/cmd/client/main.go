package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	greetv1 "github.com/drakejin/examples/connect-web/gen/greet/v1"
	"github.com/drakejin/examples/connect-web/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		//connect.WithGRPC(),
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	anyMsg := &greetv1.Foo{}
	if res.Msg.HelloAny.MessageIs(anyMsg) {
		fmt.Println("yes It's Foo")
	}
	if err := res.Msg.HelloAny.UnmarshalTo(anyMsg); err != nil {
		fmt.Println(err)
	}
	fmt.Println(anyMsg)

	jsonBytes, _ := json.Marshal(res)
	log.Println(string(jsonBytes))
}
