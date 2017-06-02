package tparse_test

import (
	"fmt"

	"github.com/shivam07a/tparse"
)

func ExampleDict_Parse() {
	tomlStr := `[Linus Torvalds]
  Found = Linux, git

  [Guido Van Rossum]
  Found = Python, Gerrit

  [Larry Wall]
  Found = Perl`
	var dict *tparse.Dict = tparse.NewDict()
	err := dict.Parse(tomlStr)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	//
}

func ExampleDict_Find() {
	tomlStr := `[Linus Torvalds]
  Found = Linux, git

  [Guido Van Rossum]
  Found = Python, Gerrit

  [Larry Wall]
  Found = Perl`
	var dict *tparse.Dict = tparse.NewDict()
	err := dict.Parse(tomlStr)
	if err != nil {
		fmt.Println(err)
	}
	e, err := dict.Find("Linus Torvalds")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)
	// Output:
	// map[Found:Linux, git]
}

func ExampleEntries_Find() {
	tomlStr := `[Linus Torvalds]
  Found = Linux, git

  [Guido Van Rossum]
  Found = Python, Gerrit

  [Larry Wall]
  Found = Perl`
	var dict *tparse.Dict = tparse.NewDict()
	err := dict.Parse(tomlStr)
	if err != nil {
		fmt.Println(err)
	}
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
}
