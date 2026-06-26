# Contributing to golings

Thank you for the interest on contributing to `golings`

Below you can find some useful information if you want to contribute to the project, be it opening a new issue or adding code.

## Adding an exercise

Two steps are required to add a new exercise to the project: adding the exercise metadata to the `info.toml` and creating the respective exercise in the `exercises` folder.

Add the metadata for your exercise in the correct order. Exercises are run in the order they are defined in the `info.toml` file. If you are unsure about the order, ask help in the pull request.

Here is an example of an exercise being added:

```toml
[[exercises]]
name = "compile1"
path = "compile/compile1.go"
mode = "compile"
hint = "hints are cool"
```

The exercise mode is very important. It tells `golings` how to run the exercise. If you are adding an exercise that expects the user to only make it compilable, use `compile` mode. If it has a test suite and you need the user to actually have the tests passing, use `test`.

### Describe the task in the file, not in `info.toml`

Each exercise file carries its own short instruction as a comment. This is the
**single source of truth** for the task: golings parses the first meaningful
comment line and shows it as the exercise's description in both the TUI and the
docs site, so there's nothing to keep in sync.

```go
// variables4
// Make me compile!
//
// See what happens when a block reuses a name with a new type.

// I AM NOT DONE
package main
```

Guidelines:

- Keep the first comment line a concise, **non-spoiler** one-liner (it becomes
  the catalog description). Add one or two extra comment lines below it only if
  the learner needs more direction — say *what* to fix, never the answer.
- The `// <name>` and `// Make me compile!` header lines are ignored, as is the
  `// I AM NOT DONE` marker.
- Only add a `desc = "..."` field in `info.toml` if you need to **override** the
  file comment (rarely needed). Prefer the file comment so intent lives in one
  place.

## Running the test suite

```sh
mise run test
```

## Issues and pull requests

There are specific templates that will guide you through opening issues or pull requests.

Thank you! ❤️
