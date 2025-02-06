package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GitHubUser represents the structure of GitHub user data
type GitHubUser struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	PublicRepos int    `json:"public_repos"`
	Followers  int    `json:"followers"`
	Following  int    `json:"following"`
	PublicGists int   `json:"public_gists"`
	AvatarURL  string `json:"avatar_url"`
}
type GitHubRepo struct {
	Name string `json:"name"`
}

// GitHubActivity represents the structure of the user's activity.
type GitHubActivity struct {
	Type      string                 `json:"type"`
	Repo      GitHubRepo             `json:"repo"`
	CreatedAt string                 `json:"created_at"`
	Payload   map[string]interface{} 
}

type Repository struct {
	Name string `json:"name"`
	HTMLURL string `json:"html_url"`
}



func FetchReadme(username, repository string) (string, error) {
	// GitHub API endpoint for fetching README
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/README.md?ref=master", username, repository)

	// Send GET request to GitHub API
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch README: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("repository not found or no README available (status code: %d)", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the content field from the JSON response
	// Assuming the response is a JSON object that includes a "content" field with base64-encoded data
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Base64 decode the content field
	encodedContent := response["content"].(string)
	decodedContent, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return "", fmt.Errorf("failed to decode README content: %v", err)
	}

	// Return the decoded README content as a string
	return string(decodedContent), nil
}

// FetchGitHubUser makes a request to GitHub API and returns user details
func FetchGitHubUser(username string) (*GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GitHub data: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("returned status code: %v", resp.StatusCode)
	}

	var user GitHubUser
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &user, nil
}
func ExitProgram(){
    os.Exit(0)
}

func FetchUserActivity(username string) ([]GitHubActivity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch activity: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("returned status code: %v", resp.StatusCode)
	}

	var activities []GitHubActivity
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return activities, nil
}

func ListRepositories(username string) ([]string, error) {
	// GitHub API endpoint to list repositories for a user
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	// Send GET request to GitHub API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repositories: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve repositories (status code: %d)", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response into a list of repositories
	var repositories []Repository
	err = json.Unmarshal(body, &repositories)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Format the repositories as "reponame - repolink"
	var repoList []string
	for _, repo := range repositories {
		repoList = append(repoList, fmt.Sprintf("%s - %s", repo.Name, repo.HTMLURL))
	}

	return repoList, nil
}