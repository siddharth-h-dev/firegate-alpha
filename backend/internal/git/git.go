package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
	"fmt"
)

const BasePath = "/etc/firegate/repos"

func GetRepoPath(repoName string) string {
	if repoName == "motd" {
		repoName = "system"
	}
	return fmt.Sprintf("%s/%s", BasePath, repoName)
}

func CommitChanges(repoName, message string) error {
	path := GetRepoPath(repoName)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	
	_,err = w.Add(".")
	if err != nil {
		return err
	}
	
	_,err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name: "Firegate Backend",
			Email: "backend@firegate.internal",
			When: time.Now(),
		},
	})
	return err
}

func GetStatus(repoName string) (string, error) {
	path := GetRepoPath(repoName)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	w, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	status, err := w.Status()
	if err != nil {
		return "", err
	}
	
	return status.String(), nil
}
