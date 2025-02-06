# GitHub CLI Tool

A simple CLI tool that interacts with the GitHub API to perform several actions such as:
- Fetching and displaying user repositories.
- Showing the contents of a specific repository.
- Listing repository information in a clean format.

## Features
- List all repositories for a given GitHub username.
- Display the `README.md` of a specific repository if it exists.
- List repositories in the format: `reponame - repolink`.

## Prerequisites

- Go 1.16 or higher installed on your system.
- An active internet connection to interact with the GitHub API.

## Commands
1. user <username>        - Fetch and display GitHub user stats
2. activity <username>    - Fetch and display recent activity for the GitHub user
3. read username reponame - Fetch and display readme of the repository
4. exit                   - Exit the program
5. list username          - Fetch and display list of all public repository names of a user
6. help                   - Display the list of available commands and their syntax

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/bhargav-yarlagadda/github-stats-cli.git
```
2. Move to the Directory:
  ```bash
cd github-stats-cli  
```
3. Build the project:
  ```bash
go mod tidy
go build -o ghcli
```