# humanize [![Test Status](https://github.com/imbue11235/humanize/workflows/Go/badge.svg)](https://github.com/imbue11235/humanize/actions?query=workflow:Go) [![codecov](https://codecov.io/gh/imbue11235/humanize/branch/main/graph/badge.svg?token=XTJ42655U1)](https://codecov.io/gh/imbue11235/humanize) [![Go Reference](https://pkg.go.dev/badge/github.com/imbue11235/humanize.svg)](https://pkg.go.dev/github.com/imbue11235/humanize)

A collection of utility functions, with built-in localization, for humanizing various types of data input.

## 🛠  Installation

Make sure to have Go installed (Version `1.16` or higher).

Install `humanize` with `go get`:

```sh
$ go get -u github.com/imbue11235/humanize
```

## 📖  Documentation

### Humanizing time

Takes two time instances, and presents the difference in a human-readable format.

#### Estimated time

This is a more loose calculation of time, where only the highest unit of time is prioritized.
E.g. `1 hour 20 minutes` becomes `1 hour` and `1 hour 55 minutes` becomes `2 hours` etc. 

##### From

Takes two time instances as input, to produce a human-readable representation of the difference in time.

```go
a := time.Parse(..., "2020-01-01")
b := time.Parse(..., "2021-02-01")

fmt.Printf("It happened almost %s", humanize.Time(a).From(b)) // => It happened almost a year ago
```

#### From now

This is a utility function, which is like calling `Time(a).From(b)`, but where `b` is automatically set to `time.Now()`

```go
a := time.Parse(..., "2021-05-05")

fmt.Printf("The file was created %s", humanize.Time(a).FromNow()) // => The file was created 5 days ago
```

#### To

The same as [From](#from), but the opposite time difference.

```go
a := time.Parse(..., "2020-01-01")
b := time.Parse(..., "2021-02-01")

fmt.Printf("It will happen %s", humanize.Time(a).To(b)) // => It will happen in a year
```

#### To now

A utility function like [From now](#from-now), which is like calling `Time(a).To(b)` where `b` is set to `time.Now()`

```go
a := time.Parse(..., "2021-05-05 22:10:00")

fmt.Printf("The plane will take off in %s", humanize.Time(a).ToNow()) // => The plane will take off in a minute
```

#### Precise time

A more precise calculation of time, where all time units is included.

```go
humanize.ExactTime(time).From(from) // => 1 hour, 23 minutes and 20 seconds ago
humanize.ExactTime(time).FromNow()  // => 3 years and 2 months ago
humanize.ExactTime(time).To(to)     // => in 6 years and 25 days
humanize.ExactTime(time).ToNow()    // => in 3 years, 6 months and 23 days
```

### Humanizing slices

Converts a string slice into a comma-separated string list with an optional limit.

```go
fmt.Printf("I went to the zoo with %s", humanize.Slice([]string{"Noah", "Marc"})) // => I went to the zoo with Noah and Marc
```

<details>
<summary markdown="span">Examples of usage ✨</summary>

```go
humanize.Slice([]string{"Joe"})                                // => Joe
humanize.Slice([]string{"Joe", "Leslie"})                      // => Joe and Leslie
humanize.Slice([]string{"Joe", "Leslie", "Carl"})              // => Joe, Leslie and Carl
humanize.Slice([]string{"Joe", "Leslie", "Carl"}, 2)           // => Joe, Leslie and one other
humanize.Slice([]string{"Joe", "Leslie", "Carl", "Yvonne"}, 2) // => Joe, Leslie and 2 others
humanize.Slice([]string{"Joe", "Leslie"}, 2)                   // => Joe and Leslie
```
</details>

### Localization

`humanize` comes prepacked with localization support, which are easily switchable on the fly.

#### Switching locales

To use a different locale with `humanize`, simply import it.
It auto-registers with the locale manager, and are ready to be used.

```go
import (
	"github.com/imbue11235/humanize"
	_ "github.com/imbue11235/humanize/locale/da"
)

func main() {
	names := []string{"Hans", "Viggo", "Klaus"}
	
	fmt.Printf("Seen by %s", humanize.Slice(names, 2)) // => Seen by Hans, Viggo and one other
	
	// switch the locale
	humanize.SetLocale("da")
	
	fmt.Printf("Set af %s", humanize.Slice(names, 2)) // => Set af Hans, Viggo og en anden
}
```

#### Registering custom locale

To register your own custom locale, simply specify a `locale.Map` mimicking the same key-value format
as the built-in locales. See [the english locale map](locale/en/locale.go) for additional details.

```go
humanize.RegisterLocale("mylocale", locale.Map{...})
```

#### Fallback locale