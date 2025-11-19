package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

type UUID [16]byte

func New() (UUID, error) {
	var u UUID
	_, err := rand.Read(u[:])
	if err != nil {
		return u, err
	}
	// The 4-bit version field as defined by Section 4.2, set to 0b0100 (4). Occupies bits 48 through 51 of octet 6.
	u[6] = (u[6] & 0x0f) | 0x40

	// The 2-bit variant field as defined by Section 4.1, set to 0b10. Occupies bits 64 and 65 of octet 8.
	u[8] = (u[8] & 0x3f) | 0x80

	return u, nil
}

func (u UUID) Bytes() []byte {
	return u[:]
}

const uuidStrLen = 36

func (u UUID) String() string {

	var retBuf [uuidStrLen]byte
	// octet is 2 hex nums, so for each of the hexOctet nums below, we want to allocate 2 spots in our retBuf
	// 		  4hexOctet "-"  =  0-7, 8th is -
	hex.Encode(retBuf[0:8], u[0:4])
	retBuf[8] = '-'
	//        2hexOctet "-"   9-12, 13th is -
	hex.Encode(retBuf[9:13], u[4:6])
	retBuf[13] = '-'
	//        2hexOctet "-"   14-17, 18th is -
	hex.Encode(retBuf[14:18], u[6:8])
	retBuf[18] = '-'
	//        2hexOctet "-"  19-22, 23rd is -
	hex.Encode(retBuf[19:23], u[8:10])
	retBuf[23] = '-'
	//        6hexOctet
	hex.Encode(retBuf[24:], u[10:16])

	return string(retBuf[:])
}
