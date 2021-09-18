package id_gen

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
	machineID uint16
)

func getMachineID() (uint16, error) {
	return machineID, nil
}

func Init(mID uint16) (err error) {
	machineID = mID
	st := sonyflake.Settings{}
	st.MachineID = getMachineID
	sonyFlake = sonyflake.NewSonyflake(st)
	return
}
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("must Call Init before GetID,err%v", err)
		return
	}
	return sonyFlake.NextID()
}
