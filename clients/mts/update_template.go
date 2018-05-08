package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateTemplate", "mts", "")
	resp = &UpdateTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateTemplateResponse struct {
	responses.BaseResponse
	RequestId common.String
	Template  UpdateTemplateTemplate
}

type UpdateTemplateTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   UpdateTemplateContainer
	Video       UpdateTemplateVideo
	Audio       UpdateTemplateAudio
	TransConfig UpdateTemplateTransConfig
	MuxConfig   UpdateTemplateMuxConfig
}

type UpdateTemplateContainer struct {
	Format common.String
}

type UpdateTemplateVideo struct {
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
	BitrateBnd UpdateTemplateBitrateBnd
}

type UpdateTemplateBitrateBnd struct {
	Max common.String
	Min common.String
}

type UpdateTemplateAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Remove     common.String
}

type UpdateTemplateTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type UpdateTemplateMuxConfig struct {
	Segment UpdateTemplateSegment
	Gif     UpdateTemplateGif
}

type UpdateTemplateSegment struct {
	Duration common.String
}

type UpdateTemplateGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}
