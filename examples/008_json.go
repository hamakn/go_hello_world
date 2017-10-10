package examples

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Id   int
	Text string
}

func JSONExample() {
	message := Message{1, "JSONExample"}

	// encoding to json
	bdata, _ := json.Marshal(message)
	// output
	os.Stdout.Write(bdata)
	fmt.Println("")

	// Array
	messages := []Message{message}
	messages = append(messages, Message{2, "JSONArrayExample"})
	bdata2, _ := json.Marshal(messages)
	os.Stdout.Write(bdata2)
	fmt.Println("")
}
