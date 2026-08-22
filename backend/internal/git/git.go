// Copyright (C) 2026 Siddharth H
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

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
