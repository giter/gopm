Go Package Manager 
=========================

Gopm (Go Package Manager) is a Go package manage and build tool for Go.

## This Version?

THIS is a specified version of gopm add `restore` command

## Commands

```
NAME:
   Gopm - Go Package Manager

USAGE:
   Gopm [global options] command [command options] [arguments...]

COMMANDS:
   list		list all dependencies of current project
   gen		generate a gopmfile for current Go project
   get		fetch remote package(s) and dependencies
   bin		download and link dependencies and build binary
   config	configure gopm settings
   run		link dependencies and go run
   test		link dependencies and go test
   build	link dependencies and go build
   install	link dependencies and go install
   clean	clean all temporary files
   update	check and update gopm resources including itself
   **restore**  restore remote package(s) to $GOPATH
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --noterm, -n		disable color output
   --strict, -s		strict mode
   --debug, -d		debug mode
   --help, -h		show help
   --version, -v	print the version
```

## License

This project is under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for the full license text.
