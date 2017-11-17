package libzcash

const (
	G1PrefixMask = 0x02
	G2PrefixMask = 0x0a
)

type Fq [32]byte
type Fq2 [64]byte

type CompressedG1 struct {
	yLsb bool
	x    Fq
}

type CompressedG2 struct {
	yGt bool
	x   Fq2
}

// TODO: serialization
type ZCProof struct {
	gA      CompressedG1
	gAPrime CompressedG1
	gB      CompressedG2
	gBPrime CompressedG1
	gC      CompressedG1
	gCPrime CompressedG1
	gK      CompressedG1
	gH      CompressedG1
}
