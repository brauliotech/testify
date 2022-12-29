package printer

import "fmt"

type Computer struct {
	// Pr is a field of type Printer interface
	Pr Printer
}

// PrintA4 prints content in an A4 paper format
func (c *Computer) PrintA4(content string) error {
	c.Pr.Print("A4", content)
	return nil
}

// MyPrinter is a type of printer that could do printing jobs
type MyPrinter struct{}

// Print is a method of type *MyPrinter
// You may give the paperSize and content to it and it will print for you.
func (e *MyPrinter) Print(paperSize string, content string) error {
	fmt.Println("Printing:", content, ", with paper size:", paperSize)
	return nil
}

// Printer is an interface which implements the Print function.
type Printer interface {
	Print(paperSize string, content string) error
}
