
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/ypeckstadt/wallhaven-sync/pkg"
	response2 "github.com/ypeckstadt/wallhaven-sync/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Wallhaven.cc collections",
	Long: `list is for retrieving the data for your Wallhaven.cc collections`,
	Run: func(cmd *cobra.Command, args []string) {

		apiKey, err := cmd.Flags().GetString("api-key")
		pkg.LogFatalWhenError(err)

		response, err := getCollections(apiKey)
		pkg.LogFatalWhenError(err)

		log.Println("Found collections (label - id):")
		for _, collection := range response.Collections {
			log.Println(collection.Label + " - " + strconv.Itoa(collection.ID))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getCollections(apiKey string) (response2.CollectionsResponse, error) {
	var collectionsResponse response2.CollectionsResponse

	// get collections
	getCollectionsResponse, err := http.Get("https://wallhaven.cc/api/v1/collections?apikey=" + apiKey)
	if err != nil {
		return collectionsResponse, err
	}

	defer getCollectionsResponse.Body.Close()

	// read the payload
	body, err := ioutil.ReadAll(getCollectionsResponse.Body)
	if err != nil {
		return collectionsResponse, err
	}

	// parse to struct
	err = json.Unmarshal(body, &collectionsResponse)
	if err != nil {
		return collectionsResponse, err
	}

	return collectionsResponse, nil
}