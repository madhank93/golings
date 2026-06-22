# Channels (deeper)

Beyond basic send/receive: buffering and direction.

- **channels1** — buffered vs unbuffered (`make(chan T, n)`)
- **channels2** — directional types (`chan<- T` send-only, `<-chan T` receive-only)

See also the `concurrent`, `select`, and `concurrency_patterns` topics.

## Resources

- [A Tour of Go: Buffered Channels](https://go.dev/tour/concurrency/3)
- [Go by Example: Channel Buffering](https://gobyexample.com/channel-buffering)
- [Go by Example: Channel Directions](https://gobyexample.com/channel-directions)
