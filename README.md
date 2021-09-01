# humanize

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
Time(time).From(from) // 1 year ago
Time(time).FromNow()  // 20 minutes ago
Time(time).To(to)     // in 3 years
Time(time).ToNow()    // in 60 minutes
```

#### Precise time

A more precise calculation of time, where all time units is included.

```go
ExactTime(time).From(from) // 1 hour, 23 minutes, 20 seconds ago
ExactTime(time).FromNow()  // 3 years and 2 months ago
ExactTime(time).To(to)     // in 6 years and 25 days
ExactTime(time).ToNow()    // in 3 years, 6 months and 23 days
```