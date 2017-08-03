# Toggl command line client

[Toggl](https://www.toggl.com/) is a powerful time tracking tool. This is a simple command line client for Toggl written in Go that allows you to track time without leaving the CLI.

## Installation

Make sure you have Go installed, and `$GOPATH` is set. Then install Toggl:

```
go get -v github.com/pro711/toggl
```

You should find a `toggl` executable in `$GOPATH/bin/`.

## Configure API Token

Obtain your API token from your Toggl [profile](https://www.toggl.com/app/profile) page. Then create a config file for Toggl under your home directory:

```
echo token: "your API token" > ~/.toggl.yaml
```

## Usage

- Start a new time entry with `toggl start [flags] description`. For example,

```
toggl start -P Study "Read Effective Go"
```

This starts a time entry "Read Effective Go" under project "Study".

- Check the status of the current time entry:

```
toggl status
```

- Stop the current time entry:

```
toggl stop
```
