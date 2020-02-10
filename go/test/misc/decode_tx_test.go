package misc

import (
	"encoding/base64"
	"testing"

	"github.com/cosmos/amino-js/go/src"
)

func TestDecodeTx(t *testing.T) {
	str := "1gHwYl3uCkOoo2GaChS536x615s3L5HNHZqLKYPpCK3tiRIUF5FHHqSe5EQ+wA9u2QAy8QEasjAaEQoFdWF0b20SCDExNjU3OTk1EhMKDQoFdWF0b20SBDUwMDAQwJoMGmoKJuta6YchAtQaCqFnshaZQp6rIkvAPyzThvCvXSDO+9AzbxVErqJPEkDWdRwgfQItPT+dDSYFMOtPqQwbbQ1j8+wfs/aBzhulg0YsRiMGZ1Z69dQmi5IT/0D/rRAb1xh6rJN7mQUN4g/FIgoxMTIyNjcyNzU0"
	original, err := base64.StdEncoding.DecodeString(str)

	originalDecoded := src.DecodeTx(original, false)

	json := string(originalDecoded)

	encoded := src.EncodeTx(originalDecoded, false)

	encodedDecoded := src.DecodeTx(encoded, false)

	json = string(originalDecoded)

	_ = err
	_ = encoded
	_ = json
	_ = encodedDecoded
	return
}

func TestDecodePosmintTx(t *testing.T) {
	//src.EncodePosmintTx()
}
