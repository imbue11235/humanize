# humanize [![Test Status](https://github.com/imbue11235/humanize/workflows/Go/badge.svg)](https://github.com/imbue11235/humanize/actions?query=workflow:Go) [![codecov](https://codecov.io/gh/imbue11235/humanize/branch/main/graph/badge.svg?token=XTJ42655U1)](https://codecov.io/gh/imbue11235/humanize) [![Go Reference](https://pkg.go.dev/badge/github.com/imbue11235/humanize.svg)](https://pkg.go.dev/github.com/imbue11235/humanize)

## Installation

```sh
$ go get github.com/imbue11235/humanize
```

## Usage

### Time

Converts a time difference (duration) into a more readable format

#### Approximate time

This is a more loose calculation of time, where only the highest unit of time is prioritized.
E.g. `1 hour 20 minutes` becomes `1 hour` and `1 hour 55 minutes` becomes `2 hours` etc. 

```go
humanize.Time(time).From(from) // => 1 year ago
humanize.Time(time).FromNow()  // => 20 minutes ago
humanize.Time(time).To(to)     // => in 3 years
humanize.Time(time).ToNow()    // => in 60 minutes
```

#### Precise time

A more precise calculation of time, where all time units is included.

```go
humanize.ExactTime(time).From(from) // => 1 hour, 23 minutes, 20 seconds ago
humanize.ExactTime(time).FromNow()  // => 3 years and 2 months ago
humanize.ExactTime(time).To(to)     // => in 6 years and 25 days
humanize.ExactTime(time).ToNow()    // => in 3 years, 6 months and 23 days
```

### Slice

Converts a string slice into a comma-separated string list with an optional limit.

```go
humanize.Slice([]string{"Joe"})                                // => Joe
humanize.Slice([]string{"Joe", "Leslie"})                      // => Joe and Leslie
humanize.Slice([]string{"Joe", "Leslie", "Carl"})              // => Joe, Leslie and Carl
humanize.Slice([]string{"Joe", "Leslie", "Carl"}, 2)           // => Joe, Leslie and one other
humanize.Slice([]string{"Joe", "Leslie", "Carl", "Yvonne"}, 2) // => Joe, Leslie and 2 others
humanize.Slice([]string{"Joe", "Leslie"}, 2)                   // => Joe and Leslie
```