package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return os.Getenv("BEARER_TOKEN")
}

func main() {
	token := loadEnv()
	id := "1243050549265612803"

	fmt.Println(token)

	url := "https://api.twitter.com/2/users/" + id + "/following"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
