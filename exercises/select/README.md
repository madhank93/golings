# Select

`select` lets a goroutine wait on multiple channel operations, proceeding with
whichever is ready. If several are ready it picks one at random; a `default`
case makes it non-blocking.

- **select1** — receive from whichever channel is ready
- **select2** — timeout by racing a receive against `time.After`
- **select3** — non-blocking receive with `default`

## Resources

- [Go by Example: Select](https://gobyexample.com/select)
- [Go by Example: Timeouts](https://gobyexample.com/timeouts)
- [Go by Example: Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations)
