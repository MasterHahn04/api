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

func WhoWhere(client *ts3.Client) ([]WhoWhereStruct, error) {
	var whoWhereList []WhoWhereStruct
	clientList, err := ClientList(client)
	if err != nil {
		return nil, err
	}
	for _, entry := range clientList {
		if entry.Type == 0 {
			var clientInfoWithChannel WhoWhereStruct
			clientInfoWithChannel.ChannelInfo, err = ChannelInfo(client, entry.ChannelId)
			if err != nil {
				return nil, err
			}
			clientInfoWithChannel.ClientList = entry
			whoWhereList = append(whoWhereList, clientInfoWithChannel)
		} else {
			continue
		}
	}
	return whoWhereList, nil
}

func WhoWhereShort(client *ts3.Client) ([]WhoWhereShortStruct, error) {
	var whoWhereShortList []WhoWhereShortStruct
	whoWhereList, err := WhoWhere(client)
	if err != nil {
		return nil, err
	}
	for _, entry := range whoWhereList {
		shortList := WhoWhereShortStruct{
			Nickname:    entry.ClientList.Nickname,
			ChannelName: entry.ChannelInfo.Name,
		}
		whoWhereShortList = append(whoWhereShortList, shortList)
	}
	return whoWhereShortList, nil
}

type WhoWhereShortStruct struct {
	Nickname    string `json:"nickname"`
	ChannelName string `json:"channel_name"`
}
type WhoWhereStruct struct {
	ClientList  ClientListStruct  `json:"client_list"`
	ChannelInfo ChannelInfoStruct `json:"channel_info"`
}

func ChannelInfo(client *ts3.Client, channelId int) (ChannelInfoStruct, error) {
	var channelInfos ChannelInfoStruct
	_, err := client.ExecCmd(ts3.NewCmd(fmt.Sprintf("channelinfo cid=%v", channelId)).WithResponse(&channelInfos))
	if err != nil {
		return ChannelInfoStruct{}, err
	}
	return channelInfos, nil
}

type ChannelInfoStruct struct {
	P_ID                          int    `ms:"pid" json:"p_id"`
	Name                          string `ms:"channel_name" json:"name"`
	Topic                         string `ms:"channel_topic" json:"topic"`
	Description                   string `ms:"channel_description" json:"description"`
	Password                      string `ms:"channel_password" json:"password"`
	Codec                         int    `ms:"channel_codec" json:"codec"`
	CodecQuality                  int    `ms:"channel_codec_quality" json:"codec_quality"`
	MaxClients                    int    `ms:"channel_maxclients" json:"max_clients"`
	MaxFamilyClients              int    `ms:"channel_maxfamilyclients" json:"max_family_clients"`
	Order                         int    `ms:"channel_order" json:"order"`
	FlagPermanent                 int    `ms:"channel_flag_permanent" json:"flag_permanent"`
	FlagSemiPermanent             int    `ms:"channel_flag_semi_permanent" json:"flag_semi_permanent"`
	FlagDefault                   int    `ms:"channel_flag_default" json:"flag_default"`
	FlagPassword                  int    `ms:"channel_flag_password" json:"flag_password"`
	CodecLatencyFactor            int    `ms:"channel_codec_latency_factor" json:"codec_latency_factor"`
	CodecIsUnencrypted            int    `ms:"channel_codec_is_unencrypted" json:"codec_is_unencrypted"`
	SecuritySalt                  string `ms:"channel_security_salt" json:"security_salt"`
	DeleteDelay                   int    `ms:"channel_delete_delay" json:"delete_delay"`
	UniqueIdentifier              string `ms:"channel_unique_identifier" json:"unique_identifier"`
	FlagMaxClientsUnlimited       int    `ms:"channel_flag_maxclients_unlimited" json:"flag_max_clients_unlimited"`
	FlagMaxFamilyClientsUnlimited int    `ms:"channel_flag_maxfamilyclients_unlimited" json:"flag_max_family_clients_unlimited"`
	FlagMaxFamilyClientsInherited int    `ms:"channel_flag_maxfamilyclients_inherited" json:"flag_max_family_clients_inherited"`
	Filepath                      string `ms:"channel_filepath" json:"filepath"`
	NeededTalkPower               int    `ms:"channel_needed_talk_power" json:"needed_talk_power"`
	ForcedSilence                 int    `ms:"channel_forced_silence" json:"forced_silence"`
	NamePhonetic                  string `ms:"channel_name_phonetic" json:"name_phonetic"`
	IconId                        int    `ms:"channel_icon_id" json:"icon_id"`
	BannerGfxUrl                  string `ms:"channel_banner_gfx_url" json:"banner_gfx_url"`
	BannerMode                    int    `ms:"channel_banner_mode" json:"banner_mode"`
	SecondsEmpty                  int    `ms:"seconds_empty" json:"seconds_empty"`
}
