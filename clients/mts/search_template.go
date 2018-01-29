package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	TemplateList SearchTemplateTemplateList
}

type SearchTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   SearchTemplateContainer
	Video       SearchTemplateVideo
	Audio       SearchTemplateAudio
	TransConfig SearchTemplateTransConfig
	MuxConfig   SearchTemplateMuxConfig
}

type SearchTemplateContainer struct {
	Format string
}

type SearchTemplateVideo struct {
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
	BitrateBnd SearchTemplateBitrateBnd
}

type SearchTemplateBitrateBnd struct {
	Max string
	Min string
}

type SearchTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type SearchTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type SearchTemplateMuxConfig struct {
	Segment SearchTemplateSegment
	Gif     SearchTemplateGif
}

type SearchTemplateSegment struct {
	Duration string
}

type SearchTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
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
