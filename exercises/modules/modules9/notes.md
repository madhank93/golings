## modules9 — go.work workspaces

```
go 1.26.0

use (
    ./app
    ./greet
)
```

**Why it works**

- Every module named in `use` becomes a main module of the workspace. `app`
  still says `require example.com/greet v1.0.0`, but inside the workspace the
  local `./greet` wins over any version, so `greet.Formal` — which v1.0.0 does
  not have — resolves.

**Under the hood**

- The `go` command finds `go.work` by walking up from the current directory,
  like it does for `go.mod`. With only `./app` listed, greet is an ordinary
  dependency, and with `GOPROXY=off` the lookup for v1.0.0 fails at setup.
- `go work init ./app` creates the file and `go work use ./greet` adds a module
  (`go work use -r .` adds every module below a directory).

**Common mistake**

- Reaching for a `replace` in `app/go.mod` instead. It works locally, but it
  ships to everyone who builds `app`. `go.work` keeps the override on your
  machine; `GOWORK=off` turns it off for a one-off build.

**Key detail:** libraries usually do not commit `go.work` — CI and users should
build against the versions in `go.mod`, not your checkout layout. Commit it only
when the repo really is one workspace (a monorepo built as a unit).

**See also:** modules5 (`replace` to a local path) · the [chapter](../README.md)

**References**

- Go Modules Reference — Workspaces: https://go.dev/ref/mod#workspaces
- Go Modules Reference — `go work use`: https://go.dev/ref/mod#go-work-use
- Tutorial — Getting started with multi-module workspaces: https://go.dev/doc/tutorial/workspaces
- Go blog — Get familiar with workspaces: https://go.dev/blog/get-familiar-with-workspaces
