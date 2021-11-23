package encodingUtils_test

import (
	"github.com/olamiko/key-value-store/encodingUtils"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {

	sentPayload := [10]byte{}
	data := encodingUtils.MyEncoding{uint8(10), uint8(1), sentPayload}

	encoding, err := encodingUtils.Encode(data)

	if err != nil {
		t.Fatalf("encoding err %v", err)
	}

	decodedData, err := encodingUtils.Decode(encoding)
	if err != nil {
		t.Fatalf("decoding err %v", err)
	}

	if decodedData.Data != data.Data {
		t.Fatalf("expected value %v, but got %v", data, decodedData)
	}

}
