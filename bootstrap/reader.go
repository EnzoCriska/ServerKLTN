//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package bootstrap

import (
	"errors"
	"net/http"

	"github.com/mainflux/mainflux"
)

// bootstrapRes represent Mainflux Response to the Bootatrap request.
// This is used as a response from ConfigReader and can easily be
// replace with any other response format.
type bootstrapRes struct {
	MFThing    string       `json:"mainflux_id"`
	MFKey      string       `json:"mainflux_key"`
	MFChannels []channelRes `json:"mainflux_channels"`
	Content    string       `json:"content"`
}

type channelRes struct {
	ID       string      `json:"id"`
	Name     string      `json:"name,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

func (res bootstrapRes) Code() int {
	return http.StatusOK
}

func (res bootstrapRes) Headers() map[string]string {
	return map[string]string{}
}

func (res bootstrapRes) Empty() bool {
	return false
}

type reader struct{}

// NewConfigReader return new reader which is used to generate response
// from the config.
func NewConfigReader() ConfigReader {
	return reader{}
}

func (r reader) ReadConfig(cfg Config) (mainflux.Response, error) {
	if len(cfg.MFChannels) < 1 {
		return bootstrapRes{}, errors.New("Invalid configuration")
	}

	var channels []channelRes
	for _, ch := range cfg.MFChannels {
		channels = append(channels, channelRes{ID: ch.ID, Name: ch.Name, Metadata: ch.Metadata})
	}

	res := bootstrapRes{
		MFKey:      cfg.MFKey,
		MFThing:    cfg.MFThing,
		MFChannels: channels,
		Content:    cfg.Content,
	}

	return res, nil
}
