package tparse

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var tomlSampleValid = `[Linus Torvalds]
Found = Linux, git

[Guido Van Rossum]
Found = Python, Gerrit

[Larry Wall]
Found = Perl`

var lTorvalds Entries = Entries{"Found": "Linux, git"}
var gvRossum Entries = Entries{"Found": "Python, Gerrit"}
var lWall Entries = Entries{"Found": "Perl"}

var tomlSampleInValid = `[Linus Torvalds]
Found = Linux, git
sd sdvlflmvdf vl dflv lkdfv flvklkmlk

[Guido Van Rossum]
Found = Python, Gerrit

[Larry Wall]
Found = Perl`

type testpairHeader struct {
	header, ans string
}

type testpairKeyVal struct {
	raw, key, val string
}

// Needed variables for getHeader funtion testing
var headerValid = []testpairHeader{
	{"[Heading]", "Heading"},
	{"[ Headi ng    ]", "Headi ng"},
	{"[       H[ad1] ng]", "H[ad1] ng"},
}
var headerInValid = []testpairHeader{
	{" [       H4ad1 ng ", ""},
	{"     H4ad1 ng ] ", ""},
	{"[       H4ad1 ng] 889 dfvdfv", ""},
	{"sdsvf [       H4ad1 ng]", ""},
	{"skcksdc [       H4ad1 ng] 889 dfvdfv", ""},
	{"[]", ""},
	{"cdcdcdcdscsdwv fdvfv", ""},
}

// Needed variables for getKeyValPair funtion testing
var keyValPairValid = []testpairKeyVal{
	{"Name = Mike", "Name", "Mike"},
	{" Name =    Mike Blitz ", "Name", "Mike Blitz"},
	{"Operations = a==1+9", "Operations", "a==1+9"},
	{" = Nothing", "", "Nothing"},
	{"[some = thing]", "[some", "thing]"},
}

var keyValPairInValid = []testpairKeyVal{
	{"Name is Mike", "", ""},
	{" Name is    Mike Blitz ", "", ""},
	{"s lamc  scdc", "", ""},
	{"[something]", "", ""},
}

func TestNewDict(t *testing.T) {
	d := NewDict()
	if d == nil {
		t.Error("Unexpected nil value for a NewDict")
	}
}

func TestgetKeyValPair(t *testing.T) {
	for _, val := range keyValPairValid {
		k, v := getKeyValPair(val.raw)
		if k != val.key || v != val.val {
			t.Error("Expected", val.key, val.val, ", got", k, v)
		}
	}

	for _, val := range keyValPairInValid {
		k, v := getKeyValPair(val.raw)
		if k != val.key || v != val.val {
			t.Error("Expected empty strings, got", k, v)
		}
	}
}

func TestgetHeader(t *testing.T) {
	for _, val := range headerValid {
		h := getHeader(val.header)
		if h != val.ans {
			t.Error("Expected", val.ans, ", got", h)
		}
	}

	for _, val := range headerValid {
		h := getHeader(val.header)
		if h != val.ans {
			t.Error("Expected", val.ans, ", got", h)
		}
	}
}

func TestDict_Parse(t *testing.T) {
	d := NewDict()
	err := d.Parse(tomlSampleValid)

	if err != nil {
		t.Error("Expected nil error but got : ", err)
	}

	d = NewDict()
	err = d.Parse(tomlSampleInValid)

	if err == nil {
		t.Error("Expected some error but got : nil")
	}
}

func TestDict_Find(t *testing.T) {
	d := NewDict()
	err := d.Parse(tomlSampleValid)

	if err != nil {
		t.Error("Expected nil error but got : ", err)
	}

	es, err := d.Find("Linus Torvalds")
	if err != nil {
		t.Error("Expected nil error but got : ", err)
	}

	if !reflect.DeepEqual(es, lTorvalds) {
		t.Error("Expected ", lTorvalds, ", got", es)
	}

	es, err = d.Find("Guido Van Rossum")
	if err != nil {
		t.Error("Expected nil error but got : ", err)
	}

	if !reflect.DeepEqual(es, gvRossum) {
		t.Error("Expected ", gvRossum, ", got", es)
	}

	es, err = d.Find("Larry Wall")
	if err != nil {
		t.Error("Expected nil error but got : ", err)
	}

	if !reflect.DeepEqual(es, lWall) {
		t.Error("Expected ", lWall, ", got", es)
	}

	es, err = d.Find("Rob Pike")
	if err == nil {
		t.Error("Expected some error but got : ", err)
	}
}

func TestEntries_Find(t *testing.T) {
	fs, err := lWall.Find("Found")
	if err != nil {
		t.Error("Expected some error but got : ", err)
	}

	if fs != "Perl" {
		t.Error("Expected \"Perl\" but got : ", fs)
	}
}

func BenchmarkDict_Parse(b *testing.B) {
	bs, _ := ioutil.ReadFile("samples/sample.toml")

	for i := 0; i < b.N; i++ {
		d := NewDict()
		d.Parse(string(bs))
	}
}
