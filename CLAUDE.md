# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

Algorithm implementations and solutions to competitive programming problems in Go. Sources include LeetCode, NeetCode, and Timus Online Judge.

## Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific problem
go test ./internal/problems/leetcode/leetcode_53/...

# Run a single test by name
go test ./internal/problems/leetcode/leetcode_53/... -run TestMaxSubArray

# Run a Timus/theme standalone program
go run ./internal/problems/timus/1119_metro/main.go

# Format code
gofmt -w .
```

## Code Architecture

```
internal/
  problems/
    leetcode/leetcode_<number>/   # LeetCode problems
    neetcode/<problem_name>/      # NeetCode problems
    timus/<number_name>/          # Timus Online Judge problems
  themes/<topic>/                 # Algorithm theme implementations (binary search, sorting, graphs, etc.)
  utils/                          # Shared utilities (file I/O, console I/O)
```

**LeetCode / NeetCode problems** use package-per-problem layout. Each problem directory contains:
- `solution.go` — the solution with a URL comment and Big-O annotations
- `solution_test.go` — table-driven tests

**Timus problems** are standalone `main` packages that read from stdin (or `INPUT.TXT`) and write to stdout (or `OUTPUT.TXT`). They do not have test files.

**Themes** are standalone `main` packages demonstrating a specific algorithm (e.g., lower/upper bound binary search, BFS, quicksort). No test files; they run directly via `go run`.

## Go Style

@.claude/go-style.md

Each solution file must also have a URL comment pointing to the problem, followed by Time/Space complexity annotations.

## Custom Commands

- `/go-review` — senior Go code review focused on correctness, idiomatic style, and test quality
- `/gen-tests` — generate a complete `solution_test.go` for the current problem from `solution.go`
