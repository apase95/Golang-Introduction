# 🏆 Golang Competitive Syntax

Welcome to the **Golang Competitive Syntax** repository! This module is a comprehensive, hands-on guide designed to help developers master Go (Golang) from the ground up, with a strong focus on **Competitive Programming (CP)** and algorithmic problem-solving. 

Instead of traditional web-based tutorials, this workspace is structured around real-world CP environments, utilizing Fast I/O operations and file-based testing (`TEST.INP` & `TEST.OUT`).

---

## ✨ Core Topics Covered

This repository is divided into 12 logical sections, guiding you from basic syntax to advanced algorithms:

1. **🧱 Language Fundamentals:** 
   - Deep dive into Go's variable types, pointers, conditionals, loops, and function declarations.
2. **⚡ Fast I/O Mechanics:**
   - Standard `fmt.Scan` is often too slow for CP. Learn to optimize execution time using `bufio.Reader` and `bufio.Writer` to prevent Time Limit Exceeded (TLE) errors.
3. **🧰 Built-ins & Standard Libraries:**
   - Mastery of Go's powerful standard libraries (`math`, `strings`, `slices`, `sort`) and built-in functions (`make`, `append`, `panic`, `recover`).
4. **🧠 Advanced Data Structures:**
   - Implementation of complex data structures that do not exist natively in Go, including:
     - **Stacks, Queues & Deques** (The "Tricky" section).
     - **Priority Queues** (Using `container/heap`).
     - **Disjoint-Set Union (DSU)** for graph problems.
     - **Prefix Sums** for fast range queries.

---

## ⚙️ Prerequisites

- **Language:** [Go](https://golang.org/dl/) (version 1.18 or higher is recommended to support generic `slices` operations).
- A terminal (ZSH/Bash) and a code editor.

---

## 🚀 How to Run & Test

Unlike standard applications, every concept in this module is paired with a `TEST.INP` (Input) and `TEST.OUT` (Expected Output) file. 

The Go scripts are hardcoded to read inputs from `TEST.INP` and print results to the terminal (or write them to `TEST.OUT`), simulating a strict Competitive Programming judge system.

### Step 1: Navigate to a topic
```bash
cd 1-Golang-Competitive-Syntax
cd 12-Adv-Data-Structure
```

### Step 2: Modify the input
You can open `TEST.INP` and change the raw data to test different edge cases.
```bash
nano TEST.INP
```

### Step 3: Run the algorithm
Execute the specific `.go` file you want to test.
```bash
go run HelloWorld.go
```

---
## Happy Coding!