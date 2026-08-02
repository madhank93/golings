## Keys

| key | does |
| --- | --- |
| `↑` `↓` / `k` `j` | move between exercises |
| `↵` | run the current exercise |
| `e` | open it in `$EDITOR`, then re-run on exit |
| `h` | hint · `x` explain (the teaching walk-through) |
| `r` | reset the exercise to its starting state |
| `n` | jump to the next unsolved exercise |
| `/` | search · `esc` clear |
| `esc` | cancel a wedged run |
| `?` | this help · `q` quit |

## While a run is going

The badge next to the exercise name shows a spinner and the elapsed time. A run
that has wedged — an infinite loop, a deadlocked test — looks exactly like a
slow one without that clock, so watch the seconds rather than the spinner.

`esc` cancels the run without killing the terminal.

## Saving a file re-runs it

The watcher re-verifies on every save, so the usual loop is: edit, save, watch
the badge. You rarely need `↵` at all.
