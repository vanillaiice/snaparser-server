# Snaparser Server

Snaparser server is a server that handles parsing of snapchat history files. 

# Installation

## Go install

```
$ go install github.com/vanillaiice/snaparser_server/cmd/snaparser_server@latest
```

## Docker

### Pull image

```sh
$ docker pull vanillaiice/snaparser_server:latest
```

### Build image

```
$ git clone https://github.com/snaparser_server
$ cd snaparser_server
$ docker build -t snaparser_server .
```

# Usage

To download your chat history data, follow the guide available on snapchat's 
[website](https://help.snapchat.com/hc/en-us/articles/7012305371156-How-do-I-download-my-data-from-Snapchat-). 
You can then do the following:

## Using go

```sh
# Generic Usage
$ snaparser_server [global options] command [command options]

# Run the server
$ snaparser_server --http --endpoint "/parse"
# or pass a config file in TOML format
$ snaparser_server --load config.toml

# You can use curl to communicate with the server
# and parse your snapchat history file.
$ curl -F 'file=@chat_history.json;type=application/json' http://localhost:8888/parse -o chats.zip
```

## Using Docker

```sh
# run server with http and enable logging
$ docker run --rm -p 8888:8888 vanillaiice/snaparser_server -t -g

# run server with https on custom port
$ docker run --rm -p 1234:1234 -v $PWD/server.crt:/server.crt -v $PWD/server.key:/server.key vanillaiice/snaparser_server -c server.crt -k server.key -g -p 1234

# run server and pass a toml config file
$ docker run --rm -p 8888:8888 -v $PWD/config.toml:/config.toml vanillaiice/snaparser_server --load config.toml

# cleaner way to pass files
$ mkdir data && cp config.toml server.key server.crt data
$ docker run --rm -p 8888:8888 -v $PWD/data:/data vanillaiice/snaparser_server --load data/config.toml
```

# Flags

```sh
NAME:
   snaparser_server - backend server for snaparser

USAGE:
   snaparser_server [global options] command [command options]

VERSION:
   0.0.1

AUTHOR:
   vanillaiice <vanillaiice1@proton.me>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --allowed-origins ORIGIN(S), -o ORIGIN(S) [ --allowed-origins ORIGIN(S), -o ORIGIN(S) ]  allow ORIGIN(S)
   --allowed-methods METHOD(S), -m METHOD(S) [ --allowed-methods METHOD(S), -m METHOD(S) ]  allow METHOD(S) (default: "POST")
   --port PORT, -p PORT                                                                     listen on PORT (default: 8888)
   --endpoint PATH, -e PATH                                                                 upload endpoint PATH (default: "/upload")
   --limiter value, -i value                                                                HTTP rate limiter type (none, lenient, normal, strict) (default:
 "none")
   --log, -g                                                                                enable logging (default: false)
   --http, -t                                                                               use HTTP instead of HTTPS (default: false)
   --key-file PATH, -k PATH                                                                 SSL secret key file PATH
   --cert-file PATH, -c PATH                                                                SSL certificate file PATH
   --load FILE, -l FILE                                                                     load TOML configuration from FILE
   --help, -h                                                                               show help
   --version, -v                                                                            print the version
```

# Author

vanillaiice

# License

GPLv3
