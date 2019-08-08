// explain the meaning of options of gcflags
// -l:disable function inlining
// -m:show the escape analysis and the choice of inlining
// -N:disable optimization

//GODEBUG
// Briefly explain the meaning of the following options of GODEBUG
// GODEBUG=gctrace=1,schedtrace=1000
//answer
// Print the debug info of gc and go scheduler's states every 1000ms

// $GOROOT is the root directory for standard library, including executables, source code and docs.
// $GOPATH is the directory for 3rd party packages.

//Briefly describe how to use pprof to monitor the performance of http server.

// Add `import _ "net/http/pprof"` in the main.go
// Run the following commands to get info correspondingly:
// CPU profile:
// go tool pprof http://localhost:6060/debug/pprof/profile --second N
// heap profile:
// go tool pprof http://localhost:6060/debug/pprof/heap
// goroutine blocking profile:
// go tool pprof http://localhost:6060/debug/pprof/block

// ./...
package main
