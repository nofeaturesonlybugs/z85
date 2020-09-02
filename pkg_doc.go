// Package z85 provides ZeroMQ Z85 encoding.
//
// Dependencies
//
// The following packages act as dependencies:
//	github.com/nofeaturesonlybugs/errors
//
// Credit
//
// This implementation is influenced in part by https://github.com/tilinna/z85
//
// Enhancements
//
// This implementation does not require the caller to allocate storage buffers before calls to Encode or Decode.
//
// This implementation provides padding facilities for inputs whose lengths are not the correct size
// for Z85; i.e. this package can be used to Z85 encode arbitrary length inputs.
//
// Usage
//
// If your inputs are guaranteed to be multiples of the correct sizes, for example when dealing
// with cryptography keys, then use the Encode() and Decode() functions.
//	Encode() expects length is multiple of 4
//	Decode() expects length is multiple of 5
//
// If your inputs are not guaranteed to be multiples of the correct sizes, for example when using
// z85 as a substitute for base64, then use the PaddedEncode() and PaddedDecode() functions.
//
// About that Padding
//
// The padding implementation is specific to this package and not a universally implemented
// solution.
//
// Z85 Spec
//
// You can view the ZeroMQ Z85 spec @ https://rfc.zeromq.org/spec:32/Z85/
package z85
