# Golings Curriculum — Beginner → Advanced

149 exercises across 47 topics, ordered as a progressive track. Run them in
order with `mise run watch` (or `./bin/golings watch`). Each topic folder has a
`README.md` with links; each exercise gives a hint via `./bin/golings hint <name>`.

Toolchain: **Go 1.26** (mise-pinned). Covers language features through Go 1.26
(generics, `any`, `min`/`max`/`clear`, range-over-int, range-over-func iterators).

---

## 🟢 Beginner

### 1. Fundamentals
| Topic | Exercises | Concepts |
|---|---|---|
| variables | 6 | declaration, `:=`, zero values, constants |
| primitive_types | 5 | numeric types, `byte`, conversions |
| if | 2 | conditionals |
| switch | 3 | `switch`, conditionless switch |
| functions | 4 | params, multiple returns |
| more_functions | 2 | recursion, variadic + slice spread |
| strings | 2 | `strings` pkg, bytes vs runes (UTF-8) |

### 2. Collections & Loops
| Topic | Exercises | Concepts |
|---|---|---|
| arrays | 2 | fixed arrays, indexing |
| slices | 4 | `make`, slicing, `append` |
| maps | 3 | keys/values, comma-ok |
| range | 3 | iterating slices & maps |

---

## 🟡 Intermediate

### 3. Types & Methods
| Topic | Exercises | Concepts |
|---|---|---|
| structs | 3 | fields, embedding, methods |
| pointers | 3 | `&`/`*`, mutation, value vs pointer |
| methods | 2 | value vs pointer receivers, named types |
| interfaces | 4 | implicit satisfaction, Stringer, type switch, embedding |
| enums | 2 | `iota`, `String()` |

### 4. Functions & Errors
| Topic | Exercises | Concepts |
|---|---|---|
| anonymous_functions | 3 | closures |
| defer | 2 | LIFO, guaranteed cleanup |
| errors | 6 | values, `%w`, `Is`/`As`, custom types, `panic`/`recover`, `Join` |

### 5. Generics & Modern Go
| Topic | Exercises | Concepts |
|---|---|---|
| generics | 4 | type params, constraints, generic `Stack`, `Reduce` |
| modern | 5 | `min`/`max`/`clear` (1.21), range-over-int (1.22), range-over-func iterators (1.23), loop variable scoping (1.22) |
| type_aliases | 3 | alias vs defined type, generic type aliases (1.24) |
| iterators | 4 | `iter.Seq`, `iter.Seq2`, `iter.Pull` |
| maps_package | 2 | `maps.Keys`/`Values`, `maps.Collect` |

### 6. Testable Design
| Topic | Exercises | Concepts |
|---|---|---|
| dependency_injection | 3 | inject `io.Writer`, interface DI, constructor injection |
| mocking | 2 | spies, stubs/fakes (test doubles) |

---

## 🔴 Advanced

### 7. Concurrency
| Topic | Exercises | Concepts |
|---|---|---|
| concurrent | 3 | goroutines, `WaitGroup`, closing channels |
| channels | 2 | buffered, directional |
| select | 3 | multiplexing, timeouts, non-blocking |
| sync | 3 | `Mutex`, `Once`, `atomic` |
| context | 3 | cancellation, deadlines, values |
| concurrency_patterns | 3 | worker pool, fan-in, pipeline |
| goroutine_safety | 3 | leak-free goroutines, `errgroup`, cancellation discipline |
| synctest | 2 | `testing/synctest` fake clock (1.25) |

### 8. Standard Library & I/O
| Topic | Exercises | Concepts |
|---|---|---|
| stdlib_essentials | 7 | JSON, `io`, `slices`, `time`, `strconv`, `regexp`, `omitzero` (1.24) |
| reflection | 2 | `reflect.TypeOf`/`ValueOf`, `Kind`, walking struct fields |
| files | 2 | `os.ReadFile`/`WriteFile`, `bufio.Scanner` |
| http_client | 1 | `http.Get`, reading responses |
| structured_logging | 3 | `log/slog`, handlers, attributes |
| unsafe_pkg | 2 | `unsafe.Pointer`, `unsafe.Slice` — and when not to |

### 9. Building Applications
| Topic | Exercises | Concepts |
|---|---|---|
| http_server | 4 | `HandlerFunc`, `ServeMux` routing (1.22), JSON, middleware |
| http_server_advanced | 3 | method patterns, path wildcards, graceful shutdown |
| cli | 2 | `flag` parsing, `FlagSet`, positional args |
| profiling | 3 | `runtime/pprof`, `ReadMemStats`, `net/http/pprof` on a custom mux |

### 10. Testing & Applied
| Topic | Exercises | Concepts |
|---|---|---|
| testing_advanced | 4 | subtests, fuzzing, httptest, benchmarks (`b.Loop`) |
| applied | 2 | `sort.Interface`, concurrency-safe store (integration) |

### 11. Capstone
One program built across eight stages — a concurrent HTTP log-ingest service.
These are the only **package-mode** exercises: each spans several files in a
directory, and each ships with the earlier stages already solved.

| Topic | Exercises | Concepts |
|---|---|---|
| capstone | 8 | sentinel errors, `RWMutex`, method patterns, worker pool, graceful shutdown, `slog` middleware, fuzzing, `net/http/pprof` |

---

## Coverage notes

- **A Tour of Go** — fully covered.
- **Go by Example** — all core language + common stdlib topics covered. Deliberately
  omitted as out-of-scope for a language roadmap: XML, base64, SHA256, text/templates,
  TCP, spawning/exec processes, signals, URL parsing.
- **Go releases** — language syntax current through 1.26. 1.24–1.26 niche features
  (generic type aliases, `testing/synctest`, `new(expr)`, self-referential constraints)
  are intentionally not drilled.
