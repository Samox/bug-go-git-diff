# Bug report go-git

The result of patch.String() is not coherent with the result of "git diff"

Run `go run main.go`
Then run `git diff a34586878c3410b6cd5cedf9ae604e366502d29f cfc333497085ec3bd924ab0d5a71e525d5ed9910`
