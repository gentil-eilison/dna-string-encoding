package encoders

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

type DNAEncoderDecoder interface {
	Encode() (any, error)
	Decode() string
}

type DNAEncoder struct {
	dna string
}

type StringDNAEncoder DNAEncoder

func NewStringDNAEncoder(dna string) (*StringDNAEncoder, error) {
	if dna == "" {
		return nil, errors.New("you must pass a dna string")
	}
	return &StringDNAEncoder{
		dna: dna,
	}, nil
}

func (dnaEncoder *StringDNAEncoder) Encode() (string, error) {
	replacer := strings.NewReplacer("A", "00", "C", "01", "G", "10", "T", "11")
	encodedDna := replacer.Replace(dnaEncoder.dna)
	return encodedDna, nil
}

var dnaNucleotideStringMap = map[string]string{
	"00": "A",
	"01": "C",
	"10": "G",
	"11": "T",
}

func (dnaEncoder *StringDNAEncoder) Decode(encodedDna string) (string, error) {
	regex, err := regexp.Compile(`\d{2}`)
	if err != nil {
		return "", err
	}
	stringBits := regex.FindAll([]byte(encodedDna), -1)
	if stringBits == nil {
		return "", errors.New("no bits were found in the dna string")
	}

	var decodedDna string
	for _, nucleotide := range stringBits {
		decodedDna += dnaNucleotideStringMap[string(nucleotide)]
	}
	return decodedDna, nil
}

type BytesDNAEncoder DNAEncoder

func NewBytesDNAEncoder(dna string) (*BytesDNAEncoder, error) {
	if dna == "" {
		return nil, errors.New("you must pass a dna string")
	}
	return &BytesDNAEncoder{
		dna: dna,
	}, nil
}

var dnaBitMapping = map[rune]int{
	'A': 0b00,
	'C': 0b01,
	'G': 0b10,
	'T': 0b11,
}

func (dnaEncoder *BytesDNAEncoder) Encode() (int, error) {
	encodedDna := 0b00

	for _, nucleotide := range dnaEncoder.dna {
		dnaBit := dnaBitMapping[nucleotide]
		encodedDna = encodedDna << 2
		encodedDna = encodedDna | dnaBit
	}

	return encodedDna, nil
}

var dnaStringBitMapping = map[int]string{
	0b00: "A",
	0b01: "C",
	0b10: "G",
	0b11: "T",
}

func (dnaEncoder *BytesDNAEncoder) Decode(encodedDna, dnaLength int) (string, error) {
	decodedDna := make([]string, dnaLength)
	for i := 0; i < dnaLength; i++ {
		lastTwoBits := encodedDna & 0b11
		decodedDna[i] = dnaStringBitMapping[lastTwoBits]
		encodedDna = encodedDna >> 2
	}
	slices.Reverse(decodedDna)
	return strings.Join(decodedDna, ""), nil
}
