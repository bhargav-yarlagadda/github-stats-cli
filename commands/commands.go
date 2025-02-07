package commands

import (
	"encoding/json"
	"fmt"
	"github-stats-cli/utils"
	"os"
)

// Define a type for command functions
type cmdFnc func([]string)

// Store registered commands
var Commands = make(map[string]cmdFnc)

// Register available commands
func InitCommands() {
	// Register the "user", "exit", and "activity" commands
	registerCommand("user", GitHubUserCommand)
	registerCommand("exit", terminateCommand)
	registerCommand("activity", FetchUserActivityCommand)
	registerCommand("help",HelpCommand)
	registerCommand("read",ReadmeCommand)
	registerCommand("list",ListCommand)
	registerCommand("set",SetTokenCommand)
	registerCommand("fork",ForkRepositoryCommand)
	registerCommand("unset",UnsetTokenCommand)
	registerCommand("clear",ClearCommand)



}

// Register a command to the map
func registerCommand(cmd string, fn cmdFnc) {
	Commands[cmd] = fn
}

func ClearCommand(args []string){
	utils.ClearScreen()
}

func UnsetTokenCommand(args[] string){
	utils.UnsetToken()
}
func ReadmeCommand(args []string) {
	// Check if the username and repository name arguments are passed
	if len(args) < 2 {
		fmt.Println("Please provide both GitHub username and repository name as arguments.")
		return
	}

	username := args[0]
	repository := args[1]

	// Fetch the README file content
	readmeContent, err := utils.FetchReadme(username, repository)
	if err != nil {
		fmt.Println("Error fetching README:", err)
		return
	}

	// Print the README content
	if readmeContent != "" {
		fmt.Printf("\nREADME for %s/%s:\n\n", username, repository)
		fmt.Println(readmeContent)
	} else {
		fmt.Println("No README found for the repository:", repository)
	}
}
func HelpCommand(args []string) {
	fmt.Println("Welcome to the GitHub Stats CLI!")
	fmt.Print("\nAvailable Commands:\n")

	// List all commands and their shor t descriptions
	fmt.Println("1. user <username>        - Fetch and display GitHub user stats")
	fmt.Println("2. activity <username>    - Fetch and display recent activity for the GitHub user")
	fmt.Println("3. read username reponame - Fetch and display readme of the repository")
	fmt.Println("4. exit                   - Exit the program")
	fmt.Println("5. list username          - Fetch and display list of all public repository names of a user.")
	fmt.Println("6. set githubToken        - sets token to authorize user to perform fork and start.")	
	fmt.Println("7. fork username reponame - sets token to authorize user to perform fork and start.")	
	fmt.Println("8. unset                  - clears token data.")	
	fmt.Println("9. clear                  - clears screen (std out)")	
}


func SetTokenCommand(args []string){
	token := args[0]
	if token == ""{
		fmt.Fprint(os.Stdout,"please enter token")
		
	}else{
		utils.SetToken(token)
	}
	 
}

func ForkRepositoryCommand(args []string){
	fmt.Fprintf(os.Stdout,"Make Sure You have set the classic PAT gihub-api token instead of fine-tuned tokens that has access to public repos and repos.\n")
	if len(args) < 2 {
		fmt.Println("Please provide a GitHub username and repository name as argument.")
		return
	}else{
		username:=args[0]
		repoName :=args[1]
		err :=utils.ForkRepository(username,repoName)
		fmt.Fprint(os.Stdout,err)
	}
}


func FetchUserActivityCommand(args []string) {
	// Check if the username argument is passed
	if len(args) < 1 {
		fmt.Println("Please provide a GitHub username as an argument.")
		return
	}
	username := args[0]

	// Fetch the activity data
	activities, err := utils.FetchUserActivity(username)
	if err != nil {
		fmt.Println("Error fetching activity:", err)
		return
	}

	// Print the activity details
	if len(activities) == 0 {
		fmt.Println("No activity found for user:", username)
		return
	}

	fmt.Printf("\n--- Recent Activity for %s ---\n", username)
	for _, activity := range activities {
		fmt.Printf("\nActivity Type: %s\n", activity.Type)
		fmt.Printf("Repository: %s\n", activity.Repo.Name)
		fmt.Printf("Created At: %s\n", activity.CreatedAt)

		// Print the Payload based on activity type
		switch activity.Type {
		case "PushEvent":
			var pushPayload struct {
				Ref string `json:"ref"`
			}
			if err := mapToStruct(activity.Payload, &pushPayload); err != nil {
				fmt.Println("Error parsing push event payload:", err)
			} else {
				fmt.Printf("Push to branch: %s\n", pushPayload.Ref)
			}

		case "PullRequestEvent":
			var prPayload struct {
				Action string `json:"action"`
			}
			if err := mapToStruct(activity.Payload, &prPayload); err != nil {
				fmt.Println("Error parsing pull request event payload:", err)
			} else {
				fmt.Printf("PR action: %s\n", prPayload.Action)
			}

		case "IssueCommentEvent":
			var issueCommentPayload struct {
				Body string `json:"body"`
			}
			if err := mapToStruct(activity.Payload, &issueCommentPayload); err != nil {
				fmt.Println("Error parsing issue comment payload:", err)
			} else {
				fmt.Printf("Comment: %s\n", issueCommentPayload.Body)
			}

		default:
			// For other event types, just print the payload as-is
			fmt.Printf("Payload: %v\n", activity.Payload)
		}
	}
	fmt.Print("\n--- End of Activity ---\n")
}

// Helper function to map a generic map to a struct
func mapToStruct(payload map[string]interface{}, out interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

// Terminate the program
func terminateCommand(args []string) {
	utils.ExitProgram()
}

// GitHubUserCommand fetches and displays GitHub user data
func GitHubUserCommand(args []string) {
	// Check if the username argument is passed
	if len(args) < 1 {
		fmt.Println("Please provide a GitHub username as an argument.")
		return
	}
	username := args[0]

	// Fetch GitHub user data
	user, err := utils.FetchGitHubUser(username)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// Display GitHub user data
	fmt.Printf("\n--- GitHub Stats for %s ---\n", user.Login)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Public Repositories: %d\n", user.PublicRepos)
	fmt.Printf("Followers: %d\n", user.Followers)
	fmt.Printf("Following: %d\n", user.Following)
	fmt.Printf("Public Gists: %d\n", user.PublicGists)
	fmt.Printf("Avatar URL: %s\n", user.AvatarURL)
	fmt.Print("\n--- End of User Stats ---\n")
}


func ListCommand(args []string) {
    if len(args) != 1 {
        fmt.Println("Usage: list <username>")
        return
    }
    username := args[0]
    repoList, err := utils.ListRepositories(username)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Returing First ",len(repoList)," for ", username)
    for _, repo := range repoList {
        fmt.Println(repo)
    }
}
