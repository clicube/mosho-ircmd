package internal

import (
	"log"
)

type Cmd struct {
	Id   int64
	Name string
}

type IrData struct {
	Interval int
	Pattern  string
}

type CmdRepo interface {
	Next() (*Cmd, error)
	Remove(*Cmd) error
}

type IrDataRepo interface {
	Get(string) (*IrData, error)
}

type Sender interface {
	Send(*IrData) error
}

func Start() error {
	var cmdrepo CmdRepo
	var irdatarepo IrDataRepo
	var err error
	var sender Sender
	cmdrepo, err = NewFirestore()
	if err != nil {
		return err
	}

	irdatarepo, err = NewDataFile("ir_pattern.json")
	if err != nil {
		return err
	}

	sender, err = NewBoardCtl()
	if err != nil {
		return err
	}

	for {
		cmd, err := cmdrepo.Next()
		if err != nil {
			return err
		}
		log.Println(cmd)

		irdata, err := irdatarepo.Get(cmd.Name)
		if err == nil {

			err = sender.Send(irdata)
			if err != nil {
				return err
			}

		}

		cmdrepo.Remove(cmd)
		if err != nil {
			return err
		}

	}
}
