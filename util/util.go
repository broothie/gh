package util

import (
	"encoding/gob"
	"io"
)

func DeepCopy[T any](v T) (result T) {
	r, w := io.Pipe()

	go func() {
		if err := gob.NewEncoder(w).Encode(v); err != nil {
			panic(err)
		}
	}()

	if err := gob.NewDecoder(r).Decode(&result); err != nil {
		panic(err)
	}

	return
}
