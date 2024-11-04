package main

import (
	"dna-string-encoding/encoders"
	"fmt"
)

func main() {
	stringEncoder, err := encoders.NewStringDNAEncoder("ATCCTG")

	if err != nil {
		fmt.Println(err)
		return
	}

	encodedDna, err := stringEncoder.Encode()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encodedDna)
	decodedDna, err := stringEncoder.Decode(encodedDna)
	fmt.Println(decodedDna)

	bytesEncoder, err := encoders.NewBytesDNAEncoder("ATCCTG")
	bytesEncodedDna, err := bytesEncoder.Encode()
	bytesDecodedDna, err := bytesEncoder.Decode(bytesEncodedDna, len("ATCCTG"))
	fmt.Println(bytesDecodedDna)
}
