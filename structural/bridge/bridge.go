package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImp1 struct{}

func (p *PrinterImp1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImp2 struct {
	Writer io.Writer
}

func (d *PrinterImp2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (np *NormalPrinter) Print() error {
	np.Printer.PrintMessage(np.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
