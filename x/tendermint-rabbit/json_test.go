package listener

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/buger/jsonparser"
)

func Print(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	fmt.Println(string(value))
	return
}
func PrintKeyVal(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	// fmt.Printf("key: %s, value: %s\n", string(key), string(value))
	fmt.Println(strings.Split(string(key), "."))
	jsonparser.ArrayEach(value, Print)
	return nil
}

func TestJson(t *testing.T) {
	data, err := ioutil.ReadFile("/home/gleb/Downloads/tx_to_chain.json")
	if err != nil {
		t.Fatal(err)
	}
	result, _, _, err := jsonparser.Get(data, "result")
	if err != nil {
		t.Fatal(err)
	}
	events, _, _, err := jsonparser.Get(result, "events")
	if err != nil {
		t.Fatal(err)
	}
	jsonparser.ObjectEach(events, PrintKeyVal)
}
