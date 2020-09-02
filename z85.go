package z85

import (
	"bytes"
	"encoding/binary"
	"github.com/nofeaturesonlybugs/errors"
)

var (
	encoder = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#"
	decoder = []byte{
		0x00, 0x44, 0x00, 0x54, 0x53, 0x52, 0x48, 0x00,
		0x4B, 0x4C, 0x46, 0x41, 0x00, 0x3F, 0x3E, 0x45,
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x40, 0x00, 0x49, 0x42, 0x4A, 0x47,
		0x51, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A,
		0x2B, 0x2C, 0x2D, 0x2E, 0x2F, 0x30, 0x31, 0x32,
		0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A,
		0x3B, 0x3C, 0x3D, 0x4D, 0x00, 0x4E, 0x43, 0x00,
		0x00, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20,
		0x21, 0x22, 0x23, 0x4F, 0x00, 0x50, 0x00, 0x00}
	sizeDecoder = len(decoder)
)

// Encode encodes slice to a Z85 encoded string; length of slice must be multiple
// of 4.
func Encode(slice []byte) (string, error) {
	size := len(slice)
	if size%4 != 0 {
		return "", errors.Errorf("Slice length must be multiple of 4")
	}
	rv := make([]byte, size*5/4)
	dest := rv
	for chunk, chunks := 0, size/4; chunk < chunks; chunk++ {
		value := binary.BigEndian.Uint32(slice[0:4])
		// Generate 5 characters
		for char := 4; char >= 0; char-- {
			dest[char] = encoder[value%85]
			value = value / 85
		}
		dest = dest[5:]
		slice = slice[4:]
	}
	return string(rv), nil
}

// Decode decodes Z85 string to a slice; length of string must be multiple of 5.
func Decode(str string) ([]byte, error) {
	size := len(str)
	if size%5 != 0 {
		return nil, errors.Errorf("String length must be multiple of 5")
	}
	rv := make([]byte, size*4/5)
	dest := rv
	for chunk, chunks := 0, size/5; chunk < chunks; chunk++ {
		value := uint32(0)
		for char := 0; char < 5; char++ {
			index := str[char] - 32
			if index < 0 || int(index) >= sizeDecoder {
				return nil, errors.Errorf("Invalid Z85 string @ input( 0x%02x )", str[char])
			}
			value = value*85 + uint32(decoder[str[char]-32])
		}
		binary.BigEndian.PutUint32(dest, value)
		dest = dest[4:]
		str = str[5:]
	}
	return []byte(rv), nil
}

// Pad pads the incoming slice to be a length that is multiple of 4.
func Pad(slice []byte) []byte {
	size := len(slice)
	need := 4 - (size % 4)
	padding := bytes.Repeat([]byte{uint8(need)}, need)
	padded := append(slice, padding...)
	return padded
}

// Trim removes the padding added by Pad().
func Trim(slice []byte) []byte {
	size := len(slice)
	if size == 0 {
		return slice
	}
	trim := int(uint8(slice[size-1]))
	if trim > 4 || trim > size {
		return slice
	}
	rv := make([]byte, size-trim)
	copy(rv, slice[0:size-trim])
	return rv

}

// PaddedDecode decodes a string that is assumed to have been padded by this package or another
// implementation that uses the same padding scheme.
func PaddedDecode(str string) ([]byte, error) {
	decoded, err := Decode(str)
	if err != nil {
		return nil, errors.Go(err)
	}
	return Trim(decoded), nil
}

// PaddedEncode encodes a slice of arbitrary length.
func PaddedEncode(slice []byte) (string, error) {
	slice = Pad(slice)
	return Encode(slice)
}
