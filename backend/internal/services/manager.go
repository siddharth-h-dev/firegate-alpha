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

package services

import (
	"os/exec"
	"fmt"
	"path/filepath"
	"firegate/internal/git"
)

type ServiceType string

const (
	ServiceNftables ServiceType = "nftables"
	ServiceUnbound ServiceType = "unbound"
	ServiceSuricata ServiceType = "suricata"
	ServiceTor ServiceType = "tor"
	ServiceMOTD ServiceType = "motd"
)

func ApplyConfig(service ServiceType, msg string) error {
	repoPath := git.GetRepoPath(string(service))
	configFile := filepath.Join(repoPath, "environments", "default", getConfigFilename(service))
	
	// Validate
	if err := validate(service, configFile); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	
	// Apply
	if err := apply(service, configFile); err != nil {
		return fmt.Errorf("apply failed: %w", err)
	}
	
	// Commit
	return git.CommitChanges(string(service), msg)
}

func getConfigFilename(s ServiceType) string {
	switch s {
	case ServiceNftables: return "nftables.nft"
	case ServiceUnbound: return "unbound.conf"
	case ServiceSuricata: return "suricata.yaml"
	case ServiceTor: return "torrc"
	case ServiceMOTD: return "motd"
	default: return ""
	}
}

func validate(s ServiceType, path string) error {
	switch s {
	case ServiceNftables:
		cmd := exec.Command("nft", "-c", "-f", path)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	case ServiceUnbound:
		cmd := exec.Command("unbound-checkconf", path)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	}	
	return nil
}

func apply(s ServiceType, path string) error {
	switch s {
	case ServiceNftables:
		cmd := exec.Command("nft", "-f", path)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	case ServiceUnbound:
		cmd := exec.Command("doas", "/sbin/rc-service", "unbound", "restart")
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	case ServiceTor:
		cmd := exec.Command("doas", "/sbin/rc-service", "tor", "restart")
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	case ServiceSuricata:
		cmd := exec.Command("doas", "/sbin/rc-service", "suricata", "restart")
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s", string(out))
		}
	}
	return nil
}	
