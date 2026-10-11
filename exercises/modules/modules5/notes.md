## modules5 — require + replace a local module

```text
// go.mod
module example.com/app

go 1.26.0

require example.com/mathx v0.0.0

replace example.com/mathx => ./mathx
```

**Why it works**

- `require` says the build depends on module `example.com/mathx`; `replace`
  says where its source is. With the right side a directory, the go command
  reads `./mathx/go.mod` and the packages beside it instead of downloading
  anything.

**Under the hood**

- Every imported package must be provided by a module in the requirement
  graph, so `replace` on its own is not enough — without the `require` the
  build fails with `no required module provides package example.com/mathx`. The
  version in `require` is a placeholder: a directory replacement has no
  version, and `v0.0.0` is the conventional stand-in.

**Common mistake**

- Pointing `replace` at a directory whose `go.mod` declares a different path.
  The target must be a module (it needs its own `go.mod`), and its `module`
  line must match the path being replaced, or the build stops with
  `module declares its path as ... but was required as ...`.

**Key detail:** `replace` is honoured only in the **main module's** `go.mod` —
the one you run the go command from. Replacements in dependencies are ignored,
so a module that relies on a local `replace` cannot be consumed as-is by
others; it is a development tool, not a publishing mechanism.

**See also:** modules1 (go.mod) · modules2 (import paths) · the
[chapter](../README.md)

**References**

- Go Modules Reference — require directive: https://go.dev/ref/mod#go-mod-file-require
- Go Modules Reference — replace directive: https://go.dev/ref/mod#go-mod-file-replace
- Managing dependencies — requiring module code in a local directory: https://go.dev/doc/modules/managing-dependencies#local_directory
