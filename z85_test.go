package z85_test

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/nofeaturesonlybugs/z85"
	"testing"
)

func TestErrors(t *testing.T) {
	{
		invalid := []byte{0x01, 0x02, 0x03}
		_, err := z85.Encode(invalid)
		if err == nil {
			t.FailNow()
		}
	}

	{
		invalid := string("1234")
		_, err := z85.Decode(invalid)
		if err == nil {
			t.FailNow()
		}

		_, err = z85.PaddedDecode(invalid)
		if err == nil {
			t.FailNow()
		}
	}
}

func TestPadTrim(t *testing.T) {
	{
		exact := []byte{0x01, 0x02, 0x03, 0x04}
		_, err := z85.Encode(exact)
		if err != nil {
			t.FailNow()
		}
	}

	{
		returned := z85.Trim(nil)
		if returned != nil {
			t.FailNow()
		}

		a := []byte{}
		returned = z85.Trim(a)
		if fmt.Sprintf("%p", a) != fmt.Sprintf("%p", returned) {
			t.FailNow()
		}
	}

	{
		badtrim := []byte{0xab, 0xcd, 0x08, 0x09}
		returned := z85.Trim(badtrim)
		if bytes.Compare(returned, badtrim) != 0 {
			t.FailNow()
		}
	}
}

func TestPad(t *testing.T) {
	var a, b []byte

	a = []byte{0x00}
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x00, 0x03, 0x03, 0x03}) != 0 {
		t.FailNow()
	}

	a = []byte{0x00, 0x00}
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x00, 0x00, 0x02, 0x02}) != 0 {
		t.FailNow()
	}

	a = []byte{0x00, 0x00, 0x00}
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x00, 0x00, 0x00, 0x01}) != 0 {
		t.FailNow()
	}

	a = []byte{0x00, 0x00, 0x00, 0x00}
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x00, 0x00, 0x00, 0x00, 0x04, 0x04, 0x04, 0x04}) != 0 {
		t.FailNow()
	}

	a = []byte{}
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x04, 0x04, 0x04, 0x04}) != 0 {
		t.FailNow()
	}

	a = nil
	b = z85.Pad(a)
	if bytes.Compare(b, []byte{0x04, 0x04, 0x04, 0x04}) != 0 {
		t.FailNow()
	}
}

func TestTrim(t *testing.T) {
	var a, b []byte

	a = []byte{0x04, 0x04, 0x04, 0x04}
	b = z85.Trim(a)
	if fmt.Sprintf("%p", a) == fmt.Sprintf("%p", b) {
		t.FailNow()
	}
	if bytes.Compare(a, b) == 0 {
		t.FailNow()
	}
}

func ExampleDecode() {
	raw := "HelloWorld"
	decoded, err := z85.Decode(raw)
	if err != nil {
		fmt.Println(err)
	}
	parts := []string{}
	for _, b := range decoded {
		parts = append(parts, fmt.Sprintf("0x%02x", b))
	}
	fmt.Println("[ " + strings.Join(parts, ", ") + " ]")
	// Output: [ 0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5b ]
}
func ExampleDecode_error() {
	raw := "HelloWorld" + string([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	_, err := z85.Decode(raw)
	if err != nil {
		fmt.Println("Received error!")
	}
	// Output: Received error!
}
func ExampleEncode() {
	raw := []byte{0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5b}
	encoded, err := z85.Encode(raw)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encoded)
	// Output: HelloWorld
}
func Example_padding() {
	var err error
	var encoded string
	// raw's size is not a multiple of 4 so will cause an error.
	raw := []byte{0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5b, 0x01, 0x02, 0x03}
	_, err = z85.Encode(raw)
	if err != nil {
		fmt.Println("Got expected error!")
	}

	// But the Padded*() functions can handle it.
	encoded, err = z85.PaddedEncode(raw)
	if err != nil {
		fmt.Println(err)
	}

	decoded, err := z85.PaddedDecode(encoded)
	if err != nil {
		fmt.Println(err)
	}

	if bytes.Compare(decoded, raw) == 0 {
		fmt.Println("It worked!")
	}
	// Output: Got expected error!
	// It worked!
}
