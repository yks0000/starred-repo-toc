# How to run this for your GitHub Account?

1. Get a Personal Access Token (pat) for your GitHub Account.
2. Save PAT as GitHub Secret. Name it `GH_TOKEN`.
3. Fork this [repo](https://github.com/yks0000/starred-repo-toc) in your GitHub account.
4. Default Schedule is to run at every 15 minutes. You can change it to run as early as every 5 minutes by updating `cron` key in GitHub action yaml.
5. Verify that GitHub action has been executed and your README.md has been updated with Starred Repositories list.

# Testing this locally?

1. Get a Personal Access Token (pat) for your GitHub Account.
2. Clone or fork and then clone this [repo](https://github.com/yks0000/starred-repo-toc)
3. Run this CLI
   1. You can build this locally using `go build -o github-stars main.go` and then run as `./github-stars generate -t <<pat>> -f README.md` OR
   2. You can run this without building using `go run main.go generate -t <<pat>> -f README.md`

Parameters:

`-t/--token` : Required  
`-f/--file`: Optional, Default=README.md