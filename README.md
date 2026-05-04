# Todo-cli

A small command-line todo app written in Go that works cross platform on: `Windows`, `Linux` and `macOS`

```
todo-cli/
├── build/           # compiled binaries
│
├── commands/
│   ├── add.go       # add a todo
│   ├── edit.go      # edit a todo
│   ├── help.go      # show help
│   ├── list.go      # list todos
│   ├── mark.go      # mark as done
│   ├── remove.go    # remove a todo
│   ├── unmark.go    # unmark as done
│
├── main.go          # entry point
├── go.mod           # module config
├── go.sum           # dependency checksums
├── README.md        # project info
└── .gitignore       # ignored files (includes build/)
```

## Commands

`add <task>`       Adds a task to the list

`rem <id>`         Removes a task from the list

`mark <id>`        Marks a task as done

`unmark <id>`      Removes the done mark

`edit <id>`        Updates an existing task

`list`             Lists all tasks (optionally filter by: `done` | `active`)

`help`             Shows available commands

---

## Installation

### 1. Install Go

Make sure you have Go installed (1.20+ recommended):

```bash
go version
```

If not installed, download it from: [https://go.dev/dl/](https://go.dev/dl/)

---

### 2. Clone

```bash
git clone https://github.com/your-username/todo-cli.git
cd todo-cli
```

---

### 3. Build

```bash
go build -o todo-cli
```

Go will automatically download all dependencies defined in `go.mod`.

---

### 4. Run

```bash
./todo-cli
```

**Windows:**

```bash
todo-cli.exe
```

---

### Optional: Run without building

```bash
go run .
```

## To do

* [x] Task editing using `edit` command
* [x] Filtering for list command
* [x] Switch to SQLite for storage
