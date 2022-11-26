# usd-101

This is an exploratory project for learning how to handle Unix Domain Sockets (UDS) in Go

## How to run

### Get binaries
- Download the latest release
- From Source Code
  1. Clone the repo
  2. Install [GoReleaser](https://goreleaser.com) and [Task](https://taskfile.dev)
  3. Run `task snap` to build a local snapshot
  
### Run the program

- Launch `uds-101 server`
- Launch `uds-101 client`
- Type any text in the client console and hit Enter to send the message
- `Ctrl+C` for terminating either of the two
  * Be aware that terminating the server will terminate all connected clients, but not viceversa
  
Fill free to run `uds-101 help` for reading about each command and available paramters.

Enjoy!
