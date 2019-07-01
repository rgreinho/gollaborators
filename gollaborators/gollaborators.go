package gollaborators

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/google/go-github/v26/github"
)

// Retrieve retrieves the collaborator information.
func Retrieve(owner, repository string, lineLength int) error {
	client := github.NewClient(nil)
	ctx := context.Background()
	collaborators, _, err := client.Repositories.ListContributors(ctx, owner, repository, &github.ListContributorsOptions{})
	if err != nil {
		return err
	}

	// Prepare the template.
	tmpl, err := template.New("test").Parse("|[<img src=\"{{.AvatarURL}}\" height=\"100px;\" style=\"border-radius: 50%;\"><br /><sub><b> {{.Name}}</b></sub>]({{.HTMLURL}})<br /> 💻🎨📜")
	if err != nil {
		return err
	}

	// Determine the maximum number of items per line.
	itemsPerLine := min(lineLength, len(collaborators))

	// Prepare the header.
	fmt.Printf("%s|\n", strings.Repeat("| ", itemsPerLine))
	fmt.Printf("%s|\n", strings.Repeat("| :---: ", itemsPerLine))

	// Prepare the content.
	count := 1
	for _, collaborator := range collaborators {
		user, _, _ := client.Users.Get(ctx, collaborator.GetLogin())
		if err := tmpl.Execute(os.Stdout, user); err != nil {
			return err
		}
		if count%itemsPerLine == 0 {
			fmt.Println("|")
		}
		count++
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
