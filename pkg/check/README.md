# The `check` Package

The `check` package is designed for performing assertions in Go tests,
particularly as a foundational layer for the `assert` package. It provides
functions that return errors instead of boolean values, allowing callers to
adjust error messages to particular context, add more contextual information
about the check, improving assertion message comprehension.
