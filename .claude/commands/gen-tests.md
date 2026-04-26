Read the solution.go file in the current problem directory and generate a complete `solution_test.go` file with table-driven tests.

## Steps

1. Read `solution.go` to understand:
   - The package name
   - All exported function signatures and their parameter/return types
   - Any types defined in the file (e.g. `TreeNode`, structs, or constructor types)
   - The problem URL from the comment

2. Generate `solution_test.go` following these rules:

### Test structure rules

- Use Go's standard `testing` package — no third-party libraries
- One `Test<FunctionName>` function per exported function or constructor
- Each test function uses a table-driven `tests := []struct{ ... }` pattern
- The loop must be `for _, tt := range tests { t.Run(tt.name, func(t *testing.T) { ... }) }`
- Failure messages use `t.Errorf("<funcName>() = %v, want %v", got, tt.want)`
- Use descriptive `name` strings that explain the scenario (e.g. `"all negatives"`, `"single element"`, `"empty input"`)

### Test case coverage rules

- Always include the examples from the problem URL (fetch if needed)
- Cover: happy path, edge cases (empty input, single element, all-same values, min/max constraints), and cases that stress the algorithm's logic
- Aim for 4–6 test cases per function

### Type-specific rules

**Simple function** (scalar/slice inputs, scalar output):
```go
tests := []struct {
    name   string
    <param> <type>
    want   <returnType>
}{ ... }
```

**Tree problems** (uses `*TreeNode`):
- Construct trees inline using nested `&TreeNode{Val: x, Left: ..., Right: ...}` literals
- `nil` root is always one test case

**Constructor / method receiver** (e.g. `Constructor(...)` returning a struct with methods):
- Do not use a table for the struct; instead write sequential sub-tests or a single test that calls the constructor and exercises each method with `t.Run` blocks
- Example pattern:
```go
func TestNumArray(t *testing.T) {
    t.Run("example from problem", func(t *testing.T) {
        obj := Constructor([]int{1, 3, 5})
        if got := obj.SumRange(0, 2); got != 9 {
            t.Errorf("SumRange() = %v, want %v", got, 9)
        }
        obj.Update(1, 2)
        if got := obj.SumRange(0, 2); got != 8 {
            t.Errorf("SumRange() = %v, want %v", got, 8)
        }
    })
}
```

### Style rules

- Same package as `solution.go` (no `_test` suffix)
- Imports: only `"testing"` unless additional packages are genuinely needed
- No comments inside test cases
- Inline struct literals, not named variables, for test inputs
- Do not reproduce the solution logic in the test

## Output

Write the generated tests directly to `solution_test.go` in the same directory as `solution.go`. Then run `go test ./...` scoped to that package to verify the tests compile and pass.
