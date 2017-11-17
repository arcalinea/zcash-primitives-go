package libzcash

const (
	ZcNumJsInputs           = 2
	ZcNumJsOutputs          = 2
	ZcNotePlaintextLeading  = 1
	ZcVSize                 = 8
	ZcRhoSize               = 32
	ZcRSize                 = 32
	ZcMemoSize              = 512
	ZcNotePlaintextSize     = ZcNotePlaintextLeading + ZcVSize + ZcRhoSize + ZcRSize + ZcMemoSize
	noteEncryptionAuthBytes = 16
	ZcNoteCiphertextSize    = ZcNotePlaintextSize + noteEncryptionAuthBytes
)

type JSDescription struct {
	vPubOld       int64
	vPubNew       int64
	anchor        [32]byte
	nullifiers    [ZcNumJsInputs][32]byte
	commitments   [ZcNumJsOutputs][32]byte
	oneTimePubKey [32]byte
	randomSeed    [32]byte
	macs          [ZcNumJsInputs][32]byte
	proof         ZCProof
	cipertexts    [ZcNumJsOutputs][ZcNoteCiphertextSize]byte
}
