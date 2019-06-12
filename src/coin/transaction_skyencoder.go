// Code generated by github.com/MDLlife/skyencoder. DO NOT EDIT.
package coin

import (
	"errors"
	"math"

	"github.com/MDLlife/MDL/src/cipher"
	"github.com/MDLlife/MDL/src/cipher/encoder"
)

// encodeSizeTransaction computes the size of an encoded object of type Transaction
func encodeSizeTransaction(obj *Transaction) uint64 {
	i0 := uint64(0)

	// obj.Length
	i0 += 4

	// obj.Type
	i0++

	// obj.InnerHash
	i0 += 32

	// obj.Sigs
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 65

		i0 += uint64(len(obj.Sigs)) * i1
	}

	// obj.In
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 32

		i0 += uint64(len(obj.In)) * i1
	}

	// obj.Out
	i0 += 4
	{
		i1 := uint64(0)

		// x.Address.Version
		i1++

		// x.Address.Key
		i1 += 20

		// x.Coins
		i1 += 8

		// x.Hours
		i1 += 8

		i0 += uint64(len(obj.Out)) * i1
	}

	return i0
}

// encodeTransaction encodes an object of type Transaction to a buffer allocated to the exact size
// required to encode the object.
func encodeTransaction(obj *Transaction) ([]byte, error) {
	n := encodeSizeTransaction(obj)
	buf := make([]byte, n)

	if err := encodeTransactionToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeTransactionToBuffer encodes an object of type Transaction to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeTransactionToBuffer(buf []byte, obj *Transaction) error {
	if uint64(len(buf)) < encodeSizeTransaction(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Length
	e.Uint32(obj.Length)

	// obj.Type
	e.Uint8(obj.Type)

	// obj.InnerHash
	e.CopyBytes(obj.InnerHash[:])

	// obj.Sigs maxlen check
	if len(obj.Sigs) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Sigs length check
	if uint64(len(obj.Sigs)) > math.MaxUint32 {
		return errors.New("obj.Sigs length exceeds math.MaxUint32")
	}

	// obj.Sigs length
	e.Uint32(uint32(len(obj.Sigs)))

	// obj.Sigs
	for _, x := range obj.Sigs {

		// x
		e.CopyBytes(x[:])

	}

	// obj.In maxlen check
	if len(obj.In) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.In length check
	if uint64(len(obj.In)) > math.MaxUint32 {
		return errors.New("obj.In length exceeds math.MaxUint32")
	}

	// obj.In length
	e.Uint32(uint32(len(obj.In)))

	// obj.In
	for _, x := range obj.In {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Out maxlen check
	if len(obj.Out) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Out length check
	if uint64(len(obj.Out)) > math.MaxUint32 {
		return errors.New("obj.Out length exceeds math.MaxUint32")
	}

	// obj.Out length
	e.Uint32(uint32(len(obj.Out)))

	// obj.Out
	for _, x := range obj.Out {

		// x.Address.Version
		e.Uint8(x.Address.Version)

		// x.Address.Key
		e.CopyBytes(x.Address.Key[:])

		// x.Coins
		e.Uint64(x.Coins)

		// x.Hours
		e.Uint64(x.Hours)

	}

	return nil
}

// decodeTransaction decodes an object of type Transaction from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeTransaction(buf []byte, obj *Transaction) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Length
		i, err := d.Uint32()
		if err != nil {
			return 0, err
		}
		obj.Length = i
	}

	{
		// obj.Type
		i, err := d.Uint8()
		if err != nil {
			return 0, err
		}
		obj.Type = i
	}

	{
		// obj.InnerHash
		if len(d.Buffer) < len(obj.InnerHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.InnerHash[:], d.Buffer[:len(obj.InnerHash)])
		d.Buffer = d.Buffer[len(obj.InnerHash):]
	}

	{
		// obj.Sigs

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Sigs = make([]cipher.Sig, length)

			for z1 := range obj.Sigs {
				{
					// obj.Sigs[z1]
					if len(d.Buffer) < len(obj.Sigs[z1]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Sigs[z1][:], d.Buffer[:len(obj.Sigs[z1])])
					d.Buffer = d.Buffer[len(obj.Sigs[z1]):]
				}

			}
		}
	}

	{
		// obj.In

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.In = make([]cipher.SHA256, length)

			for z1 := range obj.In {
				{
					// obj.In[z1]
					if len(d.Buffer) < len(obj.In[z1]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.In[z1][:], d.Buffer[:len(obj.In[z1])])
					d.Buffer = d.Buffer[len(obj.In[z1]):]
				}

			}
		}
	}

	{
		// obj.Out

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Out = make([]TransactionOutput, length)

			for z1 := range obj.Out {
				{
					// obj.Out[z1].Address.Version
					i, err := d.Uint8()
					if err != nil {
						return 0, err
					}
					obj.Out[z1].Address.Version = i
				}

				{
					// obj.Out[z1].Address.Key
					if len(d.Buffer) < len(obj.Out[z1].Address.Key) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Out[z1].Address.Key[:], d.Buffer[:len(obj.Out[z1].Address.Key)])
					d.Buffer = d.Buffer[len(obj.Out[z1].Address.Key):]
				}

				{
					// obj.Out[z1].Coins
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Out[z1].Coins = i
				}

				{
					// obj.Out[z1].Hours
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Out[z1].Hours = i
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeTransactionExact decodes an object of type Transaction from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeTransactionExact(buf []byte, obj *Transaction) error {
	if n, err := decodeTransaction(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
