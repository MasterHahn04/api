package teamspeak

import (
	"fmt"
	"github.com/multiplay/go-ts3"
)

func Connect(addr string, port string, user string, passwd string, server int) (*ts3.Client, error) {
	client, err := ts3.NewClient(addr + ":" + port)
	if err != nil {
		return nil, err
	}
	err = client.Login(user, passwd)
	if err != nil {
		return nil, err
	}
	err = client.Use(1)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ClientList(client *ts3.Client) ([]ClientListStruct, error) {
	list, err := client.Server.ClientList()
	if err != nil {
		return nil, err
	}
	var clientListFull []ClientListStruct
	for _, entry := range list {
		client := ClientListStruct{
			ClientId:     entry.ID,
			ChannelId:    entry.ChannelID,
			DatabaseId:   entry.DatabaseID,
			Nickname:     entry.Nickname,
			Type:         entry.Type,
			IsClientAway: entry.Away,
			AwayMessage:  entry.AwayMessage,
		}
		clientListFull = append(clientListFull, client)
	}
	return clientListFull, nil
}

type ClientListStruct struct {
	ClientId     int    `json:"client_id"`
	ChannelId    int    `json:"channel_id"`
	DatabaseId   int    `json:"database_id"`
	Nickname     string `json:"nickname"`
	Type         int    `json:"type"`
	IsClientAway bool   `json:"is_client_away"`
	AwayMessage  string `json:"away_message"`
}

func WhoWhere(client *ts3.Client) {
	clientList, err := ClientList(client)
	if err != nil {

	}
	for _, entry := range clientList {
		if entry.Type == 1 {

		} else {
			continue
		}
	}
}

type WhoWhereStruct struct {
	ClientList ClientListStruct
}

func ChannelInfo(client *ts3.Client, channelId int) {
	list, err := client.ExecCmd(ts3.NewCmd(fmt.Sprintf("channelinfo cid=%v", channelId)).WithResponse())
	if err != nil {

	}

}

type ChannelInfo struct {
	P_ID int `ms:"pid" json:"p_ID"`

	pid=19 channel_name=Mannschaftsheim\sV channel_topic channel_description channel_password channel_codec=4 channel_codec_quality=6 channel_maxclients=-1 channel_maxfamilyclients=-1 channel_order=23 channel_flag_permanent=1 channel_flag_semi_permanent=0 channel_flag_default=0 channel_flag_password=0 channel_codec_latency_factor=1 channel_codec_is_unencrypted=0 channel_security_salt channel_delete_delay=0 channel_unique_identifier=0fe7005e-c197-419e-aca8-f88adc11ed53 channel_flag_maxclients_unlimited=1 channel_flag_maxfamilyclients_unlimited=1 channel_flag_maxfamilyclients_inherited=0 channel_filepath=files\/virtualserver_1\/channel_24 channel_needed_talk_power=0 channel_forced_silence=0 channel_name_phonetic=Mannschaftsheim\s5 channel_icon_id=2053633368 channel_banner_gfx_url channel_banner_mode=0 seconds_empty=785237
}