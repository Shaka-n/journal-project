## Journal Project
TODO: Summary

## How to deploy
Currently, since we're just starting our deployments are incredibly manual.  Please dont make fun of us!
Steps to deploy:
1.  `cd journal-project-backend`
1.  `GOOS=linux GOARCH=amd64 go build .`
1.  Upload the cross compiled binary as a github release
1.  Copy the URL of the release
1.  ssh onto the EC2 server running the code
1.  Kill the existing server
1.  Use wget to download the release
1.  `./journal-project-backend &`
1.  `exit`

