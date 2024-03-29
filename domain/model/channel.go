package model

type Channel struct {
	ID                int64  `json:"id"`
	ChannelURL        string `json:"channelURL"`
	ChannelID         string `json:"channelID"`
	ChannelName       string `json:"channelName"`
	UniqueID          string `json:"uniqueID"`
	ChannelType       string `json:"channelType"`
	ManualChannelName string `json:"manualChannelName"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
