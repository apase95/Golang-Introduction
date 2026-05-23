# 🛠️ My-CLI-Workspace

**My-CLI-Workspace** is a personal Command Line Interface (CLI) application built with **Go**. This project aims to provide a minimalist, efficient, and fully offline workspace, integrating daily essential tools for developers directly into the terminal (ZSH/Bash).

The system is divided into 3 independent modules, managing data through a local JSON storage system. This allows you to maintain maximum focus without ever leaving your keyboard.

---

## ✨ Core Features
### **📝 Task Management:** 
   - An isolated REPL (Interactive Mode) environment for quick task operations (add, list, toggle, delete).
   - Intuitive color-coded status display for better visibility.
<img width="962" height="651" alt="image" src="https://github.com/user-attachments/assets/5b3b3e86-78b5-46a5-bd79-09c18607ff74" />


### **💸 Expense Tracker:**
   - Lightning-fast personal finance tracking using simple commands.
   - Automatically calculates and displays the "Total Spent".
<img width="964" height="643" alt="image" src="https://github.com/user-attachments/assets/369eca63-902f-4812-862f-39ea4cd04114" />



### **🎵 Local Music Player:**
   - Headless CLI music player utilizing `mpv`.
   - **MPRIS Integration:** Displays song titles and supports media controls directly on the OS Top Bar (Linux/GNOME).
   - Features a real-time playback timer and Vim-style keyboard shortcuts for seamless navigation and seeking.
<img width="962" height="651" alt="image" src="https://github.com/user-attachments/assets/1fb0baa2-b39c-47d8-8847-52cec0902765" />



---

## ⚙️ Prerequisites

- **Language:** [Go](https://golang.org/dl/) (version 1.20 or higher).
- **Music Player:** `mpv` and `mpv-mpris` are required (For Linux/Fedora).
```bash
sudo dnf install mpv mpv-mpris
```

---

## 🚀 Installation & Setup

### Step 1: Build the Executable Binary
```bash
git clone https://github.com/apase95/Golang-Introduction.git
cd Golang-Introduction/3-My-CLI-Workspace
go mod tidy
go build -o workspace main.go
```

### Step 2: Initialize Data Directories
```bash
mkdir -p ~/Downloads/My-CLI-Workspace-Data/Storage
mkdir -p ~/Downloads/My-CLI-Workspace-Data/Music
```

### Step 3: Configure ~/.zshrc (Aliases & Functions)
```bash
nano ~/.zshrc
```
```bash
export WORKSPACE_BIN="/home/<YOUR-PATH>/GoLang-Introduction/3-My-CLI-Workspace/workspace"

open_task() {
    echo "============================= TASK MANAGEMENT MODE ==============================="
    echo "💻 Controls: add \"title\", list, done <id>, del <id>. Type 'q' or 'exit' to quit."
    echo "=================================================================================="
    while true; do
        printf "task> "
        read -r user_input        
        if [[ "$user_input" == "q" || "$user_input" == "exit" || "$user_input" == "quit" ]]; then
            echo "🔴 Exited Task mode."
            break
        fi
        if [[ -n "$user_input" ]]; then
            local clean_input=${user_input#task }            
            eval "$WORKSPACE_BIN task $clean_input"
        fi
    done
}

open_expense() {
    echo "================================== EXPENSE TRACKER MODE ================================="
    echo "💻 Controls: add <amount> \"note\", list, delete <id>, clear. Type 'q' or 'exit' to quit."
    echo "========================================================================================="
    while true; do
        printf "expense> "
        read -r user_input        
        if [[ "$user_input" == "q" || "$user_input" == "exit" || "$user_input" == "quit" ]]; then
            echo "🔴 Exited Expense mode."
            break
        fi
        if [[ -n "$user_input" ]]; then
            local clean_input=${user_input#expense }            
            eval "$WORKSPACE_BIN expense $clean_input"
        fi
    done
}

open_music() {
    local MUSIC_DIR="$HOME/Downloads/My-CLI-Workspace-Data/Music"
    local target_dir="$MUSIC_DIR"
    local flags=()
    for arg in "$@"; do
        if [[ "$arg" == -* ]]; then
            flags+=("$arg")
        else
            target_dir="$MUSIC_DIR/$arg"
        fi
    done
    $WORKSPACE_BIN music play "$target_dir" "${flags[@]}"
}
```
```bash
source ~/.zshrc
```

---

## 📖 Usage Guide

### 1. Task Management
Type `open_task` anywhere in your terminal to enter the Task Mode.
- `add "Task Description"`: Add a new task.
- `list`: View all tasks.
- `done <ID>`: Toggle the completion status of a task.
- `del <ID>`: Delete a task from the system.
- `q`: Exit Task Mode.


### 2. Expense Tracker
Type `open_expense` to enter the Financial Tracker Mode.
- `add <amount> "note"`: Log a new expense (e.g., add 50000 "Morning Coffee").
- `list`: View your expense history and the total amount spent.
- `delte <ID>`: Delete a expense from the system.
- `clear`: Delete all expenses from the system.
- `q`: Exit Expense Mode.

### 3. Local Music Player
Use the `open_music` command followed by an optional subfolder name and shuffle flag.
- `open_music`: Play all tracks in the root Music directory.
- `open_music -s`: Play in random order (Shuffle).
- `open_music Pop`: Play tracks inside the Pop subfolder.
In-App Playback Controls:
- `Enter or n`: Skip to the Next track.
- `p`: Go back to the Previous track.
- `k`: Pause / Resume playback.
- `l` / `j`: Seek forward 5s / Seek backward 5s.
- `xx:yy`: Seek to an exact timestamp (e.g., `01:25`).
- `q` : Stop the music and exit the player.

---
### Happy Coding!
