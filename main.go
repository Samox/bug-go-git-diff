package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	repo, _ := git.PlainOpen(".")
	hash := plumbing.NewHash("a34586878c3410b6cd5cedf9ae604e366502d29f")
	parentHash := plumbing.NewHash("cfc333497085ec3bd924ab0d5a71e525d5ed9910")
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
		if change.From.Name == "mrr.csv" || change.To.Name == "mrr.csv" {
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
