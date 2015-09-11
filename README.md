unistat
=======

Unistat is a tool for calculating simple statistics on a set of unicode characters.


## Getting Started

To install, simply run:

```sh
$ go get github.com/benbjohnson/unistat
```

Then run `unistat` against your file:

```sh
$ unistat /path/to/myfile.txt
Control:   0
Digit:     27534
Graphic:   106517
Letter:    63330
Lower:     58824
Mark:      0
Number:    27534
Print:     106517
Punct:     10134
Space:     2069
Symbol:    3450
Title:     0
Upper:     4506

Total:     106517
```

You can also pipe in data from STDIN:

```sh
$ cat /path/to/myfile.txt | unistat
```