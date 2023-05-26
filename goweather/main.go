package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	// "io/ioutil"

	// "io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	apiKey := goDotEnvVariable("API")
	reader := bufio.NewReader(os.Stdin)
	userinput, _ := reader.ReadString('\n')

	url := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + userinput

	fmt.Print(url)

	response, err := http.Get(strings.TrimSpace(url))

	if err != nil {
		fmt.Print(err)
		panic("Error happend while fetch")
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err)
		panic("Error happend while reading")
	}

	// fmt.Print(string(responseData))

}
