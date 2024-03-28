# Snaparser Server

Server that handles parsing of snapchat history file. 

# Installation

```
$ go install github.com/vanillaiice/snaparser_server/cmd/snaparser_server@latest
```

# Usage

To download your chat history data, follow the guide available on snapchat's 
[website](https://help.snapchat.com/hc/en-us/articles/7012305371156-How-do-I-download-my-data-from-Snapchat-). 
You can then do the following:

```sh
# Generic Usage
$ snaparser_server [global options] command [command options]

# Run the server
$ snaparser_server --http --endpoint "/parse"
# or pass a config file in TOML format
$ snaparser_server --load config.toml

# You can use curl to communicate with the server
# and parse your snapchat history file.
curl -F 'file=@chat_history.json;type=application/json' http://localhost:8888/parse -o chats.zip
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
