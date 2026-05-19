# alGo

Algorithm implementations and solutions to competitive programming problems in Go.

## Sources

- [LeetCode](https://leetcode.com) — daily problems and topic practice
- [NeetCode](https://neetcode.io) — curated problem sets
- [Timus Online Judge](https://acm.timus.ru) — classic competitive programming problems

## Structure

```
internal/
  problems/
    leetcode/leetcode_<number>/   # solution.go + solution_test.go
    neetcode/<problem_name>/      # solution.go + solution_test.go
    timus/<number_name>/          # standalone main.go (stdin/stdout)
  themes/<topic>/                 # algorithm demos (binary search, BFS, sorting, …)
  utils/                          # shared I/O helpers for Timus problems
```

## Usage

```bash
# Run all tests
go test ./...

# Run tests for a specific problem
go test ./internal/problems/leetcode/leetcode_53/...

# Run a single test by name
go test ./internal/problems/leetcode/leetcode_53/... -run TestMaxSubArray

# Run a Timus problem
go run ./internal/problems/timus/1119_metro/main.go

# Run an algorithm theme demo
go run ./internal/themes/binarySearch/main.go
```
