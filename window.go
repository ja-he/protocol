// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "strconv"

// ShowMessageParams params of ShowMessage Notification.
type ShowMessageParams struct {
	// Message is the actual message.
	Message string `json:"message"`

	// Type is the message type.
	Type MessageType `json:"type"`
}

// MessageType type of ShowMessageParams type.
type MessageType float64

const (
	// MessageTypeError an error message.
	MessageTypeError MessageType = 1
	// MessageTypeWarning a warning message.
	MessageTypeWarning MessageType = 2
	// MessageTypeInfo an information message.
	MessageTypeInfo MessageType = 3
	// MessageTypeLog a log message.
	MessageTypeLog MessageType = 4
)

// String implements fmt.Stringer.
func (m MessageType) String() string {
	switch m {
	case MessageTypeError:
		return "error"
	case MessageTypeWarning:
		return "warning"
	case MessageTypeInfo:
		return "info"
	case MessageTypeLog:
		return "log"
	default:
		return strconv.FormatFloat(float64(m), 'f', -10, 64)
	}
}

// Enabled reports whether the level is enabled.
func (m MessageType) Enabled(level MessageType) bool {
	return level > 0 && m >= level
}

// messageTypeMap map of MessageTypes.
var messageTypeMap = map[string]MessageType{
	"error":   MessageTypeError,
	"warning": MessageTypeWarning,
	"info":    MessageTypeInfo,
	"log":     MessageTypeLog,
}

// ToMessageType converts level to the MessageType.
func ToMessageType(level string) MessageType {
	mt, ok := messageTypeMap[level]
	if !ok {
		return MessageType(0) // unknown
	}

	return mt
}

// ShowMessageRequestParams params of ShowMessage Request.
type ShowMessageRequestParams struct {
	// Actions is the message action items to present.
	Actions []MessageActionItem `json:"actions"`

	// Message is the actual message
	Message string `json:"message"`

	// Type is the message type. See {@link MessageType}
	Type MessageType `json:"type"`
}

// MessageActionItem item of ShowMessageRequestParams action.
type MessageActionItem struct {
	// Title a short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

// LogMessageParams params of LogMessage Notification.
type LogMessageParams struct {
	// Message is the actual message
	Message string `json:"message"`

	// Type is the message type. See {@link MessageType}
	Type MessageType `json:"type"`
}

// WorkDoneProgressCreateParams params of WorkDoneProgressCreate request.
//
// @since 3.15.0.
type WorkDoneProgressCreateParams struct {
	// Token is the token to be used to report progress.
	Token ProgressToken `json:"token"`
}

// WorkDoneProgressCreateParams params of WorkDoneProgressCancel request.
//
// @since 3.15.0.
type WorkDoneProgressCancelParams struct {
	// Token is the token to be used to report progress.
	Token ProgressToken `json:"token"`
}
