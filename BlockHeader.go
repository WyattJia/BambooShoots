// bitcoin block header

type BlockHeader struct {
    Version int32

	PrevBlock chainhash.Hash

	MerkleRoot chainhash.hash

	Timestamp time.time

	Bits uint32

	Nonce uint32
}


// bitcoin block
// 
// Field    Description    Size
// Magic no    value always 0xD9B4BEF9    4 bytes
// Blocksize    number of bytes following up to end of block    4 bytes
// Blockheader    consists of 6 items    80 bytes
// Transaction counter    positive integer VI = VarInt    1 - 9 bytes
// transactions    the (non empty) list of transactions    -many transactions
// 














