package encodingUtils

import (
	"bytes"
	"encoding/binary"
)

// Encoding format
//
// start byte - a byte
// command byte - a byte
// size byte - 4 bytes
// data byte - as much as needed
// end byte - a byte

type myEncoding struct {
	start   uint8
	command uint8
	size    uint32
	data    []string
	end     uint8
}

func encode(toEncode myEncoding) ([]bytes, err) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, toEncode)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(byteData []bytes) (myEncoding, error) {
	var decodedData myEncoding
	buf := bytes.NewReader(byteData)
	err := binary.Read(buf, binary.BigEndian, &decodedData)
	if err != nil {
		return nil, err
	}
	return decodedData, nil

}
