package encodingUtils

import (
	"bytes"
	"encoding/binary"
	//"fmt"
)

// Encoding format
//
// start byte - a byte
// command byte - a byte
// size byte - 4 bytes
// data byte - as much as needed
// end byte - a byte

type MyEncoding struct {
	Start   uint8
	Command uint8
	Data    [10]byte
}

func Encode(toEncode MyEncoding) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, &toEncode)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(byteData []byte) (*MyEncoding, error) {
	var decodedData MyEncoding

	buf := bytes.NewReader(byteData)

	err := binary.Read(buf, binary.BigEndian, &decodedData)
	if err != nil {
		return nil, err
	}
	return &decodedData, nil

}
