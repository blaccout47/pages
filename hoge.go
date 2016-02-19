package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Status int

const (
	StatusStart Status = iota
	StatusProcessing
	StatusEnd
)

func (s Status) String() string {
	switch s {
	case StatusStart:
		return "start"
	case StatusProcessing:
		return "processing"
	case StatusEnd:
		return "end"
	}
	return ""
}

func main() {
	fmt.Println("TypeOf: " + reflect.TypeOf(StatusStart).Name())
	fmt.Printf("String: %s\n", StatusStart)
	fmt.Printf("Int: %d\n", StatusStart)

	fmt.Print("Println: ")
	fmt.Println(StatusStart)

	fmt.Println("=== interface ===")
	InterfaceTest(StatusEnd)
}

func InterfaceTest(status Status) {
	data := map[string]interface{}{
		"status": status,
	}
	fmt.Print("interface: ")
	fmt.Println(data)

	fmt.Println("test1")
	jsonData, _ := json.Marshal(data)
	fmt.Print("json: ")
	fmt.Println(string(jsonData))

	data = map[string]interface{}{
		"status": status.String(),
	}

	fmt.Println("test2")
	jsonData, _ = json.Marshal(data)
	fmt.Print("json: ")
	fmt.Println(string(jsonData))
}
