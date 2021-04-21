package dependencyInjection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, message string) {
	fmt.Fprintf(writer, "Hello, %s!", message)
}
