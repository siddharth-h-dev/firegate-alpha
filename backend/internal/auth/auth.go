// Copyright (C) 2026 Siddharth H
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

package auth

import (
	"fmt"
	"os/exec"
)

const authBinary = "/usr/local/bin/firegate-auth"

func Verify(username, password string) error {
	cmd := exec.Command(authBinary, username, password)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("auth failed: exit code %d", exitErr.ExitCode())
		}
		return fmt.Errorf("failed to run auth helper: %w", err)
	}
	return nil
}
