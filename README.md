# Golang-Introduction

This repository documents a comprehensive learning journey in Go (Golang). It transitions from basic language syntax and competitive programming techniques to advanced backend engineering, concurrency patterns, and the development of a real-world Command Line Interface (CLI) application.

## Repository Structure

The repository is divided into three main progressive modules:

### 1. Golang-Competitive-Syntax
This module focuses on the core mechanics of the Go programming language, specifically tailored for algorithmic problem-solving and competitive programming.
- **Language Fundamentals:** Variables, pointers, loops, conditionals, and functions.
- **Performance:** Fast I/O operations using `bufio` to prevent execution timeouts.
- **Data Structures:** Implementations of essential data structures such as Linked Lists, Stacks, Queues, Deques, Priority Queues (Heap), and Disjoint-Set Union (DSU).
- **Standard Library:** Practical usage of built-in packages like `math`, `strings`, `slices`, and `sort`.

### 2. Golang-Backend-Core
This module bridges the gap between basic scripting and production-ready software engineering. It covers backend service development and DevOps tooling.
- **Core Mechanics:** Interfaces, Generics, and Struct Embedding.
- **Advanced Concurrency:** In-depth exploration of Goroutines, Channels, and industry-standard concurrency patterns (Worker Pools, Pipelines, Mutex/Atomic, Context Timeouts).
- **Backend Fundamentals:** HTTP routing (Go 1.22+ `ServeMux`), middleware pipelines, JSON data handling, and graceful shutdowns.
- **Architecture & Database:** Implementing a clean MVC (Model-View-Controller) architecture, database connection pooling, and ORM integration using GORM with PostgreSQL.
- **DevOps CLI Tools:** Executing shell commands (`os/exec`), parsing command-line flags, and efficient file I/O operations.

### 3. My-CLI-Workspace
A capstone project applying the concepts learned in the previous modules. It is a minimalist, fully offline CLI application built using the `spf13/cobra` framework.
- **Task Management:** An interactive REPL mode to add, list, update, and delete daily tasks.
- **Expense Tracker:** A quick financial tracking tool to log expenses and calculate totals.
- **Local Music Player:** A headless media player integrating `mpv`, featuring real-time playback timers, OS-level MPRIS integration (Top Bar display), and Vim-style keyboard controls for seeking and navigation.

## Directory Tree

```text
.
├── 1-Golang-Competitive-Syntax
│   ├── 1-Variables-Types
│   ├── 2-Fast-IO
│   ├── 3-Conditional-Statement
│   ├── 4-Loop
│   ├── 5-Function
│   ├── 6-Built-in-Function
│   ├── 7-String-Processing
│   ├── 8-Standard-Library-Package
│   ├── 9-Map-Set
│   ├── 10-Struct
│   ├── 11-Tricky
│   └── 12-Adv-Data-Structure
├── 2-Golang-Backend-Core
│   ├── 1-Go-Core-Mechanics
│   ├── 2-Concurrency
│   ├── 3-Backend-Fundamentals
│   ├── 4-Basic-MVC-Architecture-ORM
│   └── 5-DevOps-CLI-Tools
├── 3-My-CLI-Workspace
│   ├── cmd
│   │   ├── expense.go
│   │   ├── music.go
│   │   ├── root.go
│   │   └── task.go
│   └── internal
│       ├── models
│       │   ├── expense.go
│       │   └── task.go
│       ├── services
│       └── storage
│           └── storage.go
└── README.md
```

---
## Happy Coding!!!