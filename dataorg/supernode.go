package dataorg

import (
	"github.com/erician/gpdDB/common/gpdconst"

	"github.com/erician/gpdDB/utils/conv"
)

//SuperNodeHeader like the NodeHeader, the struct of SuperNodeHeader
//is just to show the format of a block, not used
//And it occpies the first block, block num is 0
type SuperNodeHeader struct {
	blkID       int64 //use the block number as the blkID
	rootNodeID  int64 //it is the root node's blkID
	allPairsNum int64 //the number of key-values
	nextBlkNum  int64 //the next block number to be allocated starting from 1
	Len         int16 //the lenght of the SuperNode
	version     int8  //the gpdDB version used for parsing data
}

//SuperNode offset of all field
const (
	SuperNodeOffBlkID       int64 = 0
	SuperNodeOffRootNodeID  int64 = 8
	SuperNodeOffAllPairsNum int64 = 16
	SuperNodeOffNextBlkNum  int64 = 24
	SuperNodeOffLen         int64 = 32
	SuperNodeOffVersion     int64 = 34
)

//SuperNode all fields size
const (
	SuperNodeFieldSizeBlkID       int64 = 8
	SuperNodeFieldSizeRootNodeID  int64 = 8
	SuperNodeFieldSizeAllPairsNum int64 = 8
	SuperNodeFieldSizeNextBlkNum  int64 = 8
	SuperNodeFieldSizeLen         int64 = 2
	SuperNodeFieldSizeVersion     int64 = 1
)

//some field value of SuperNode
const (
	SuperNodeConstValueBlkID   int64 = 0
	SuperNodeConstValueLen     int16 = 0x23
	SuperNodeConstValueVersion int8  = 0x01

	SuperNodeInitValueAllPairsNum int64 = 0                            //initvalue means the value will be changed
	SuperNodeInitValueNextBlkNum  int64 = gpdconst.RootNodeInitBlockID //same with gpdconst.RootNodeInitBlockID
	SuperNodeInitValueRootNodeID  int64 = gpdconst.NotAllocatedBlockID
)

//const values of SuperNode
const (
	SuperNodeSize int64 = 4 * gpdconst.KB
)

//SuperNodeInit init the SuperNode
func SuperNodeInit(superNode []byte) {
	SuperNodeSetBlkID(superNode, SuperNodeConstValueBlkID)
	SuperNodeSetRootNodeID(superNode, SuperNodeInitValueRootNodeID)
	SuperNodeSetAllPairsNum(superNode, SuperNodeInitValueAllPairsNum)
	SuperNodeSetNextBlkNum(superNode, SuperNodeInitValueNextBlkNum)
	SuperNodeSetLen(superNode, SuperNodeConstValueLen)
	SuperNodeSetVersion(superNode, SuperNodeConstValueVersion)
}

//SuperNodeSetField set a field
func SuperNodeSetField(superNode []byte, pos int, bs []byte, bsStart int, len int) int {
	return NodeSetField(superNode, pos, bs, bsStart, len)
}

//SuperNodeGetField get a field
func SuperNodeGetField(superNode []byte, pos int, fieldLen int) []byte {
	return NodeGetField(superNode, pos, fieldLen)
}

//SuperNodeSetBlkID set the supernode's blkID
func SuperNodeSetBlkID(superNode []byte, blkID int64) {
	bs, _ := conv.Itob(blkID)
	SuperNodeSetField(superNode, int(SuperNodeOffBlkID), bs, 0, int(SuperNodeFieldSizeBlkID))
}

//SuperNodeGetBklID get the blkID
func SuperNodeGetBklID(superNode []byte) (blkID int64) {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffBlkID), int(SuperNodeFieldSizeBlkID))
	blkID, _ = conv.Btoi(bs)
	return
}

//SuperNodeSetRootNodeID set the rootNodeID
func SuperNodeSetRootNodeID(superNode []byte, rootNodeID int64) {
	bs, _ := conv.Itob(rootNodeID)
	SuperNodeSetField(superNode, int(SuperNodeOffRootNodeID), bs, 0, int(SuperNodeFieldSizeRootNodeID))
}

//SuperNodeGetRootNodeID get the rootNodeID
func SuperNodeGetRootNodeID(superNode []byte) (rootNodeID int64) {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffRootNodeID), int(SuperNodeFieldSizeRootNodeID))
	rootNodeID, _ = conv.Btoi(bs)
	return
}

//SuperNodeSetAllPairsNum set the supernode's allPairsNum
func SuperNodeSetAllPairsNum(superNode []byte, allPairsNum int64) {
	bs, _ := conv.Itob(allPairsNum)
	SuperNodeSetField(superNode, int(SuperNodeOffAllPairsNum), bs, 0, int(SuperNodeFieldSizeAllPairsNum))
}

//SuperNodeGetAllPairsNum get the allPairsNum
func SuperNodeGetAllPairsNum(superNode []byte) (pairsNum int64) {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffAllPairsNum), int(SuperNodeFieldSizeAllPairsNum))
	pairsNum, _ = conv.Btoi(bs)
	return
}

//SuperNodeSetNextBlkNum set the supernode's nextBlkNum
func SuperNodeSetNextBlkNum(superNode []byte, nextBlkNum int64) {
	bs, _ := conv.Itob(nextBlkNum)
	SuperNodeSetField(superNode, int(SuperNodeOffNextBlkNum), bs, 0, int(SuperNodeFieldSizeNextBlkNum))
}

//SuperNodeGetNextBlkNum get the nextBlkNum
func SuperNodeGetNextBlkNum(superNode []byte) (nextNum int64) {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffNextBlkNum), int(SuperNodeFieldSizeNextBlkNum))
	nextNum, _ = conv.Btoi(bs)
	return
}

//SuperNodeSetLen set the supernode's len
func SuperNodeSetLen(superNode []byte, len int16) {
	bs, _ := conv.Itob(len)
	SuperNodeSetField(superNode, int(SuperNodeOffLen), bs, 0, int(SuperNodeFieldSizeLen))
}

//SuperNodeGetLen get the len
func SuperNodeGetLen(superNode []byte) int16 {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffLen), int(SuperNodeFieldSizeLen))
	len, _ := conv.Btoi(bs)
	return int16(len)
}

//SuperNodeSetVersion set the supernode's version
func SuperNodeSetVersion(superNode []byte, version int8) {
	bs, _ := conv.Itob(version)
	SuperNodeSetField(superNode, int(SuperNodeOffVersion), bs, 0, int(SuperNodeFieldSizeVersion))
}

//SuperNodeGetVersion get the version
func SuperNodeGetVersion(superNode []byte) int8 {
	bs := SuperNodeGetField(superNode, int(SuperNodeOffVersion), int(SuperNodeFieldSizeLen))
	version, _ := conv.Btoi(bs)
	return int8(version)
}
