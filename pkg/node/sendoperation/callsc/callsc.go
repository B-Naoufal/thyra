package callsc

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/massalabs/thyra/pkg/node/base58"
)

const CallSCOpID = uint64(4)

//nolint:tagliatelle
type OperationDetails struct {
	MaxGaz     int64       `json:"max_gas"`
	GazPrice   string      `json:"gas_price"`
	Coins      string      `json:"coins"`
	TargetAddr string      `json:"target_addr"`
	TargetFunc string      `json:"target_func"`
	Param      interface{} `json:"param"`
}

//nolint:tagliatelle
type Operation struct {
	CallSC OperationDetails `json:"CallSC"`
}

type CallSC struct {
	address    []byte
	function   string
	parameters []byte
	gazPrice   uint64
	gazLimit   uint64
	coins      uint64
}

func New(address []byte, function string, parameters []byte, gazPrice uint64, gazLimit uint64, coins uint64,
) *CallSC {
	return &CallSC{
		address: address, function: function, parameters: parameters,
		gazPrice: gazPrice, gazLimit: gazLimit, coins: coins,
	}
}

func (c *CallSC) Content() interface{} {
	return &Operation{
		CallSC: OperationDetails{
			MaxGaz:     int64(c.gazLimit),
			GazPrice:   fmt.Sprint(c.gazPrice),
			Coins:      fmt.Sprint(c.coins),
			TargetAddr: "A" + base58.CheckEncode(append(make([]byte, 1), c.address...)),
			TargetFunc: c.function,
			Param:      hex.EncodeToString(c.parameters),
		},
	}
}

func (c *CallSC) Message() []byte {
	msg := make([]byte, 0)
	buf := make([]byte, binary.MaxVarintLen64)

	// operationId
	nbBytes := binary.PutUvarint(buf, CallSCOpID)
	msg = append(msg, buf[:nbBytes]...)

	// maxGaz
	nbBytes = binary.PutUvarint(buf, c.gazLimit)
	msg = append(msg, buf[:nbBytes]...)

	// Coins
	nbBytes = binary.PutUvarint(buf, c.coins)
	msg = append(msg, buf[:nbBytes]...)

	// gazPrice
	nbBytes = binary.PutUvarint(buf, c.gazPrice)
	msg = append(msg, buf[:nbBytes]...)

	// target address
	msg = append(msg, c.address...)

	// target function
	nbBytes = binary.PutUvarint(buf, uint64(len([]byte(c.function))))
	msg = append(msg, buf[:nbBytes]...)
	msg = append(msg, []byte(c.function)...)

	// param
	nbBytes = binary.PutUvarint(buf, uint64(len(c.parameters)))
	msg = append(msg, buf[:nbBytes]...)
	msg = append(msg, c.parameters...)

	return msg
}
