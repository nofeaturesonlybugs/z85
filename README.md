[![Go Reference](https://pkg.go.dev/badge/github.com/nofeaturesonlybugs/z85.svg)](https://pkg.go.dev/github.com/nofeaturesonlybugs/z85)
[![Go Report Card](https://goreportcard.com/badge/github.com/nofeaturesonlybugs/z85)](https://goreportcard.com/report/github.com/nofeaturesonlybugs/z85)
[![Build Status](https://app.travis-ci.com/nofeaturesonlybugs/z85.svg?branch=master)](https://app.travis-ci.com/nofeaturesonlybugs/z85)
[![codecov](https://codecov.io/gh/nofeaturesonlybugs/z85/branch/master/graph/badge.svg)](https://codecov.io/gh/nofeaturesonlybugs/z85)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A golang implementation of ZeroMQ Z85 encoding as specified at https://rfc.zeromq.org/spec:32/Z85/

Credit

This implementation is influenced in part by https://github.com/tilinna/z85

Enhancements

This implementation does not require the caller to allocate storage buffers before calls to Encode or Decode.

This implementation provides functions to pad inputs to a length that is a multiple of 4; a complementary trim function is also provided.
