package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddTemplateRequest struct {
	requests.RpcRequest
	Container            string `position:"Query" name:"Container"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	TransConfig          string `position:"Query" name:"TransConfig"`
	MuxConfig            string `position:"Query" name:"MuxConfig"`
	Video                string `position:"Query" name:"Video"`
	Audio                string `position:"Query" name:"Audio"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AddTemplateRequest) Invoke(client *sdk.Client) (resp *AddTemplateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddTemplate", "mts", "")
	resp = &AddTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddTemplateResponse struct {
	responses.BaseResponse
	RequestId common.String
	Template  AddTemplateTemplate
}

type AddTemplateTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   AddTemplateContainer
	Video       AddTemplateVideo
	Audio       AddTemplateAudio
	TransConfig AddTemplateTransConfig
	MuxConfig   AddTemplateMuxConfig
}

type AddTemplateContainer struct {
	Format common.String
}

type AddTemplateVideo struct {
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
	BitrateBnd AddTemplateBitrateBnd
}

type AddTemplateBitrateBnd struct {
	Max common.String
	Min common.String
}

type AddTemplateAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Remove     common.String
	Volume     AddTemplateVolume
}

type AddTemplateVolume struct {
	Level  common.String
	Method common.String
}

type AddTemplateTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type AddTemplateMuxConfig struct {
	Segment AddTemplateSegment
	Gif     AddTemplateGif
}

type AddTemplateSegment struct {
	Duration common.String
}

type AddTemplateGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}
