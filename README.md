# OpenClaude

A full-featured terminal UI for [Claude Code](https://docs.anthropic.com/en/docs/claude-code).

[![CI](https://github.com/johmara/openclaude/actions/workflows/release.yml/badge.svg)](https://github.com/johmara/openclaude/actions/workflows/release.yml)
[![GitHub Release](https://img.shields.io/github/v/release/johmara/openclaude)](https://github.com/johmara/openclaude/releases)
[![npm](https://img.shields.io/npm/v/@johmara/openclaude)](https://www.npmjs.com/package/@johmara/openclaude)

<!-- TODO: Add screenshot -->

## Installation

### Homebrew

```sh
brew install johmara/tap/openclaude
```

### Go

```sh
go install github.com/johmara/openclaude@latest
```

### npm / npx

```sh
npx @johmara/openclaude
```

## Prerequisites

OpenClaude requires the [Claude Code CLI](https://docs.anthropic.com/en/docs/claude-code) to be installed and authenticated. OpenClaude acts as a frontend and delegates all AI interactions to Claude Code.

## Features

- Streaming markdown rendering with syntax highlighting
- Tool call visualization (file edits, bash commands, etc.)
- Multiple concurrent sessions
- 3 built-in color themes
- Fuzzy file picker for attaching context
- Command palette for quick access to actions
- Sidebar with cost, token, and session tracking

## Keybindings

`Ctrl+X` is the **leader key** — press it, then a second key within 2 seconds.

| Key | Action |
|-----|--------|
| Enter | Send message |
| Esc | Cancel generation / close dialog |
| Ctrl+C ×2 | Quit (double-press) |
| Ctrl+K | Command palette |
| Ctrl+X s | Session switcher |
| Ctrl+X t | Theme picker |
| Ctrl+X f | File picker |
| Ctrl+X ? | Help |
| Ctrl+X n | New session |
| PgUp / PgDn | Scroll messages |

## Themes

- **Nord** (default)
- **Catppuccin Mocha**
- **Dracula**

Cycle through themes with `Ctrl+X t`.

## Configuration

| Variable | Description | Default |
|----------|-------------|---------|
| `CLAUDE_PATH` | Path to the Claude Code CLI binary | `claude` |

## Building from source

```sh
git clone https://github.com/johmara/openclaude.git
cd openclaude
make build
```

## License

MIT
