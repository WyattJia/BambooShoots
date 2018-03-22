// bitcoin block header

type BlockHeader struct {
	// version of the block
    Version int32

	// hash of previous block
	PrevBlock chainhash.Hash

    // Merkle tree reference to hash of all transactions for the block.
	MerkleRoot chainhash.hash

	// time of block was created.
	Timestamp time.time

	// difficulty target for the block.
	Bits uint32

	// Nonce used to generate the block.
	Nonce uint32
}


// bitcoin block
// 
// Field          Description                             Size
// Magic no       value always 0xD9B4BEF9                 4 bytes
// Blocksize      number of bytes following up to end of block    4 bytes
// Blockheader    consists of 6 items                     80 bytes
// Transaction counter    positive integer VI = VarInt    1 - 9 bytes
// transactions   the (non empty) list of transactions    -many transactions
// 

