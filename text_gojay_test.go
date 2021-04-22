// SPDX-FileCopyrightText: Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestDidOpenTextDocumentParams(t *testing.T) {
	testDidOpenTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	testDidChangeTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	testTextDocumentContentChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	testTextDocumentChangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	testWillSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	testDidSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	testTextDocumentSaveRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	testDidCloseTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
