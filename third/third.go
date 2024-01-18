package third

import (
	"errors"
	"protocol/constant"
)

func (x *FcmUpdateTokenReq) Check() error {
	if x.PlatformID > constant.AdminPlatformID || x.PlatformID < constant.IOSPlatformID {
		return errors.New("platformID is invalidate")
	}
	if x.FcmToken == "" {
		return errors.New("FcmToken is empty")
	}
	if x.Account == "" {
		return errors.New("account is empty")
	}
	return nil
}

func (x *SetAppBadgeReq) Check() error {
	if x.UserID == "" {
		return errors.New("UserID is empty")
	}
	return nil
}
