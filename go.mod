module github.com/farhansabbir/telnet

go 1.23.5

// require (
// 	github.com/farhansabbir/go-ping v1.0.6
// 	golang.org/x/net v0.34.0 // indirect
// )

require github.com/farhansabbir/go-ping v1.1.0

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.9.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
)

replace github.com/farhansabbir/go-ping => ./go-ping/
