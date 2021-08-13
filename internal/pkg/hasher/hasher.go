package hasher

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
)

type hasher struct{}

func New() *hasher {
	return &hasher{}
}

func (m hasher) GetHashFromStruct(o interface{}) (hash string, err error) {
	hs := sha256.New()

	err = gob.NewEncoder(hs).Encode(o)
	if err != nil {
		return "", errors.Wrapf(err, "cannot get cache key for struct: %#v", o)
	}

	return fmt.Sprintf("%x", hs.Sum(nil)), nil
}
