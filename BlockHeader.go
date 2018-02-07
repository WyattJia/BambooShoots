// bitcoin block header

type BlockHeader struct {
    Version int32

	PrevBlock chainhash.Hash

	MerkleRoot chainhash.hash

	Timestamp time.time

	Bits uint32

	Nonce uint32
}
