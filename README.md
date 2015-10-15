# Intro
Memstream is an expandable ReadWriteSeeker for Golang that works with an internally managed
byte buffer. Operation is usage is intended to be seamless and smooth.

In situations where the maximum read/write sizes are known, a fixed []byte/byte buffer
will likely offer better performance.

# Docs
As with all Go packages, documentation can be found on
[godoc.org](https://godoc.org/github.com/traherom/memstream).

# Install
     $ go get github.com/traherom/memstream
