# Gophlings

Gophlings is a set of small Go programs with something wrong in each one.
Fix the program, run the tests, and move on. The path follows
[A Tour of Go](https://go.dev/tour/): packages, types, control flow, and concurrency.

This is the Go companion to [ziglings](https://codeberg.org/ziglings/exercises).

## Run

Install [Go](https://go.dev/dl/) 1.22 or newer, then from this directory:

```
go run .
```

The runner stops at the first exercise that fails. Open that directory, read
the comments, fix the code, and run `go run .` again.

```
go run . -list
go run . -n 4
go run . -from 11
go run . -watch
```

`-watch` reruns the current exercise when you save a file in it.
`-n` and `-from` take the number at the start of the directory name.

To test one exercise directly:

```
go test ./04_variables/...
```

`01_functions` is a program. Run it with `go run ./01_functions`.

The exercises live in one module. Each directory is a separate package, so
run `go test` on that directory rather than from inside a nested module.

Finished programs are in `solutions/`. Try the exercise before you open the
matching file.

## Exercises

- [x] Packages (`00_packages`)
- [x] Functions (`01_functions`)
  - [x] Public
  - [x] Private
  - [x] Parameters
  - [x] Returns (single and multiple)
- [x] Strings (`02_strings`)
  - [x] UTF-8
  - [x] Hello World
- [x] Rune (`03_runes`)
  - [x] Single character
- [x] Variables (`04_variables`)
  - [x] Initializers
  - [x] Short declarations (`:=`)
- [x] Types (`05_types`)
- [x] Zero values (`06_zero`)
- [x] Type conversion (`07_conversion`)
- [x] Type inference (`08_inference`)
- [x] Constant (`09_constants`)
- [x] Numeric constant (`10_numeric`)
- [x] For (`11_for`)
  - [x] Init
  - [x] Condition
  - [x] Post statement
  - [x] Forever
  - [x] Range
- [x] Conditional (`12_conditional`)
  - [x] Short if
  - [x] Else
  - [x] Switch
- [x] Defer (`13_defer`)
  - [x] Stacked
- [x] Pointers (`14_pointers`)
- [x] Struct (`15_struct`)
  - [x] Fields
  - [x] Pointers
  - [x] Receiver
  - [x] Pointer receiver
- [x] Array (`16_array`)
- [x] Slice (`17_slice`)
  - [x] Defaults
  - [x] Nil
  - [x] Make
  - [x] Multidimensional
  - [x] Appending
- [x] Maps (`18_maps`)
  - [x] Mutating
- [x] Functions (`19_closures`)
  - [x] HTTP function handlers
  - [x] Closures
- [x] Interface (`20_interface`)
  - [x] Type alias
  - [x] Method signatures
  - [x] Implement implicitly
  - [x] Working with nils
  - [x] Empty (`interface{}` or `any`)
- [x] Type assertion (`21_assertion`)
- [x] Goroutine (`22_goroutine`)
- [x] Channels (`23_channels`)
  - [x] Buffer
- [x] Select (`24_select`)
- [x] Mutex (`25_mutex`)
