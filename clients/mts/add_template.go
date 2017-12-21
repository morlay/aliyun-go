package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTemplateRequest struct {
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

func (r AddTemplateRequest) Invoke(client *sdk.Client) (response *AddTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		AddTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddTemplateResponse struct {
	RequestId string
	Template  AddTemplateTemplate
}

type AddTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   AddTemplateContainer
	Video       AddTemplateVideo
	Audio       AddTemplateAudio
	TransConfig AddTemplateTransConfig
	MuxConfig   AddTemplateMuxConfig
}

type AddTemplateContainer struct {
	Format string
}

type AddTemplateVideo struct {
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
	BitrateBnd AddTemplateBitrateBnd
}

type AddTemplateBitrateBnd struct {
	Max string
	Min string
}

type AddTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
	Volume     AddTemplateVolume
}

type AddTemplateVolume struct {
	Level  string
	Method string
}

type AddTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type AddTemplateMuxConfig struct {
	Segment AddTemplateSegment
	Gif     AddTemplateGif
}

type AddTemplateSegment struct {
	Duration string
}

type AddTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
