# humanize [![Test Status](https://github.com/imbue11235/humanize/workflows/Go/badge.svg)](https://github.com/imbue11235/humanize/actions?query=workflow:Go) [![codecov](https://codecov.io/gh/imbue11235/humanize/branch/main/graph/badge.svg?token=XTJ42655U1)](https://codecov.io/gh/imbue11235/humanize) [![Go Reference](https://pkg.go.dev/badge/github.com/imbue11235/humanize.svg)](https://pkg.go.dev/github.com/imbue11235/humanize)

> A collection of utility functions, with [built-in localization](#built-in-locales), for humanizing various types of data input.

## 🛠  Installation

Make sure to have Go installed (Version `1.16` or higher).

Install `humanize` with `go get`:

```sh
$ go get -u github.com/imbue11235/humanize
```

## 💻  Usage

### Humanizing time

Takes two time instances, and presents the difference in a human-readable format.

#### Estimated time

This is a more loose calculation of time, where only the highest unit of time is prioritized.
E.g. `1 hour 20 minutes` becomes `1 hour` and `1 hour 55 minutes` becomes `2 hours` etc. 

##### Time from _x_

Takes two time instances as input, to produce a human-readable representation of the difference in time.

```go
a := time.Parse(..., "2020-01-01")
b := time.Parse(..., "2021-02-01")

fmt.Printf("It happened almost %s", humanize.Time(a).From(b)) // => It happened almost a year ago
```

##### Time from now

This is a utility function, which is like calling `Time(a).From(b)`, but where `b` is automatically set to `time.Now()`

```go
a := time.Parse(..., "2021-05-05")

fmt.Printf("The file was created %s", humanize.Time(a).FromNow()) // => The file was created 5 days ago
```

##### Time to _x_

The same as [From](#time-from-x), but the opposite time difference.

```go
a := time.Parse(..., "2020-01-01")
b := time.Parse(..., "2021-02-01")

fmt.Printf("It will happen %s", humanize.Time(a).To(b)) // => It will happen in a year
```

##### Time to now

A utility function like [FromNow](#time-from-now), which is like calling `Time(a).To(b)` where `b` is set to `time.Now()`

```go
a := time.Parse(..., "2021-05-05 22:10:00")

fmt.Printf("The plane will take off %s", humanize.Time(a).ToNow()) // => The plane will take off in a minute
```

#### Exact time

A more precise calculation of time, where all time units is included.

##### Exact time from _x_

Takes two time instances as input, to produce a string-representation of the exact difference in time.

```go
a := time.Parse(..., "2020-01-01 22:00:05")
b := time.Parse(..., "2021-02-01 02:50:22")

fmt.Printf("It happened exactly %s", humanize.ExactTime(a).From(b)) // => It happened exactly 1 year, 30 days, 4 hours, 45 minutes and 22 seconds ago
```

##### Exact time from now

This is a utility function, which is like calling `ExactTime(a).From(b)`, but where `b` is automatically set to `time.Now()`.

```go
a := time.Parse(..., "2021-06-06 22:05:05")

fmt.Printf("The file was deleted %s", humanize.ExactTime(a).FromNow()) // => The file was deleted 5 minutes and 5 seconds ago
```

##### Exact time to _x_

The same as [From](#exact-time-from-x), but the opposite time difference.

```go
a := time.Parse(..., "2021-05-03 15:00:00")
b := time.Parse(..., "2021-05-08 18:30:00")

fmt.Printf("It's my birthday %s", humanize.ExactTime(a).To(b)) // => It's my birthday in 5 days, 3 hours and 30 minutes
```

##### Exact time to now

A utility function like [FromNow](#exact-time-from-now), which is like calling `ExactTime(a).To(b)` where `b` is set to `time.Now()`

```go
a := time.Parse(..., "2021-03-02 12:00:33")

fmt.Printf("The train will depart %s", humanize.ExactTime(a).ToNow()) // => The train will depart in 2 minutes and 33 seconds
```
---

### Humanizing duration

#### Estimated duration

Similar to [estimated time](#estimated-time), this is a loose calculation of given duration.

```go
fmt.Printf("My dog is %s old", humanize.Duration(time.Hour * 24 * 68)) // => My dog is 2 months old
```
<details>
<summary>Additional examples of usage</summary>

```go
humanize.Duration(time.Hour * 24 * 38)  // => a month
humanize.Duration(time.Hour * 24 * 400) // => a year
humanize.Duration(time.Hour * 24 * 800) // => 2 years
```
</details>

#### Exact duration

Similar to [exact time](#exact-time), this is a strict and precise calculation of the given duration.

```go
fmt.Printf("The offer ends in %s", humanize.ExactDuration(time.Hour * 70)) // => The offer ends in 2 days and 22 hours
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.ExactDuration(time.Hour * 70)                                // => 2 days and 22 hours
humanize.ExactDuration(time.Hour*3 + time.Minute*33 + time.Second*55) // => 3 hours, 33 minutes and 55 seconds
humanize.ExactDuration(time.Hour * 24 * 8)                            // => 8 days
```
</details>

---

### Humanizing slices

Converts a string slice into a comma-separated string list with an optional limit.

```go
fmt.Printf("I went to the zoo with %s", humanize.Slice([]string{"Noah", "Marc"})) // => I went to the zoo with Noah and Marc
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.Slice([]string{"Joe"})                                // => Joe
humanize.Slice([]string{"Joe", "Leslie"})                      // => Joe and Leslie
humanize.Slice([]string{"Joe", "Leslie", "Carl"})              // => Joe, Leslie and Carl
humanize.Slice([]string{"Joe", "Leslie", "Carl"}, 2)           // => Joe, Leslie and one other
humanize.Slice([]string{"Joe", "Leslie", "Carl", "Yvonne"}, 2) // => Joe, Leslie and 2 others
humanize.Slice([]string{"Joe", "Leslie"}, 2)                   // => Joe and Leslie
```
</details>

---

### Humanizing sizes

Transforms byte-sizes into the closest related multi-byte unit size (MB, GB etc.)

#### Bytes

Uses the SI prefixes (powers of `10`, e.g. `1000b = 1kB`) for converting the bytes into their human-readable representation.

```go
fmt.Printf("The size of 'cats.jpg' is %s", humanize.Bytes(2500000)) // => The size of 'cats.jpg' is 2.5 MB 
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.Bytes(2000000000)          // => 2.0 GB
humanize.Bytes(1000000000000)       // => 1.0 TB
humanize.Bytes(1000000000000000000) // => 1.0 EB
```
</details>

#### Binary bytes

Uses the binary system (powers of `2`, e.g. `1024b = 1KiB`) for converting the bytes into their human-readable representation.

```go
fmt.Printf("The size of 'dogs.jpg' is %s", humanize.BinaryBytes(2500000)) // => The size of 'dogs.jpg' is 2.4 MiB
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.BinaryBytes(2500000)          // => 2.4 MiB
humanize.BinaryBytes(10000000)         // => 9.5 MiB
humanize.BinaryBytes(1000000000000000) // => 909 TiB
```
</details>

#### Short-form binary bytes

Using the same system as [binary bytes](#binary-bytes), but in a GNU-like format.

```go
fmt.Printf("vacation.zip | %s", humanize.ShortFormBinaryBytes(1000000000000)) // => vacation.zip | 931G
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.ShortFormBinaryBytes(35324355)             // => 34M
humanize.ShortFormBinaryBytes(2000000000)           // => 1.9G
humanize.ShortFormBinaryBytes(13400000000000000000) // => 12E
```
</details>

---

### Humanizing fractions

Transforms a float value into a formatted human-readable fraction

```go
fmt.Printf("You can have %s of the cake", humanize.Fraction(0.25)) // => You can have 1/4 of the cake
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.Fraction(2.625) // => 2 5/8
humanize.Fraction(0.5)   // => 1/2
humanize.Fraction(1.66)  // => 1 33/50
```
</details>

---

### Humanizing fuzzy text

Sometimes, being able to print out information to the user, directly from a data structure, a key 
in a JSON object or similar, is nice, instead of redefining it all, word for word.

#### Text

Formats a fuzzy text as a common sentence, capitalizing the first letter of the first word, and lower-casing the rest.

```go
fmt.Print(humanize.FuzzyText("some-!!@@----Wierd_____format")) // => Some wierd format
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.FuzzyText("my.key")                  // => My key
humanize.FuzzyText("snake-case")              // => Snake case
humanize.FuzzyText("a_text_with_underscores") // => A text with underscores
```
</details>

#### Custom format

Extracts words from a fuzzy text and constructs a string from the words, using the provided formatter on every extracted word.
If the formatter is `nil`, the words will be concatenated in their natural state.

```go
fmt.Printf("The receipt contains your %s", humanize.FormatFuzzyText("customer__id", strings.ToUpperCase)) // => The receipt contains your CUSTOMER ID
```

<details>
<summary>Additional examples of usage</summary>

With custom formatter:
```go
myCustomFormatter := func(index int, word string) {
    if index == 1 {
        return strings.ToUpperCase(word)
    }

    return strings.Title(word)
}

humanize.FormatFuzzyText("app-id", myCustomFormatter) // => App ID
```

With pre-defined formatter:
```go
humanize.FormatFuzzyText("app-id", strings.Title) // => App Id
```

</details>

---

### Humanizing numbers

Converting small or larger numbers into a shorter form of the number to a human-friendly text representation

#### Int

Converts an integer into a readable string representation, rounding the volume and using the [names of large numbers](https://en.wikipedia.org/wiki/Names_of_large_numbers) as a suffix.

```go
fmt.Printf("I have %s followers", humanize.Int(1589035)) // => I have 1.6 million followers
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.Int(999)      // => 999
humanize.Int(125000)   // => 125 thousand
humanize.Int(15600000) // => 15.6 million
```
</details>

#### Including symbol

Does the same as [Int](#int), but returns the symbol identifier, rather than the fully translated suffix. E.g. `million = M`

```go
fmt.Printf("I have $%s on my bank account", humanize.IntWithSymbol(785030)) // => I have $785K on my bank account
```

<details>
<summary>Additional examples of usage</summary>

```go
humanize.Int(999)      // => 999
humanize.Int(125000)   // => 125K
humanize.Int(15600000) // => 15.6M
```
</details>

## 🌍 Localization

`humanize` comes prepacked with localization support, which are easily switchable on the fly.

#### Switching locales

To use a different locale with `humanize`, simply import the locale you need, and register it.
To view a list of currently available locales, see [built-in locales](#built-in-locales).

```go
import (
	"github.com/imbue11235/humanize"
	"github.com/imbue11235/humanize/locale/da"
)

func main() {
	// register the locale
	humanize.RegisterLocale(da.Code, da.Locale)
	
	names := []string{"Hans", "Viggo", "Klaus"}
	
	fmt.Printf("Seen by %s", humanize.Slice(names, 2)) // => Seen by Hans, Viggo and one other
	
	// switch the locale
	humanize.SetLocale(da.Code)
	
	fmt.Printf("Set af %s", humanize.Slice(names, 2)) // => Set af Hans, Viggo og én anden
}
```

#### Registering custom locale

To register your own custom locale, simply specify a `locale.Map` mimicking the same key-value format
as the built-in locales. See [the english locale map](locale/en/locale.go) for additional details.

```go
humanize.RegisterLocale("my-locale", locale.Map{...})
```
#### Fallback locale

If the current selected locale does not support or define the translation needed, you can define a locale to fall back to. Make sure that the locale is registered beforehand or it will result in an error.

```go
humanize.SetFallbackLocale("my-fallback-locale")
```

#### Built-in locales

Currently, the following locales are included in the `humanize` package:

- [Danish](locale/da/locale.go)
- [English](locale/en/locale.go)
- [German](locale/de/locale.go)

#### Contributing locales

If you find your language is not on the list, and you want to add it, please [submit a PR](https://github.com/imbue11235/humanize/pulls).
It would be greatly appreciated and help the package become even more usable across languages.

List of wanted locales:

- [ ] Arabic
- [ ] Dutch
- [ ] Finnish
- [ ] French
- [ ] Italian
- [ ] Japanese
- [ ] Korean
- [ ] Mandarin Chinese
- [ ] Norwegian
- [ ] Russian
- [ ] Spanish
- [ ] Swedish
- [ ] Vietnamese
- [ ] Add your own

## 📜 License

This project is licensed under the [MIT license](LICENSE).
