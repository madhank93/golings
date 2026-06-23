# Slices

A slice is a dynamically-sized, flexible view into an underlying array. Create
with `make([]T, len, cap)` or a literal, grow with `append`, and take sub-slices
with `s[low:high]`. Slices have a length and a capacity.

## Resources

- [Go blog: Slices intro](https://go.dev/blog/slices-intro)
- [A Tour of Go: Slices](https://go.dev/tour/moretypes/7)
- [Go by Example: Slices](https://gobyexample.com/slices)
