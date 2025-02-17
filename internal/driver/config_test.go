// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"strings"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

func TestCreateASCIIConnectionInfo_dataBits7(t *testing.T) {
	address := "/dev/USB0tty"
	baudRate := 19200
	dataBits := 7
	stopBits := 1
	parity := "N"
	unitID := uint8(255)
	protocols := map[string]models.ProtocolProperties{
		ProtocolASCII: {
			Address:  address,
			UnitID:   "255",
			BaudRate: "19200",
			DataBits: "7",
			StopBits: "1",
			Parity:   "N",
		},
	}

	connectionInfo, err := createConnectionInfo(protocols)

	if err != nil {
		t.Fatalf("Fail to create connectionInfo. Error: %v", err)
	}
	if connectionInfo.Protocol != ProtocolASCII || connectionInfo.Address != address || connectionInfo.UnitID != unitID ||
		connectionInfo.BaudRate != baudRate || connectionInfo.DataBits != dataBits || connectionInfo.StopBits != stopBits ||
		connectionInfo.Parity != parity {
		t.Fatalf("Unexpect test result. %v should match to %v ", connectionInfo, protocols)
	}
}

func TestCreateRTUConnectionInfo_unitID255(t *testing.T) {
	address := "/dev/USB0tty"
	baudRate := 19200
	dataBits := 8
	stopBits := 1
	parity := "N"
	unitID := uint8(255)
	protocols := map[string]models.ProtocolProperties{
		ProtocolRTU: {
			Address:  address,
			UnitID:   "255",
			BaudRate: "19200",
			DataBits: "8",
			StopBits: "1",
			Parity:   "N",
		},
	}

	connectionInfo, err := createConnectionInfo(protocols)

	if err != nil {
		t.Fatalf("Fail to create connectionInfo. Error: %v", err)
	}
	if connectionInfo.Protocol != ProtocolRTU || connectionInfo.Address != address || connectionInfo.UnitID != unitID ||
		connectionInfo.BaudRate != baudRate || connectionInfo.DataBits != dataBits || connectionInfo.StopBits != stopBits ||
		connectionInfo.Parity != parity {
		t.Fatalf("Unexpect test result. %v should match to %v ", connectionInfo, protocols)
	}
}

func TestCreateConnectionInfo_unitID0(t *testing.T) {
	address := "/dev/USB0tty"
	baudRate := 19200
	dataBits := 8
	stopBits := 1
	parity := "N"
	unitID := uint8(0)
	protocols := map[string]models.ProtocolProperties{
		ProtocolRTU: {
			Address:  address,
			UnitID:   "0",
			BaudRate: "19200",
			DataBits: "8",
			StopBits: "1",
			Parity:   "N",
		},
	}

	connectionInfo, err := createConnectionInfo(protocols)

	if err != nil {
		t.Fatalf("Fail to create connectionInfo. Error: %v", err)
	}
	if connectionInfo.Protocol != ProtocolRTU || connectionInfo.Address != address || connectionInfo.UnitID != unitID ||
		connectionInfo.BaudRate != baudRate || connectionInfo.DataBits != dataBits || connectionInfo.StopBits != stopBits ||
		connectionInfo.Parity != parity {
		t.Fatalf("Unexpect test result. %v should match to %v ", connectionInfo, protocols)
	}
}

func TestCreateConnectionInfo_unitIdOutOfRange(t *testing.T) {
	address := "/dev/USB0tty"
	unitID := "256"
	protocols := map[string]models.ProtocolProperties{
		ProtocolRTU: {
			Address:  address,
			UnitID:   unitID,
			BaudRate: "19200",
			DataBits: "8",
			StopBits: "1",
			Parity:   "N",
		},
	}

	_, err := createConnectionInfo(protocols)

	if err == nil || !strings.Contains(err.Error(), "value out of range") {
		t.Fatalf("Unexpect test result, unitID %v should out of ranage, %v", unitID, err)
	}
}

func TestCreateConnectionInfo_invalidParity(t *testing.T) {
	address := "/dev/USB0tty"
	parity := "invalid-parity"
	protocols := map[string]models.ProtocolProperties{
		ProtocolRTU: {
			Address:  address,
			UnitID:   "1",
			BaudRate: "19200",
			DataBits: "8",
			StopBits: "1",
			Parity:   parity,
		},
	}

	_, err := createConnectionInfo(protocols)

	if err == nil || !strings.Contains(err.Error(), "invalid parity value, it should be N(None) or O(Odd) or E(Even)") {
		t.Fatalf("Unexpect test result, parity %v should be invalid, %v", parity, err)
	}
}

func TestCreateTCPConnectionInfo(t *testing.T) {
	address := "0.0.0.0"
	port := 502
	unitID := uint8(255)
	protocols := map[string]models.ProtocolProperties{
		ProtocolTCP: {
			Address: address,
			Port:    "502",
			UnitID:  "255",
		},
	}

	connectionInfo, err := createConnectionInfo(protocols)

	if err != nil {
		t.Fatalf("Fail to create connectionInfo. Error: %v", err)
	}
	if connectionInfo.Protocol != ProtocolTCP || connectionInfo.Address != address ||
		connectionInfo.Port != port || connectionInfo.UnitID != unitID {
		t.Fatalf("Unexpect test result. %v should match to %v ", connectionInfo, protocols)
	}
}

func TestCreateTCPConnectionInfo_unitIdOutOfRange(t *testing.T) {
	address := "0.0.0.0"
	unitID := "256"
	protocols := map[string]models.ProtocolProperties{
		ProtocolTCP: {
			Address: address,
			Port:    "502",
			UnitID:  unitID,
		},
	}

	_, err := createConnectionInfo(protocols)

	if err == nil || !strings.Contains(err.Error(), "value out of range") {
		t.Fatalf("Unexpect test result, unitID %v should out of ranage, %v", unitID, err)
	}
}

func TestCreateTCPConnectionInfo_portOutOfRange(t *testing.T) {
	address := "0.0.0.0"
	port := "65536"
	protocols := map[string]models.ProtocolProperties{
		ProtocolTCP: {
			Address: address,
			Port:    port,
			UnitID:  "1",
		},
	}

	_, err := createConnectionInfo(protocols)

	if err == nil || !strings.Contains(err.Error(), "value out of range") {
		t.Fatalf("Unexpect test result, port %v should out of ranage, %v", port, err)
	}
}
