package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"unicode"
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprintln(m.Stderr, err)
		os.Exit(1)
	}
}

// Main represents the main program execution.
type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Run executes the program.
func (m *Main) Run(args ...string) error {
	// Parse command line flags.
	fs := flag.NewFlagSet("unistat", flag.ContinueOnError)
	fs.SetOutput(m.Stderr)
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Validate flags.
	if fs.NArg() > 1 {
		return errors.New("too many files specified")
	}

	// Determine if stdin or a CLI arg should be used as input.
	r, err := m.Reader(fs.Arg(0))
	if err != nil {
		return err
	}
	defer r.Close()

	// Read each rune and calculate stats.
	stats, err := m.Stat(bufio.NewReader(r))
	if err != nil {
		return err
	}

	// Print stats to terminal.
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Control:", stats.ControlN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Digit:", stats.DigitN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Graphic:", stats.GraphicN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Letter:", stats.LetterN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Lower:", stats.LowerN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Mark:", stats.MarkN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Number:", stats.NumberN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Print:", stats.PrintN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Punct:", stats.PunctN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Space:", stats.SpaceN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Symbol:", stats.SymbolN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Title:", stats.TitleN)
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Upper:", stats.UpperN)
	fmt.Fprintln(m.Stdout, "")
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Multibyte:", stats.MultiByteN)
	fmt.Fprintln(m.Stdout, "")
	fmt.Fprintf(m.Stdout, "%-10s %d\n", "Total:", stats.TotalN)
	fmt.Fprintln(m.Stdout, "")

	return nil
}

// Reader returns a file handle if filename is specified. Otherwise returns stdin.
// The reader must be closed by the caller.
func (m *Main) Reader(filename string) (io.ReadCloser, error) {
	// Use stdin if there's no file specified.
	if filename == "" {
		fmt.Fprintln(m.Stderr, "reading from stdin\n")
		return ioutil.NopCloser(m.Stdin), nil
	}

	// Otherwise open the file.
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Stat calculates statistics for all runes read from r.
func (m *Main) Stat(r io.RuneReader) (Stats, error) {
	var stats Stats

	for {
		// Read next character.
		ch, sz, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return stats, err
		}

		// Calculate stats.
		stats.TotalN++
		if unicode.IsControl(ch) {
			stats.ControlN++
		}
		if unicode.IsDigit(ch) {
			stats.DigitN++
		}
		if unicode.IsGraphic(ch) {
			stats.GraphicN++
		}
		if unicode.IsLetter(ch) {
			stats.LetterN++
		}
		if unicode.IsLower(ch) {
			stats.LowerN++
		}
		if unicode.IsMark(ch) {
			stats.MarkN++
		}
		if unicode.IsNumber(ch) {
			stats.NumberN++
		}
		if unicode.IsPrint(ch) {
			stats.PrintN++
		}
		if unicode.IsPunct(ch) {
			stats.PunctN++
		}
		if unicode.IsSpace(ch) {
			stats.SpaceN++
		}
		if unicode.IsSymbol(ch) {
			stats.SymbolN++
		}
		if unicode.IsTitle(ch) {
			stats.TitleN++
		}
		if unicode.IsUpper(ch) {
			stats.UpperN++
		}
		if sz > 1 {
			stats.MultiByteN++
		}
	}

	return stats, nil
}

// Stats represents unicode stats for a set of runes.
type Stats struct {
	TotalN     int
	ControlN   int
	DigitN     int
	GraphicN   int
	LetterN    int
	LowerN     int
	MarkN      int
	NumberN    int
	PrintN     int
	PunctN     int
	SpaceN     int
	SymbolN    int
	TitleN     int
	UpperN     int
	MultiByteN int
}
