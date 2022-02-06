# fun: a functional programming library for (generic) Go

A functional programming library for Go 1.18+. Currently mostly inspired by array functions in
JavaScript, but will likely grow over time.

> Yes, there are many libraries out there like this. This is mostly a learning exercise for me.

## Current API

```go
// Any types
func ForEach[T any]([]T, func(T))
func Map[T any]([]T, func(T) T) []T
func Filter[T any]([]T, func(T) bool) []T
func Concat[T any]([]T, []T) []T
func Every[T any]([]T, func(T) bool) bool
func Any[T any]([]T, fn func(T) bool) bool
func Reduce[T any]([]T, fn func(T, T) T) T
func Reverse[T any]([]T) []T
func Pop[T any]([]T, int) (T, []T)
func Push[T any]([]T, T) []T

// Comparable types
func Find[T comparable]([]T, func(T) bool) *T
func Includes[T comparable]([]T, T) bool
func Index[T comparable]([]T, T) int
func LastIndexOf[T comparable]([]T, T) int
func Replace[T comparable]([]T, T, T) []T
func ReplaceAll[T comparable]([]T, T, T) []T
func Delete[T comparable]([]T, T) []T
func DeleteAll[T comparable]([]T, T) []T

// Ordered types
func Max[T Ordered]([]T) T
func Min[T Ordered]([]T) T
```
