package wire

import "io"

type MsgGetBlockTxn struct {

}

func (m MsgGetBlockTxn) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

func (m MsgGetBlockTxn) BtcEncode(writer io.Writer, u uint32, encoding MessageEncoding) error {
	return nil
}

func (m MsgGetBlockTxn) Command() string {
	return CmdGetBlockTxn
}

func (m MsgGetBlockTxn) MaxPayloadLength(u uint32) uint32 {
	return 0
}

func NewMsgGetBlockTxn() *MsgGetBlockTxn {
	return &MsgGetBlockTxn{}
}