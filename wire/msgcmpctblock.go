package wire

import "io"

type MsgCmpctBlock struct {
	BlockHeader *BlockHeader
	Nonce uint64
	ShortIds []byte
	PrefilledTransaction []*MsgTx
}

func (msg *MsgCmpctBlock)AddShortID(id byte){
	msg.ShortIds = append(msg.ShortIds, id)
}

func (msg *MsgCmpctBlock)AddTransaction(tx *MsgTx){
	msg.PrefilledTransaction = append(msg.PrefilledTransaction, tx)
}

func (msg *MsgCmpctBlock) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

func (msg *MsgCmpctBlock) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return nil
}

func (msg *MsgCmpctBlock) Command() string {
	return CmdCmpctBlock
}

func (msg *MsgCmpctBlock) MaxPayloadLength(pver uint32) uint32 {
	return MaxBlockPayload
}

func NewMsgCmpctBlock(blockheader BlockHeader) *MsgCmpctBlock {
	return &MsgCmpctBlock{
		BlockHeader: &blockheader,
		ShortIds: make([]byte, 0),
		PrefilledTransaction: make([]*MsgTx, 0),
	}
}