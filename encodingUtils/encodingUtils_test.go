package encodingUtils_test

import (
	"github.com/olamiko/key-value-store/encodingUtils"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {

	data := encodingUtils.MyEncoding{uint8(10), uint8(1), [1]string{"foo=bar"}, uint8(01)}

	encoding, _ := encodingUtils.Encode(data)

	decodedData, err := encodingUtils.Decode(encoding)
	t.Fatalf("err %v", err)
	//if derefencedData != data {
	t.Fatalf("expected value %v, but got %v", data, decodedData)
	//}

}
