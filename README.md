# cli-chat

Terminal based chat with multiple functionality which has its own
data transfer protocol, written in [go](https://go.dev/).

## Install

You can build it from source with the [go](https://go.dev/) compiler,
or you can check the binaries in the [last release](https://github.com/Guillex387/cli-chat/releases/latest).

> Note: In linux, if do you have any error log with "libc" word
>
> ```bash
> # Note: In linux, if do you have any
> # error log with libc word, check the command below
>
> sudo apt install glibc-source
> # If you are in a distro not based on debian search
> # for installing the libc in your distro
> ```

## Usage

To open a connection view the help message.

```bash
$ cli-chat -help

Terminal based chat with multiple functionality which has its own data transfer protocol.

Options:
  -ip string
    	IP direction to connect to a host (client only) (default "127.0.0.1")
  -mode string
    	Choose the execution mode of the program, "client", "server" (default "client")
  -name string
    	The display name in the chat (client only) (default "unknow")
  -port string
    	The port to listen or connect to a server (default "8000")
```

### Commands

Inside the chat you have different commands:

- `/end` for end the chat connection
- `/kill username` for delete a user (host only)

## License

cli-chat Copyright (c) 2023 Guillex387. All rights reserved.

Licensed under the [GNU AGPLv3](https://github.com/Guillex387/cli-chat/blob/master/LICENSE) license.
