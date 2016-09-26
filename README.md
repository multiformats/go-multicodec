# go-multicodec - [multicodec](https://github.com/jbenet/multicodec) in Go

self-describing serialization. This is a Golang implementation of https://github.com/jbenet/multicodec

### Godoc: https://godoc.org/github.com/jbenet/multicodec

### Supported codecs

- `/protobuf`
- `/cbor`
- `/json`

## Usage

```go
import (
  "os"
  "io"

  cbor "github.com/multiformats/go-multicodec/cbor"
  json "github.com/multiformats/go-multicodec/json"
)

func main() {
  dec := cbor.Multicodec().NewDecoder(os.Stdin)
  enc := json.Multicodec().NewEncoder(os.Stdout)

  for {
    var item interface{}

    if err := dec.Decode(&item); err == io.EOF {
      break
    } else if err != nil {
      panic(err)
    }

    if err := enc.Encode(&item); err != nil {
      panic(err)
    }
  }
}
```
