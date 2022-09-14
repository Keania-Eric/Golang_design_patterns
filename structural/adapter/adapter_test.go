package adapter

import "testing"

func TestAdapter(t *testing.T) {

	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &LegacyPrinterImp{}, Msg: msg}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("message did not match %s \n", returnedMsg)
	}

	// checking for nil
	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()

	if returnedMsg != "Hello World!" {
		t.Errorf("message didn't match %s \n", returnedMsg)
	}
}
