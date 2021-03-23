package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Label struct for GitHub labels
type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

// GetLabels Get GitHub labels of the repository
func GetLabels(repositoryURL string) ([]Label, error) {
	URL := fmt.Sprintf("%v/labels", repositoryURL)

	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, fmt.Errorf("Couldn't make a new request in GetLabel: %v", err)
	}

	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return nil, fmt.Errorf("Could get the environment variable GITHUB_TOKEN: %v", err)
	}

	token = fmt.Sprintf("bearer %v", token)

	request.Header.Add("Authorization", token)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("Response error in GetLabel: %v", err)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("couldn't convert response body to []byte: %v", err)
	}

	var labels []Label

	err = json.Unmarshal(body, &labels)

	if err != nil {
		return nil, fmt.Errorf("problem unmarshalling the reponse body: %v", err)
	}

	return labels, nil
}