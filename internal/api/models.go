package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

//Int64 is int64
type Int64 int64

//MarshalInt64 is for marshaling the Int64
func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(i, 10))
	})
}

//UnmarshalInt64 is for marshaling the Int64
func UnmarshalInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
