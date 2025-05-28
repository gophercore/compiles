Compiles
===

Quickly check if a Go module compiles without performing a full build.

This tool helps identify compile-time issues in your Go code without producing binaries. **It also checks if test code compiles.**

## Usage

```
Usage of compiles:
  -exclude-tests
        do not check test-only package or test executable
  -tags string
        comma-separated list of build tags to include
```

Run the following command:

```
go run github.com/gophercore/compiles@latest ./...
```

The first run may take a while as it caches the latest version of `compiles`.

## Install

To permanently install this tool as `compiles`, run:

```
go install github.com/gophercore/compiles@latest
```
