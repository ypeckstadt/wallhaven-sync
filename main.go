package main

import (
	"encoding/json"
	response2 "github.com/ypeckstadt/wallhaven-sync/response"
	"github.com/ypeckstadt/wallhaven-sync/result"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	outputFolder := "/Users/****/Downloads/wallhaven"
	apiKey := "eX****"
	username := "****"

	err := prepareOutputFolder(outputFolder)
	logFatalWhenError(err)

	existingFiles, err := getOutputFolderFiles(outputFolder)
	logFatalWhenError(err)

	collections, err := getCollections(apiKey)
	logFatalWhenError(err)

	// Get "Default" collection ID
	var defaultCollectionID int
	for _, collection := range collections.Collections {
		if collection.Label == "Default" {
			defaultCollectionID = collection.ID
			break
		}
	}

	// check if default collection id has been found
	if defaultCollectionID > 0 {
		log.Print("Default collection ID has been found")
		syncResult, err := syncForCollection(defaultCollectionID, username, apiKey, existingFiles, outputFolder)
		logFatalWhenError(err)

		log.Println("========================================================")
		log.Println("========================================================")
		log.Printf("%d existing wallpapers have been skipped", syncResult.SkippedPicturesCount)
		log.Printf("%d new wallpapers have been added", syncResult.NewPicturesCount)
		log.Printf("%d wallpapers have been deleted", syncResult.DeletedPicturesCount)
	}
}

func getOutputFolderFiles(folder string) (map[string]bool, error) {
	outputFolderFiles := make(map[string]bool) // k: id, v: found or not during sync

	// read all files in the output folder and add to map
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return outputFolderFiles, err
	}

	// add all files to map for lookup, set to 0 so we can later add a
	for _, f := range files {
		outputFolderFiles[f.Name()] = false
	}
	return outputFolderFiles, nil
}

func prepareOutputFolder(folder string) error {
	_, err := os.Stat(folder)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) {
		return os.Mkdir(folder, 0700)
	}
	return nil
}

func logFatalWhenError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func syncForCollection(collectionID int, username string, apiKey string, files map[string]bool, folder string) (result.SyncResult, error) {
	currentPage := 1
	keepSyncing := true
	syncResult := result.SyncResult{}

	// keep syncing until all pages are loaded
	for keepSyncing {
		log.Printf("Syncing page %d ...", currentPage)
		// get the pictures in the default collection
		getPicturesResponse, err := http.Get("https://wallhaven.cc/api/v1/collections/" + username + "/" + strconv.Itoa(collectionID)+ "?apikey=" + apiKey + "&page=" + strconv.Itoa(currentPage))
		if err != nil {
			return syncResult, err
		}
		defer getPicturesResponse.Body.Close()

		// read the payload
		body, err := ioutil.ReadAll(getPicturesResponse.Body)
		if err != nil {
			return syncResult, err
		}

		var wallpapersResponse response2.WallpaperResponse
		err = json.Unmarshal(body, &wallpapersResponse)
		if err != nil {
			return syncResult, err
		}


		for _, wallpaper := range wallpapersResponse.Wallpapers {
			// determine file extension by file type
			var extension string
			if wallpaper.FileType == "image/png" {
				extension = ".png"
			}
			if wallpaper.FileType == "image/jpeg" {
				extension = ".jpg"
			}

			fileName := wallpaper.ID + extension

			// check if the picture is already downloaded and store in the output folder
			if _, ok := files[fileName]; !ok {
				log.Printf("Saving new file %s",fileName)

				// create output for output
				output, err := os.Create(folder + "/" + fileName)
				if err != nil {
					return syncResult, err
				}
				defer output.Close()

				// download output and save to target folder
				resp, err := http.Get(wallpaper.Path)
				if err != nil {
					return syncResult, err
				}
				defer resp.Body.Close()


				// copy output to output
				_, err = io.Copy(output, resp.Body)
				if err != nil {
					return syncResult, err
				}

				syncResult.NewPicturesCount++
			} else { // file is found, mark as being found in the sync
				files[wallpaper.ID + extension] = true
				syncResult.SkippedPicturesCount++
			}
		}

		if currentPage < wallpapersResponse.Meta.LastPage {
			currentPage++
		} else {
			keepSyncing = false
		}
	}

	// remove pictures that are not in the collection anymore
	for fileName, hasBeenFoundInSync := range files {
		if !hasBeenFoundInSync {
			log.Printf("Deleting file %s",fileName)
			err := os.Remove(folder + "/" + fileName)
			if err != nil {
				return syncResult, err
			}
			syncResult.DeletedPicturesCount++
		}
	}
	return syncResult, nil
}
