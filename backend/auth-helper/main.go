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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: firegate-auth <user> <password>")
		os.Exit(2)
	}
	username, password := os.Args[1], os.Args[2]
	
	data, err := os.ReadFile("/etc/shadow")
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error")
		os.Exit(1)
	}
	
	for _,line := range strings.Split(string(data), "\n") {
		parts := strings.SplitN(line, ":", 8)
		if len(parts) < 2 || parts[0] != username {
			continue
		}
		
		hash := parts[1]
		if hash == "!" || hash == "*" || hash == "" {
			os.Exit(1)
		}
		
		fields := strings.SplitN(hash, "$", 4)
		if len(fields) != 4 {
			os.Exit(1)
		}
		
		cmd := exec.Command("openssl", "passwd", "-"+fields[1], "-salt", fields[2], "-stdin")
		cmd.Stdin = strings.NewReader(password + "\n")
		out, err := cmd.Output()
		if err != nil || string(out) != hash+"\n" {
			os.Exit(1)
		}
		
		os.Exit(0)
	}
	os.Exit(1)
}
