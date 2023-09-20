My printer drivers are... exotic. Got it to work on a home PC. The various mobile phones need the printer too. It also doesn't really like good old postscript. So, this is a go program that serves a http form with file upload as well as printing options, for access for my devices on the local network. If you have a weird printer, you too, can resort to the horror-show of printing over HTTP.

For a dynamically linked binary, `go build -o exe main.go && strip exe`.

For a static binary (on linux), `go build -tags netgo -ldflags '-extldflags "-static"' -o exe main.go && strip exe`
