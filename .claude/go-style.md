# Go Style

Write Go that is idiomatic, simple, and easy to maintain.

- Prefer clear, straightforward control flow over clever abstractions.
- Follow standard Go formatting with `gofmt` and keep imports clean and organized.
- Prefer concrete types by default. Avoid interfaces for single-use code.
- Prefer the standard library. No external dependencies in algorithm solutions.
- Add comments only for non-obvious logic — hidden constraints, subtle invariants, or algorithm-specific tricks.
- Write tests that cover happy paths, edge cases, and boundary conditions. Use table-driven tests.
