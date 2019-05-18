package internal

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

type BoardCtl struct{}

func NewBoardCtl() (*BoardCtl, error) {
	return &BoardCtl{}, nil
}

type result struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func (b *BoardCtl) Send(irdata *IrData) error {
	out, err := exec.Command(
		"boardctl",
		"cmd",
		strconv.Itoa(irdata.Interval),
		irdata.Pattern,
	).Output()

	// outputをパースしてみる
	var res result
	jerr := json.Unmarshal(out, &res)

	// パースできなければ
	if jerr != nil {

		if err != nil {
			return err
		} else {
			return fmt.Errorf(string(out))
		}

	}

	if res.Result != "ok" {
		return fmt.Errorf(res.Message)
	}

	return nil
}
