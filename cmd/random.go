package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command will give you a random dad joke, fetched from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"

	responseBytes := getJokeData(url)

	randJoke := Joke{}

	if err := json.Unmarshal(responseBytes, &randJoke); err != nil {
		log.Printf("Could not unmarshal response - %v\n", err)
	}

	fmt.Println(randJoke.Joke)
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a dad joke - %v\n", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/learn-go/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not request a dad joke - %v\n", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not request a dad joke - %v\n", err)
	}

	return responseBytes
}
