# webhook2status
webhook2status: an awesome and simple request generator to change Slack status via webhook, IFTTT and curl.

Live version is here: https://webhook2status.ksswre.net/

## Installation
1. Create Slack app from Slack App Directory and issue `CLIENT_ID` and `CLIENT_ID_SECRET`.
1. Setup `.env` file using template `.env.example`.
1. Install dependencies: echo, godotenv, oauth2 (TODO: use dep).
1. `go run main.go` to launch webserver, visit `localhost:8888` from your web browser.

## TODO
- Add request generator to update Slack status ([API Doc](https://api.slack.com/methods/users.profile.set)).

## LICENSE
MIT
