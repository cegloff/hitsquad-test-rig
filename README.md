# hellotool

A simple CLI tool to greet users.

## Installation

Build the tool using Go:

```bash
go build -o hellotool main.go
```

## Usage

### Root Command
Run the tool without arguments to see the version.

```bash
./hellotool
```
Output: `hellotool v0.1.0`

### Version Flag
You can also use the version flag:

```bash
./hellotool --version
```
Output: `hellotool v0.1.0`

### Greet Subcommand
Greeting the world (default):

```bash
./hellotool greet
```
Output: `Hello, world!`

Greeting a specific person:

```bash
./hellotool greet --name Alice
```
Output: `Hello, Alice!`
