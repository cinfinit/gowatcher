# GoWatcher: Because Your F5 Key Deserves a Break 🎯

[![Go Version](https://img.shields.io/github/go-mod/go-version/cinfinit/gowatcher)](https://github.com/cinfinit/gowatcher)
[![Dev Details](https://img.shields.io/badge/dev_details-green)](https://pkg.go.dev/github.com/cinfinit/gowatcher)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> *The embeddable live-reload library that doesn't make you install yet another CLI tool.*

---

## What Fresh Heck is This?

Let's face it: you're tired of:
- Installing existing CLIs on every machine you work on
- Your teammate using version 1.2.3 while you're on 1.4.0 and "it works on my machine" 
- Managing external processes that sometimes zombie out
- Explaining to new devs "oh, you need to install this CLI tool first"

**GoWatcher** is the SIMPLE. It's a **library**, not a CLI tool. You import it, you use it, it works. No external tools, no version mismatches, no separate processes to babysit.

---

## Why Does This Exist? (A Tragedy in Three Acts)

### Act I: The Problem
You're building a Go app. You make a change. You compile. You run. You repeat. Your F5 key is wearing out. You need live reload.

### Act II: The Existing "Solutions"
You install existing CLIs. Great! But wait...
- "What version are you using?" 
- "It doesn't work on my machine."
- "Why is there a `.newconfig.toml` file in my repo?"
- "The process didn't terminate, let me kill -9 everything..."

### Act III: THE Solution
**What if your app just... reloaded itself?** No external tools. No config files. Just import and go.

---

## How It Works (In Plain English)

1. **You import it**: `import "github.com/cinfinit/gowatcher"`
2. **You call it**: `gowatcher.Watch(".")` in your `main()`
3. **Magic happens**: In dev mode (`-tags dev`), it watches your files
4. **You save a file**: It rebuilds your app and restarts it
5. **Profit**: No more manual recompiling

**In production?** It's a no-op. Your binary doesn't even know it exists. Zero bloat. Nada. Zilch.

---

## The "Why Should I Care?" Section

### ✅ Pros (The Good Stuff)

- **No CLI tool roulette**: Pin it in `go.mod`, version it, forget it
- **Zero production overhead**: `//go:build dev` tags mean it doesn't exist in prod builds
- **Programmable**: Add hooks, customize build args, do whatever you want
- **Self-contained**: No separate process to manage, no zombie processes
- **Team-friendly**: Everyone gets the same version automatically
- **Cross-platform**: Works on Windows, Mac, Linux (thanks, `fsnotify`)

### ❌ Cons (The "Fine, I Guess" Stuff)

- **It's Go-only**: Shocking, I know , Actually it's not , it's just for Go files , so not a shocking thing :|
- **Build tags**: You need to remember `-tags dev` (but that's it!)

---

## Quick Start (Literally 30 Seconds)

### Step 1: Get It
```bash
go get github.com/cinfinit/gowatcher@latest
```

### Step 2: Use It
```go
package main

import (
    "net/http"
    "github.com/cinfinit/gowatcher"
)

func main() {
    // Start the watcher (only active with -tags dev)
    gowatcher.Watch(".")
    
    // Your actual app
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hot reload is pretty cool, ngl"))
    })
    
    http.ListenAndServe(":8080", nil)
}
```

### Step 3: Run It
```bash
# Development (with live reload)
go run -tags dev .

# Production (no watcher, no bloat)
go run .
```

### Step 4: Edit Something
Change your code, save the file, and **BOOM** — automatic reload. No F5 required.

---

## Advanced Wizardry

Want to feel like a hacker? Use the config option:

```go
gowatcher.WatchWithConfig(gowatcher.Config{
    Dir:       "./cmd/myapp",       // Watch a specific directory
    OnReload:  func() {            // Do stuff before reload
        fmt.Println("Ooh, hot reload! 🔥")
    },
    BuildArgs: []string{"-ldflags", "-s -w"}, // Strip debug info
    RunArgs:   []string{"--port=8080"},       // Args to your app
})
```

---


## FAQ (The Questions Nobody Asked But Here They Are)

### Q: Will this make my production binary huge?
**A**: Nope! The `watch.go` file (compiled in production) is literally an empty function. It's like it doesn't even exist.

### Q: What about Windows?
**A**: Works fine. It even uses `.exe` for the temp binary. We thought of everything.

### Q: Can I use this with Docker?
**A**: Absolutely! Just remember to pass `-tags dev` in your Dockerfile if you want hot reload in containers.

### Q: Does it watch `.go` files only?
**A**: Currently watches `.go`, `.mod`, and `.sum` files. Why? Because those are the ones that matter.

### Q: I'm a Vim/Emacs/VSCode user. Will it work?
**A**: Yes. It watches files, not editors. We're not biased (much).

### Q: Can I contribute?
**A**: Please do! The code is simple, the idea is solid, and we'd love your help.

---

## The "Don't Be Boring" Section

Still reading? Cool. Here's what makes GoWatcher special:

- **It's not another CLI tool**: We're solving the "which version of existing CLI?" problem
- **It's lean**: Around 150 lines of actual code in dev mode — clean, readable, no bloat
- **It's idiomatic Go**: Build tags, context cancellation, `defer` statements — we use all the tricks
- **It's SIMPLE**: Yeah it's SIMPLE , nothing fancy.

---

## License

MIT — because we're not monsters.

---

## Support

- Found a bug? Open an issue
- Want a feature? Open an issue (or PR, we're not picky)
- Want to say thanks? Star the repo ⭐

---

*GoWatcher: Because your time is better spent coding than pressing F5.* 🚀

---

## About the Author

This library was created by a developer [cinfinit](https://github.com/cinfinit) who got tired of:
- Manually restarting apps after every code change
- Explaining to teammates how to install yet another CLI tool
- Watching some random CLI crash and leaving zombie processes everywhere
- Googling "how to kill a stuck Go process" at 2 AM

The author is probably right now debugging something that "worked yesterday" and questioning their life choices.

**Author's mantra**: "Why press F5 when your code can press itself?"

Follow for more libraries born from developer's midnight coding sessions and questionable life decisions.

---

*Made with 🍕, `go build -tags dev`, and the deep-seated belief that pressing F5 should be illegal.*
