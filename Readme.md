# tparse
tparse is a golang library for parsing simple toml like syntax

For Documentation visit : (Link)[https://godoc.org/github.com/shivam07a/tparse]

## Sample Usage
```go
tomlStr := `[Linus Torvalds]
Found = Linux, git

[Guido Van Rossum]
Found = Python, Gerrit

[Larry Wall]
Found = Perl`
var dict *tparse.Dict = tparse.NewDict()
dict.Parse(tomlStr)
e, err := dict.Find("Linus Torvalds")
if err != nil {
  fmt.Println(err)
  return
}
found, err := e.Find("Found")
if err != nil {
  fmt.Println(err)
  return
}
fmt.Println(found)

// Output:
// Linux, git
```
