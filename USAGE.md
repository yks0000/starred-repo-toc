# Generate README.md with your starred repository Table of Content (TOC) for your GitHub Account?

1. Get a Personal Access Token (pat) for your [GitHub Account](https://github.com/settings/tokens).
2. Fork this [repo](https://github.com/yks0000/starred-repo-toc) into your GitHub account.
3. Save PAT as GitHub Secret in forked repository. Name it `GH_TOKEN`. You can set this up under `Settings > Security > Secrets > Actions`
4. Default Schedule is to run at every 30 minutes. You can change it to run as early as every 5 minutes by updating `cron` key in [GitHub action yaml](.github/workflows/generate-md.yml). Adjust accordingly to avoid throttling/rate limiting. A 30 minutes corn schedule is safe.
5. Verify that GitHub action has been executed and your README.md has been updated with Starred Repositories list.

# Testing this locally

1. Get a Personal Access Token (pat) for your GitHub Account.
2. Fork and then `git clone` forked repo.
3. Run this CLI
   1. You can build this locally using `go build -o github-stars main.go` and then run as `./github-stars generate -t <<pat>> -f README.md` OR
   2. You can run this without building using `go run main.go generate -t <<pat>> -f README.md`

Parameters:

`-t/--token` : Required  
`-f/--file`: Optional, Default: `README.md`
