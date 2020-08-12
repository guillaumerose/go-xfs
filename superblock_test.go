package xfs

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

const block1 = "5846534200001000000000000079eefb0000000000000000000000000000000095e7129735d44a2db7735ed1dc283b050000000000080006000000000000008000000000000000810000000000000082000000010002fa40000000290000000000000a00b4f5020002000008726f6f7400000000000000000c09090312000019000000000001bee00000000000001c5500000000004b8e070000000000000000ffffffffffffffffffffffffffffffff0608000000000008000000000000000000000000000000010000018a0000018a00000000000000050000000300000000b38cba870000000400000000000022120000005700004a2e000000000000000000000000000000000000000000000000"
const block2 = "5846534200001000000000000008f00000000000000000000000000000000000910678fff77e4a7d8d5386f2ac47a82300000000000800060000000000000080000000000000008100000000000000820000000100023c00000000040000000000000a00b4b5020002000008726f6f7400000000000000000c090903120000190000000000006a00000000000000007200000000000286c20000000000000000ffffffffffffffffffffffffffffffff0000000000000008000000000000000000000000000000010000018a0000018a00000000000000050000000300000000fee0eb6f00000004ffffffffffffffff0000000500003f32000000000000000000000000000000000000000000000000"

func TestSuperblock1(t *testing.T) {
	bin, err := hex.DecodeString(block1)
	require.NoError(t, err)
	rs, err := NewSuperblockWithReader(bytes.NewReader(bin))
	assert.NoError(t, err)
	assert.Equal(t, rs.Data(), &SuperblockData{
		MagicNum:   Magic,
		BlockSize:  4096,
		DBlocks:    7991035,
		RBlocks:    0,
		RExtents:   0,
		UUID:       [16]uint8{0x95, 0xe7, 0x12, 0x97, 0x35, 0xd4, 0x4a, 0x2d, 0xb7, 0x73, 0x5e, 0xd1, 0xdc, 0x28, 0x3b, 0x05},
		LogStart:   0x80006,
		RootIno:    0x80,
		RBIno:      0x81,
		RSumIno:    0x82,
		RExtSize:   0x1,
		AgBlocks:   195136,
		AgCount:    41,
		RBmBlocks:  0,
		LogBlocks:  0xa00,
		VersionNum: 0xb4f5,
		SectSize:   512,
		InodeSize:  512,
		InoPBlock:  0x8,
		FName:      [48]uint8{0x72, 0x6f, 0x6f, 0x74, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc, 0x9, 0x9, 0x3, 0x12, 0x0, 0x0, 0x19, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xbe, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x55, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4b, 0x8e, 0x7, 0x0, 0x0, 0x0, 0x0},
		BlockLog:   0,
		SectLog:    0,
		InodeLog:   0,
		InoPBLog:   0,
		AgBlkLog:   255,
		RextSLog:   255,
		ImaxPct:    255,
		Icount:     0xffffffffffffffff,
		Ifree:      0xffffffffff060800,
		FDBlocks:   0x8000000,
		FRExtents:  0x0,
	})
}

func TestSuperblock2(t *testing.T) {
	bin, err := hex.DecodeString(block2)
	require.NoError(t, err)
	rs, err := NewSuperblockWithReader(bytes.NewReader(bin))
	require.NoError(t, err)
	fmt.Println(rs)
}
