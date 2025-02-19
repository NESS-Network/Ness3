// Code generated by github.com/MDLlife/skyencoder. DO NOT EDIT.
package blockdb

import (
	"errors"
	"math"

	"github.com/MDLlife/MDL/src/cipher/encoder"
	"github.com/MDLlife/MDL/src/coin"
)

// encodeSizeHashPairsWrapper computes the size of an encoded object of type hashPairsWrapper
func encodeSizeHashPairsWrapper(obj *hashPairsWrapper) uint64 {
	i0 := uint64(0)

	// obj.HashPairs
	i0 += 4
	{
		i1 := uint64(0)

		// x1.Hash
		i1 += 32

		// x1.PrevHash
		i1 += 32

		i0 += uint64(len(obj.HashPairs)) * i1
	}

	return i0
}

// encodeHashPairsWrapper encodes an object of type hashPairsWrapper to a buffer allocated to the exact size
// required to encode the object.
func encodeHashPairsWrapper(obj *hashPairsWrapper) ([]byte, error) {
	n := encodeSizeHashPairsWrapper(obj)
	buf := make([]byte, n)

	if err := encodeHashPairsWrapperToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeHashPairsWrapperToBuffer encodes an object of type hashPairsWrapper to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeHashPairsWrapperToBuffer(buf []byte, obj *hashPairsWrapper) error {
	if uint64(len(buf)) < encodeSizeHashPairsWrapper(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.HashPairs length check
	if uint64(len(obj.HashPairs)) > math.MaxUint32 {
		return errors.New("obj.HashPairs length exceeds math.MaxUint32")
	}

	// obj.HashPairs length
	e.Uint32(uint32(len(obj.HashPairs)))

	// obj.HashPairs
	for _, x := range obj.HashPairs {

		// x.Hash
		e.CopyBytes(x.Hash[:])

		// x.PrevHash
		e.CopyBytes(x.PrevHash[:])

	}

	return nil
}

// decodeHashPairsWrapper decodes an object of type hashPairsWrapper from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeHashPairsWrapper(buf []byte, obj *hashPairsWrapper) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.HashPairs

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.HashPairs = make([]coin.HashPair, length)

			for z1 := range obj.HashPairs {
				{
					// obj.HashPairs[z1].Hash
					if len(d.Buffer) < len(obj.HashPairs[z1].Hash) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.HashPairs[z1].Hash[:], d.Buffer[:len(obj.HashPairs[z1].Hash)])
					d.Buffer = d.Buffer[len(obj.HashPairs[z1].Hash):]
				}

				{
					// obj.HashPairs[z1].PrevHash
					if len(d.Buffer) < len(obj.HashPairs[z1].PrevHash) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.HashPairs[z1].PrevHash[:], d.Buffer[:len(obj.HashPairs[z1].PrevHash)])
					d.Buffer = d.Buffer[len(obj.HashPairs[z1].PrevHash):]
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeHashPairsWrapperExact decodes an object of type hashPairsWrapper from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeHashPairsWrapperExact(buf []byte, obj *hashPairsWrapper) error {
	if n, err := decodeHashPairsWrapper(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
