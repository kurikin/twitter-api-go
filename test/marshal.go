package main

import (
	"encoding/json"
	"fmt"
)

type TweetStruct struct {
	TweetId  string
	UserName string
	UserId   string
}

func main() {
	tweetData := TweetStruct{TweetId: "test", UserName: "kurikin", UserId: "0309"}

	jsonData, err := json.Marshal(tweetData)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(string(jsonData))
}
