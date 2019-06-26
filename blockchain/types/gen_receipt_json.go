// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		Status          hexutil.Uint   `json:"status"`
		Bloom           Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs            []*Log         `json:"logs"              gencodec:"required"`
		TxHash          common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress common.Address `json:"contractAddress"`
		GasUsed         hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
	}
	var enc Receipt
	enc.Status = hexutil.Uint(r.Status)
	enc.Bloom = r.Bloom
	enc.Logs = r.Logs
	enc.TxHash = r.TxHash
	enc.ContractAddress = r.ContractAddress
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		Status          *hexutil.Uint   `json:"status"`
		Bloom           *Bloom          `json:"logsBloom"         gencodec:"required"`
		Logs            []*Log          `json:"logs"              gencodec:"required"`
		TxHash          *common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress *common.Address `json:"contractAddress"`
		GasUsed         *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Status != nil {
		r.Status = uint(*dec.Status)
	}
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*dec.GasUsed)
	return nil
}
