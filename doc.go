/*
Package goenv implements a GOPATH wrapper tool to isolate dependencies among many Go codebases.

This implementation utilizes a user-local directory and environment variables to track the
state of active and inactive GOPATHs on a system.

Environment

All known GOPATHs are stored within a directory in the current user's home directory.  This
path is

	~/.goenv

Each known GOPATH is stored within a config file named after the Go package.  Each config
is loaded when the package is "activated" by the user.  During activation, multiple
environment variables are modified to ensure that the Go toolchain will use the specified
package's GOPATH and installed dependencies outside of the vendor directory as well as ensure
that the user's system will be be able to resolve binaries in the activated GOPATH successfully.

Environment Variables

GOENV_PATH:
	Explanation
GOENV_GOPATH:
	Explanation
GOENV_PGOPATH:
	Explanation
GOENV_UNIXPATH:
	Explanation
GOENV_PUNIXPATH:
	Explanation
*/
package main
