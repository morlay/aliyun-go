package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *SearchTemplateRequest) Invoke(client *sdk.Client) (resp *SearchTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchTemplate", "mts", "")
	resp = &SearchTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchTemplateResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Long
	PageNumber   common.Long
	PageSize     common.Long
	TemplateList SearchTemplateTemplateList
}

type SearchTemplateTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   SearchTemplateContainer
	Video       SearchTemplateVideo
	Audio       SearchTemplateAudio
	TransConfig SearchTemplateTransConfig
	MuxConfig   SearchTemplateMuxConfig
}

type SearchTemplateContainer struct {
	Format common.String
}

type SearchTemplateVideo struct {
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
	BitrateBnd SearchTemplateBitrateBnd
}

type SearchTemplateBitrateBnd struct {
	Max common.String
	Min common.String
}

type SearchTemplateAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Remove     common.String
}

type SearchTemplateTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type SearchTemplateMuxConfig struct {
	Segment SearchTemplateSegment
	Gif     SearchTemplateGif
}

type SearchTemplateSegment struct {
	Duration common.String
}

type SearchTemplateGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type SearchTemplateTemplateList []SearchTemplateTemplate

func (list *SearchTemplateTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchTemplateTemplate)
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
