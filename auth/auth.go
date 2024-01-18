package auth

import (
	"errors"
	"github.com/panguim-x/protocol/constant"
)

func (x *UserTokenReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.PlatformID > constant.AdminPlatformID || x.PlatformID < constant.IOSPlatformID {
		return errors.New("platform is invalidate")
	}
	return nil
}

func (x *ForceLogoutReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.PlatformID > constant.AdminPlatformID || x.PlatformID < constant.IOSPlatformID {
		return errors.New("platformID is invalidate")
	}
	return nil
}

func (x *ParseTokenReq) Check() error {
	if x.Token == "" {
		return errors.New("userID is empty")
	}
	return nil
}
