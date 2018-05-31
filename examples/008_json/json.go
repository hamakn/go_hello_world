package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Id   int
	Text string
}

func main() {
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

	// Load
	msgs1 := []Message{}
	json.Unmarshal([]byte(`[{"id":1,"text":"hoge"},{"id":2,"text":"fuga"}]`), &msgs1)
	fmt.Println(msgs1)

	// Load from bytes(io.Writer)
	msgs2 := []Message{}
	buf := bytes.NewBufferString(`[{"id":1,"text":"hoge"},{"id":2,"text":"fuga"}]`)
	json.NewDecoder(buf).Decode(&msgs2)
	fmt.Println(msgs2)
}
