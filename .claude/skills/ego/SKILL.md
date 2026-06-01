```markdown
# ego Development Patterns

> Auto-generated skill from repository analysis

## Overview

This skill teaches you the core development patterns and workflows used in the `ego` Go codebase. You will learn about file organization, coding conventions, testing approaches, and efficient ways to perform large-scale refactors. This guide is especially useful for contributors aiming to maintain consistency and leverage automated workflows in the project.

## Coding Conventions

### File Naming

- Use **snake_case** for all file names.
  - Example: `elog_api.go`, `memory_encoder_test.go`

### Import Style

- Use **relative imports** within the project.
  - Example:
    ```go
    import "../core/econf"
    ```

### Export Style

- Use **named exports** for functions, types, and variables.
  - Example:
    ```go
    // Exported function
    func NewLogger() *Logger {
        // ...
    }
    ```

### Commit Patterns

- Mixed commit types, often starting with a prefix such as `refactor`.
- Commit messages are concise, averaging around 60 characters.
  - Example: `refactor: replace interface{} with any in all packages`

## Workflows

### Bulk Refactor Across Codebase

**Trigger:** When the codebase needs to be updated to use a new language feature or to comply with updated best practices (e.g., replacing `interface{}` with `any` in Go 1.18+).

**Command:** `/bulk-refactor interface-to-any`

**Step-by-step:**

1. **Identify** the deprecated or changed construct (e.g., `interface{}`).
2. **Search and replace** all occurrences across the codebase.
   - Example replacement:
     ```go
     // Before
     func DoSomething(v interface{}) {}

     // After
     func DoSomething(v any) {}
     ```
3. **Update related test files** to match the new usage.
4. **Commit changes** to all affected files with a clear message.
5. **Open a pull request** and merge after review.

**Files Involved:**
- Affects multiple files across `client/`, `core/`, `server/`, `internal/`, `examples/`, and `task/` directories.

**Frequency:** ~1x/month

## Testing Patterns

- **Testing framework:** Not explicitly detected; likely uses Go's built-in `testing` package.
- **Test file pattern:** Files containing `.test.` in their names.
  - Example: `interceptor_test.go`, `conf_test.go`, `memory_encoder_test.go`
- **Test structure:** Standard Go test functions.
  - Example:
    ```go
    func TestFunctionName(t *testing.T) {
        // test logic
    }
    ```

## Commands

| Command                        | Purpose                                                   |
|---------------------------------|-----------------------------------------------------------|
| /bulk-refactor interface-to-any | Perform a sweeping refactor from `interface{}` to `any`.  |
```
