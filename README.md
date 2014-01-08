# Denser server

Denser is a super simple client/server pair (this is the server) that keeps
track of the remote IP address of a machine. It is very specific to my
usecase. The idea is that the server sits on a machine I control, and a client
sits on a machine with a frequently changing IP address. Every 5 minutes, the
client pings the server and tells it what the current public IP address is (as
given by [icanhazip](http://icanhazip.com/).

That's pretty much it. Both the client and the server use the same configuration
file, `~/.denser`. It can contain the following:

```toml
endpoint = "localhost:3333/set"
port = 3333
```

`port` is optional and defaults to `3245` if it isn't set in the configuration file.
