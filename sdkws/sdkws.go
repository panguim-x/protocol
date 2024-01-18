package sdkws

import (
	"errors"
)

func (x *MsgData) Check() error {
	if x.SendID == "" {
		return errors.New("sendID is empty")
	}
	if x.Content == nil {
		return errors.New("content is empty")
	}
	return nil
}
