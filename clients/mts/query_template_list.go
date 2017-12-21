package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTemplateListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateIds          string `position:"Query" name:"TemplateIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryTemplateListRequest) Invoke(client *sdk.Client) (response *QueryTemplateListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryTemplateListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTemplateList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryTemplateListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryTemplateListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryTemplateListResponse struct {
	RequestId    string
	TemplateList QueryTemplateListTemplateList
	NonExistTids QueryTemplateListNonExistTidList
}

type QueryTemplateListTemplate struct {
	Id          string
	Name        string
	State       string
	Container   QueryTemplateListContainer
	Video       QueryTemplateListVideo
	Audio       QueryTemplateListAudio
	TransConfig QueryTemplateListTransConfig
	MuxConfig   QueryTemplateListMuxConfig
}

type QueryTemplateListContainer struct {
	Format string
}

type QueryTemplateListVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Remove     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd QueryTemplateListBitrateBnd
}

type QueryTemplateListBitrateBnd struct {
	Max string
	Min string
}

type QueryTemplateListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type QueryTemplateListTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type QueryTemplateListMuxConfig struct {
	Segment QueryTemplateListSegment
	Gif     QueryTemplateListGif
}

type QueryTemplateListSegment struct {
	Duration string
}

type QueryTemplateListGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type QueryTemplateListTemplateList []QueryTemplateListTemplate

func (list *QueryTemplateListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTemplateListTemplate)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryTemplateListNonExistTidList []string

func (list *QueryTemplateListNonExistTidList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
