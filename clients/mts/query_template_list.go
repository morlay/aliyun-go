package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTemplateListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateIds          string `position:"Query" name:"TemplateIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryTemplateListRequest) Invoke(client *sdk.Client) (resp *QueryTemplateListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTemplateList", "mts", "")
	resp = &QueryTemplateListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTemplateListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TemplateList QueryTemplateListTemplateList
	NonExistTids QueryTemplateListNonExistTidList
}

type QueryTemplateListTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   QueryTemplateListContainer
	Video       QueryTemplateListVideo
	Audio       QueryTemplateListAudio
	TransConfig QueryTemplateListTransConfig
	MuxConfig   QueryTemplateListMuxConfig
}

type QueryTemplateListContainer struct {
	Format common.String
}

type QueryTemplateListVideo struct {
	Codec      common.String
	Profile    common.String
	Bitrate    common.String
	Crf        common.String
	Width      common.String
	Height     common.String
	Fps        common.String
	Gop        common.String
	Preset     common.String
	ScanMode   common.String
	Bufsize    common.String
	Maxrate    common.String
	PixFmt     common.String
	Degrain    common.String
	Qscale     common.String
	Remove     common.String
	Crop       common.String
	Pad        common.String
	MaxFps     common.String
	BitrateBnd QueryTemplateListBitrateBnd
}

type QueryTemplateListBitrateBnd struct {
	Max common.String
	Min common.String
}

type QueryTemplateListAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Remove     common.String
}

type QueryTemplateListTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type QueryTemplateListMuxConfig struct {
	Segment QueryTemplateListSegment
	Gif     QueryTemplateListGif
}

type QueryTemplateListSegment struct {
	Duration common.String
}

type QueryTemplateListGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
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

type QueryTemplateListNonExistTidList []common.String

func (list *QueryTemplateListNonExistTidList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
