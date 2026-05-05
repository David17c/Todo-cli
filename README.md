# Todo-cli

A small command-line todo app written in Go that works cross platform on: `Windows` and `Linux`.

```
todo-cli/
├── build/           # compiled binaries
│
├── commands/
│   ├── add.go       # add a new todo
│   ├── edit.go      # edit an existing todo
│   ├── help.go      # display CLI usage and available commands
│   ├── list.go      # list all todos
│   ├── mark.go      # mark a todo as completed
│   ├── remove.go    # remove a todo
│   ├── clear.go     # remove all todos
│   ├── unmark.go    # mark a todo as not completed
│
├── cli/
│   ├── parser.go    # parse input and dispatch commands
│
├── interactive/
│   ├── repl.go      # interactive REPL for managing todos
│
├── main.go          # application entry point
├── go.mod           # module definition and dependencies
├── go.sum           # dependency checksums
├── README.md        # project documentation and usage
└── .gitignore       # List of files Git ignores
```

## Commands

`add <task>`       Adds a task to the list

`rem <id>`         Removes a task from the list

`clear`            Removes all tasks at once

`mark <id>`        Marks a task as done

`unmark <id>`      Removes the done mark

`edit <id>`        Updates an existing task

`list`             Lists all tasks (optionally filter by: `done` | `active`)

`help`             Shows available commands

---

## Installation (Recommended way)

Go to: [Releases](https://github.com/David17c/todo-cli/releases) and look for the newest version made for your OS.

## Installation

#### 1. Install Go

Make sure you have Go installed (1.20+ recommended):

```bash
go version
```

If not installed, download it from: https://go.dev/doc/install

---

#### 2. Clone

If you have [Git](https://git-scm.com/) installed run.
```bash
git clone https://github.com/David17c/todo-cli.git
```
Or if you don't

Go to [The main page](https://github.com/David17c/todo-cli) of this repo, click the green code button and select "Download as ZIP", extract the folder then continue.

---

#### 3. Build

Use the `cd` command to navigate to the root folder.

Then run.

```bash
go build -o build
```
This will create an executable file in the build folder.

---

#### 4. Run

From the root of the project on Linux run:
```bash
./build/todo-cli
```

From the root of the project on Windows run:
```bash
./build/todo-cli.exe
```

---
## To do

* [x] Task editing using `edit` command.
* [x] Filtering for list command.
* [x] Switch to SQLite for storage.
* [x] Add a clear command that removes all tasks.

> **_NOTE:_**  Todo-cli probably works on macOS but has not been tested since i don't own one.
