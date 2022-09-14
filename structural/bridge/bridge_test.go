package bridge

import (
	"errors"
	"testing"
)

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content received but writer was empty")
	return
}
func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImp1{}

	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error trying to use API 1 implementation: Message: %s\n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	testWriter := TestWriter{}
	api2 := PrinterImp2{Writer: &testWriter}

	expectedMsg := "Hello"
	err := api2.PrintMessage(expectedMsg)
	if err != nil {
		t.Errorf("Error trying to use API2 implementation %s\n", err.Error())
	}
	if testWriter.Msg != expectedMsg {
		t.Fatalf("API2 did not write correcty on the io Write \n Actual:%s\n Expected : %s", testWriter.Msg, expectedMsg)
	}
}

func TestNormalPrinter(t *testing.T) {
	testWriter := TestWriter{}
	expectedMsg := "Hello io writer"
	normal := NormalPrinter{Msg: expectedMsg, Printer: &PrinterImp2{Writer: &testWriter}}

	err := normal.Print()

	if err != nil {
		t.Errorf(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("The expected message on the io.Writer doesn't match actual. \n Actual: %s\n Expected: %s \n", testWriter.Msg, expectedMsg)
	}
}

func TestPackPrinter(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMsg := "Message from Packt: Hello io.Writer"

	packt := PacktPrinter{Msg: passedMessage, Printer: &PrinterImp1{}}

	err := packt.Print()

	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{Msg: passedMessage, Printer: &PrinterImp2{Writer: &testWriter}}

	err = packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("The expected message on the io.Writer doesn't match actual \n Actual: %s\n Expected: %s \n", testWriter.Msg, expectedMsg)
	}
}
