package main

import (
	"fmt"
	"syscall"
)

type Error struct {
	Code         ErrNo         /* The error code returned by SQLite */
	ExtendedCode ErrNoExtended /* The extended error code returned by SQLite */
	SystemErrno  syscall.Errno /* The system errno returned by the OS through SQLite, if applicable */
	// contains filtered or unexported fields
}

func main() {
	fmt.Println("hello world")
}
