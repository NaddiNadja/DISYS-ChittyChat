# DISYS-ChittyChat

Submission of solution for Chitty-Chat service made by:

- Theis Stensgaard, thhs
- Nadja Brix Koch, nako
- Caspar Marschall, cafm
- Frederik Rothe, frot

## User manual

### Starting the server

1. Open a terminal and `cd` to the project's directory.

2. Run command

    `$ docker build -t chittychat --no-cache .`

3. Run command

    `$ docker run -p 9080:9080 -tid chittychat`

### Connecting as a client

1. To connect to the chitty-chat service, open a terminal and `cd` to the project's directory.

2. Run command

    `$ go run client/client.go -sender {your name here}`

3. Enjoy chatting with other clients (that are possibly also you, but in another terminal window)

#### Multiple chat rooms

It is possible to create your own chat room and only chat with other people in the same room.

1. When joining, add `-room {name of room here}` to the command

    `$ go run client/client.go -sender {your name here}`

2. Wait for others to join the same room and chat with them!.