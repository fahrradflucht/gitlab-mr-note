package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/xanzy/go-gitlab"
)

type noteData struct {
	Input string
}

func main() {
	noteData := noteData{Input: readStdin()}
	gitlabClient := gitlab.NewClient(nil, os.Getenv("GITLAB_ACCESS_TOKEN"))
	projectID := os.Getenv("CI_PROJECT_ID")
	mrIID := guessMergeRequestIID(gitlabClient, projectID)
	noteTemplate, _ := template.New("default").Parse("```\n{{.Input}}\n```")

	note := new(bytes.Buffer)
	noteTemplate.Execute(note, noteData)

	mergeRequestNoteOptions := &gitlab.CreateMergeRequestNoteOptions{
		Body: gitlab.String(note.String())}

	if _, _, err := gitlabClient.Notes.CreateMergeRequestNote(
		projectID,
		mrIID,
		mergeRequestNoteOptions); err != nil {
		log.Fatal(err)
	}
}

func guessMergeRequestIID(client *gitlab.Client, projectID string) int {
	mrListRequestOptions := &gitlab.ListProjectMergeRequestsOptions{State: gitlab.String("opened")}

	mrs, _, err := client.MergeRequests.ListProjectMergeRequests(projectID, mrListRequestOptions)

	if err != nil {
		log.Fatal(err)
	}

	for _, mr := range mrs {
		if mr.SHA == os.Getenv("CI_COMMIT_SHA") {
			return mr.IID
		}
	}

	err = errors.New("not found: no merge request with current commits sha present")
	log.Fatal(err)
	panic(err)
}

func readStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	result := ""
	for scanner.Scan() {
		result += fmt.Sprintln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
