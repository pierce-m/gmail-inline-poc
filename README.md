# gmail-inline-poc

## Setup

### Ngrok Tunneling
First, set up an ngrok tunnel. Download ngrok if you don't already have it: ` brew install ngrok/ngrok/ngrok`

Start a TCP tunnel like `ngrok tcp 1025`. `1025` is the local port that will receive traffic for the mini-pov.

Go into Starwood's admin console. Navigate to `Apps > Google Workspace > Settings for Gmail > Hosts` and edit the `test pipe` route. Ensure this route's `Single host` value matches that of the ngrok tunnel you just started.

### Starting the server
Open a new terminal window and navigate to the root of this repo. Do `go run main.go` to start the server. Now outbound emails will be routed to this server on their way to the destination.
