# 🧁 Sugabred

**Sugabred** is a powerful macOS CLI tool designed to **perfectly replicate a development machine** including all developer tools, system configurations, CLI utilities, environment variables, editors, and state while leaving out personal files.

The vision:  
> 🧠 **One command should turn any fresh Mac into your exact dev machine down to every installed tool, binary, and hidden config.**

Sugabred generates a complete snapshot of your machine in a shareable YAML format. You (or your team) can then use that snapshot to recreate a near-identical environment on another macOS system clean, fast and reliably.

---

## 🧭 Why Sugabred?

Modern developer environments are deeply customized: dozens of tools, version managers, plugins, environment variables, symlinks, configs, and more.

But when setting up a new machine or onboarding a teammate we start from scratch. Even with dotfiles and Brewfiles, things go missing. Editors behave differently. Commands fail silently.

Sugabred solves this by:
- Automatically detecting **everything** that makes your Mac dev-ready
- Creating a YAML snapshot that is **versioned, portable, and reproducible**
- Allowing restoration of that environment on another machine reliably

---

## ✨ Features

- ✅ Captures **Homebrew** packages and casks with versions  
- ✅ Records **environment variables**  
- ✅ Detects **macOS version, architecture, and timestamp**  
- ✅ Automatically discovers hidden directories like `.nvm`, `.pyenv`, `.ivy2`, etc.  
- ✅ CLI tools, dev tools, IDEs, shells, and system binaries  
- ✅ Leaves out private documents and personal files  
- ✅ Architecture-aware: adapts between Intel and Apple Silicon installs  

---

## 📦 Install

```bash
go install github.com/ajaypanthagani/sugabred@latest
````

Make sure `$GOPATH/bin` is in your `$PATH`.

---

## 🚀 Usage

### 📸 `sugabred snapshot`

Take a snapshot of your current macOS development environment and store it in `sugabred.snapshot.yaml`.

```bash
sugabred snapshot
```

### 🔼 `sugabred up`

Sets up dev environment based on the specifications in `sugabred.snapshot.yaml`.

```bash
sugabred up
```

### 🧪 `sugabred doctor` (coming soon)

Check if the current machine matches a provided snapshot and identify what's missing or different.

---

## 📂 Example Output

```yaml
timestamp: "2025-06-27T10:45:00Z"
architecture: "arm64"
macos_version: "14.4"
brew_packages:
  - name: go
    version: "1.22.1"
  - name: node
    version: "20.5.0"
brew_casks:
  - name: visual-studio-code
    version: "1.80.0"
env_vars:
  PATH: "/usr/local/bin:/opt/homebrew/bin"
  NODE_ENV: "development"
```

---

## 🛠 Tech Stack

* **Go**
* **Cobra** for CLI
* **Ginkgo + Gomega** for tests
* **Interfaces & Dependency Injection** for testability
* **Modular design** for extensibility (more collectors coming soon)

---

## 🧪 Testing

```bash
ginkgo ./...
```

Or run with:

```bash
go test ./... -v
```

---

## 📃 License

Apache 2.0. See [`LICENSE`](LICENSE) for details.

---

## 👨‍🍳 Author

Created with ❤️ by [Ajay Panthagani](mailto:ajaypanthagani321@gmail.com)

---

## 💡 Roadmap

* [ ] Restore environment from snapshot
* [ ] Support VS Code, JetBrains plugins
* [ ] Git config, SSH setup
* [ ] Dotfile syncing
* [ ] App preference backups (macOS plist)
* [ ] Cross-architecture software mapping
