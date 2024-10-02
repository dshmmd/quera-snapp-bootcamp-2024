package resolver

import "io"

type Resolver func(reader io.Reader) (string, error)
