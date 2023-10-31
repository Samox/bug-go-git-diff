package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	repo, _ := git.PlainOpen("/Users/sammyteillet/.datadrift/default")
	hash := plumbing.NewHash("dbe8af0a6c67ee0e5e903b601cdb8ae8dd11bae0")
	parentHash := plumbing.NewHash("0057446b4b461cdcbf27dfb020395098190a229a")
	commit, err := repo.CommitObject(hash)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(commit.Author.Name)
	}

	parentCommit, err := repo.CommitObject(parentHash)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(parentCommit.Author.Name)
	}

	// Retrieve the trees from the commits
	tree, err := commit.Tree()
	if err != nil {
		log.Fatalf("Could not fetch tree: %s", err)
	}
	parentTree, err := parentCommit.Tree()
	if err != nil {
		log.Fatalf("Could not fetch parent tree: %s", err)
	}

	// Calculate the diff between the trees
	changes, err := object.DiffTree(parentTree, tree)
	if err != nil {
		log.Fatalf("Could not calculate diff: %s", err)
	} else {
		log.Println("Changes:", changes)
	}

	// Iterate over the changes and find the diff for the specific file
	for _, change := range changes {
		log.Println(change.From.Name, change.To.Name)
		if change.From.Name == "mrr_7.csv" || change.To.Name == "mrr_7.csv" {
			fmt.Printf("Change detected in file 'mrr_7':\n")

			// Get the content diff
			patch, err := change.Patch()
			if err != nil {
				log.Fatalf("Could not generate patch: %s", err)
			}
			fmt.Println("Diff:\n", patch.String())
		}
	}
}
