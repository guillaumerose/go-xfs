package xfs

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	Magic uint32 = 0x58465342
)

type SuperblockData struct {
	// https://elixir.bootlin.com/linux/latest/source/fs/xfs/libxfs/xfs_format.h#L96
	MagicNum   uint32
	BlockSize  uint32
	DBlocks    uint64
	RBlocks    uint64
	RExtents   uint64
	UUID       [16]uint8
	LogStart   uint64
	RootIno    uint64
	RBIno      uint64
	RSumIno    uint64
	RExtSize   uint32
	AgBlocks   uint32
	AgCount    uint32
	RBmBlocks  uint32
	LogBlocks  uint32
	VersionNum uint16
	SectSize   uint16
	InodeSize  uint16
	InoPBlock  uint16
	FName      [48]uint8
	BlockLog   uint8
	SectLog    uint8
	InodeLog   uint8
	InoPBLog   uint8
	AgBlkLog   uint8
	RextSLog   uint8
	ImaxPct    uint8
	Icount     uint64
	Ifree      uint64
	FDBlocks   uint64
	FRExtents  uint64
}

type Superblock struct {
	data *SuperblockData
}

func (sb *Superblock) Data() *SuperblockData {
	return sb.data
}

func NewSuperblockWithReader(rs io.ReadSeeker) (sb *Superblock, err error) {
	sbd := new(SuperblockData)

	err = binary.Read(rs, binary.BigEndian, sbd)
	if err != nil {
		return nil, err
	}

	if sbd.MagicNum != Magic {
		return nil, fmt.Errorf("magic number not correct, got %x", sbd.MagicNum)
	}

	return &Superblock{
		data: sbd,
	}, nil
}
