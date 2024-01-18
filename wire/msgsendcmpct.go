package wire

import "io"

// MsgSendCmpct wireshark判定未知的报文
type MsgSendCmpct struct {

}

func (m *MsgSendCmpct) BtcDecode(reader io.Reader, u uint32, encoding MessageEncoding) error {
	return nil
}

func (m *MsgSendCmpct) BtcEncode(writer io.Writer, u uint32, encoding MessageEncoding) error {
	return nil
}

func (m *MsgSendCmpct) Command() string {
	return CmdSendCmpct
}

func (m *MsgSendCmpct) MaxPayloadLength(u uint32) uint32 {
  	return 0
}

func NewMsgSendCmpct() *MsgSendCmpct {
	return &MsgSendCmpct{}
}