# golife

## Introduction

GoLife is a Multiplayer Interactive simulation of Conway's Game of Life!

The backend processing of the board -as well as all data broadcasting and server state management- is written in Go,
using Protobuf over Websockets to communicate with a React JS frontend. The data format is heavily encoded, in addition
to being packaged in binary Protobuf messages (to avoid the overhead of JSON serialization/deserialization), and as such
is able to handle many simultaneous client connections and broadcasting of a complex and large simulation grid. The backend
processing is also heavily parallelized/parallelizable, albeit with a bottleneck for the multiplexed broadcasting of the
board data.

## Installation
GoLife requires either Go 1.12+ (I am developing w/ 1.13 currently) in addition to Node/yarn for the UI, or, I have
included sample Dockerfiles and a `docker-compose.yml` for a Dockerized deployment.

The backend application can be built using `go build server/server.go`, which will download all dependencies necessary to
run the backend server, and then produce a `server.exe` (or simply `server` on UNIX systems) executable. You can then
execute this program normally (the server will run on port 5000 by default).

You can run the frontend UI using:
```
cd ui
yarn    //downloads dependencies (or use npm install; untested)
yarn start  //start the dev webserver on port 3000 (or use npm start; untested)
```

Alternatively, to deploy via docker-compose, use the command `docker-compose up --build -d`.

You'll want to change the `REACT_APP_SERVICE_URL` in `Dockerfile.ui.prod` to reflect your relevant hostname for your deployment;
if accessing the Docker UI container from the same machine as your deployment, `localhost:5000` should suffice.
