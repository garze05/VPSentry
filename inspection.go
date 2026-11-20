/*
***************************
     Inspection Module
***************************

This module provides information to other modules
to verify current status of the different procedures
indicating if they are present or not
*/

package main

import (
	"errors"
	"os/user"
	"slices"
)

type UserState struct {
	UserExists bool
	IsUserSudo bool
}

func GetUserState(username string) (UserState, error) {
	serverUser, err := user.Lookup(username)
	if err != nil {

		var unknownErr user.UnknownUserError

		if errors.As(err, &unknownErr) {
			return UserState{
				UserExists: false,
				IsUserSudo: false,
			}, nil
		}

		return UserState{}, err
	}

	groupIds, err := serverUser.GroupIds()
	if err != nil {
		return UserState{}, err
	}

	sudoGroup, err := user.LookupGroup("sudo")
	if err != nil {
		var unknownGroupErr user.UnknownGroupError
		
		// Group does not exist in system
		if errors.As(err, &unknownGroupErr) {
			return UserState{UserExists: true, IsUserSudo: false}, nil
		}
		return UserState{}, err
	}

	isSudo := slices.Contains(groupIds, sudoGroup.Gid) || serverUser.Uid == "0"

	return UserState{
		UserExists: true,
		IsUserSudo: isSudo,
	}, nil
}
