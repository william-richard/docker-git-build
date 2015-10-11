package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/libgit2/git2go"
)

type Repository struct {
	git_repo git.Repository
	//	watcher  Watcher
}

func NewRepository(clone_path string, remote_url string) (Repository, error) {
	var myrepo Repository
	// ensure the parent of the clone_path exists
	clone_path = filepath.Clean(clone_path)
	// MkdirAll will do nothing if the dir already exists
	os.MkdirAll(clone_path, 0755)
	// Ensure the repo has been initialized
	git_repo, err := git.InitRepository(clone_path, false)
	if err != nil {
		return myrepo, fmt.Errorf("Error initializing repo:", err.Error())
	}
	// ensure the remote on the repo matches the requested remote
	// TODO add support for multiple remotes. Does that mean multiple watchers?
	// remove all the existing remotes
	remotes, err := git_repo.Remotes.List()
	if err != nil {
		return myrepo, fmt.Errorf("Error listing remotes:", err.Error())
	}
	for _, remote_name := range remotes {
		err = git_repo.Remotes.Delete(remote_name)
		if err != nil {
			return myrepo, fmt.Errorf("Error deleting remote", remote_name, ":", err.Error())
		}
	}
	// add the one we want
	git_repo.Remotes.Create("origin", remote_url)

	myrepo = Repository{*git_repo}
	return myrepo, nil
}
