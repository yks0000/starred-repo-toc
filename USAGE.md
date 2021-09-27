# How to run this ?

1. Get a Personal Access Token (pat) for your GitHub Account
2. Clone this [repo](https://github.com/yks0000/starred-repo-toc):
   1. You can build this locally using `go build -o github-stars main.go ` and then run as `./github-stars generate -t <<pat>> -f README.md` OR
   2. You can run this without building using `go run main.go generate -t <<pat>> -f README.md`

`-t/--token` : Required  
`-f/--file`: Optional, Default=README.md

ProTip: Schedule this as GitHub Actions for updating list automatically. Check here for [github-action config](.github/workflows/generate-md.yml)