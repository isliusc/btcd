package wire

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io"
)

type MsgBlockTxn struct {
	Hash    chainhash.Hash
	Transactions []*MsgTx
}

func (msg *MsgBlockTxn)AddTransaction(tx *MsgTx){
	msg.Transactions = append(msg.Transactions, tx)
}

func (msg *MsgBlockTxn)SetBlockHash(hash []byte) error{
	err := msg.Hash.SetBytes(hash)
	if err != nil {
		return err
	}
	return nil
}

func (msg *MsgBlockTxn) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {

	err := readElement(r, &msg.Hash)
	if err != nil {
		return err
	}
	txCount, err := ReadVarInt(r, pver)
	if err != nil {
		return err
	}

	msg.Transactions = make([]*MsgTx, 0, txCount)
	for i := uint64(0); i < txCount; i++ {
		tx := MsgTx{}
		err := tx.BtcDecode(r, pver, enc)
		if err != nil {
			return err
		}
		msg.Transactions = append(msg.Transactions, &tx)
	}

	return nil
}

func (msg *MsgBlockTxn) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	err := writeElement(w, msg.Hash)
	if err != nil {
		return err
	}
	err = WriteVarInt(w, pver, uint64(len(msg.Transactions)))
	if err != nil {
		return err
	}

	for _, tx := range msg.Transactions {
		err = tx.BtcEncode(w, pver, enc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (msg *MsgBlockTxn) Command() string {
	return CmdBlockTxn
}

func (msg *MsgBlockTxn) MaxPayloadLength(u uint32) uint32 {
	return MaxBlockPayload
}

func NewMsgBlockTxn() *MsgBlockTxn{
	return &MsgBlockTxn{
		Hash: chainhash.Hash{},
		Transactions: make([]*MsgTx, 0),
	}
}