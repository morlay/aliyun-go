package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateTemplateRequest struct {
	requests.RpcRequest
	Container            string `position:"Query" name:"Container"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MuxConfig            string `position:"Query" name:"MuxConfig"`
	Video                string `position:"Query" name:"Video"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TemplateId           string `position:"Query" name:"TemplateId"`
	Name                 string `position:"Query" name:"Name"`
	TransConfig          string `position:"Query" name:"TransConfig"`
	Audio                string `position:"Query" name:"Audio"`
}

func (req *UpdateTemplateRequest) Invoke(client *sdk.Client) (resp *UpdateTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateTemplate", "", "")
	resp = &UpdateTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateTemplateResponse struct {
	responses.BaseResponse
	RequestId string
	Template  UpdateTemplateTemplate
}

type UpdateTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   UpdateTemplateContainer
	Video       UpdateTemplateVideo
	Audio       UpdateTemplateAudio
	TransConfig UpdateTemplateTransConfig
	MuxConfig   UpdateTemplateMuxConfig
}

type UpdateTemplateContainer struct {
	Format string
}

type UpdateTemplateVideo struct {
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
	BitrateBnd UpdateTemplateBitrateBnd
}

type UpdateTemplateBitrateBnd struct {
	Max string
	Min string
}

type UpdateTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type UpdateTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type UpdateTemplateMuxConfig struct {
	Segment UpdateTemplateSegment
	Gif     UpdateTemplateGif
}

type UpdateTemplateSegment struct {
	Duration string
}

type UpdateTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
