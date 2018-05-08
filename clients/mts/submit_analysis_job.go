package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitAnalysisJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AnalysisConfig       string `position:"Query" name:"AnalysisConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitAnalysisJobRequest) Invoke(client *sdk.Client) (resp *SubmitAnalysisJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitAnalysisJob", "mts", "")
	resp = &SubmitAnalysisJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAnalysisJobResponse struct {
	responses.BaseResponse
	RequestId   common.String
	AnalysisJob SubmitAnalysisJobAnalysisJob
}

type SubmitAnalysisJobAnalysisJob struct {
	Id               common.String
	UserData         common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	CreationTime     common.String
	PipelineId       common.String
	Priority         common.String
	TemplateList     SubmitAnalysisJobTemplateList
	InputFile        SubmitAnalysisJobInputFile
	AnalysisConfig   SubmitAnalysisJobAnalysisConfig
	MNSMessageResult SubmitAnalysisJobMNSMessageResult
}

type SubmitAnalysisJobTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   SubmitAnalysisJobContainer
	Video       SubmitAnalysisJobVideo
	Audio       SubmitAnalysisJobAudio
	TransConfig SubmitAnalysisJobTransConfig
	MuxConfig   SubmitAnalysisJobMuxConfig
}

type SubmitAnalysisJobContainer struct {
	Format common.String
}

type SubmitAnalysisJobVideo struct {
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
	BitrateBnd SubmitAnalysisJobBitrateBnd
}

type SubmitAnalysisJobBitrateBnd struct {
	Max common.String
	Min common.String
}

type SubmitAnalysisJobAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
}

type SubmitAnalysisJobTransConfig struct {
	TransMode common.String
}

type SubmitAnalysisJobMuxConfig struct {
	Segment SubmitAnalysisJobSegment
	Gif     SubmitAnalysisJobGif
}

type SubmitAnalysisJobSegment struct {
	Duration common.String
}

type SubmitAnalysisJobGif struct {
	Loop       common.String
	FinalDelay common.String
}

type SubmitAnalysisJobInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitAnalysisJobAnalysisConfig struct {
	QualityControl    SubmitAnalysisJobQualityControl
	PropertiesControl SubmitAnalysisJobPropertiesControl
}

type SubmitAnalysisJobQualityControl struct {
	RateQuality     common.String
	MethodStreaming common.String
}

type SubmitAnalysisJobPropertiesControl struct {
	Deinterlace common.String
	Crop        SubmitAnalysisJobCrop
}

type SubmitAnalysisJobCrop struct {
	Mode   common.String
	Width  common.String
	Height common.String
	Top    common.String
	Left   common.String
}

type SubmitAnalysisJobMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
}

type SubmitAnalysisJobTemplateList []SubmitAnalysisJobTemplate

func (list *SubmitAnalysisJobTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAnalysisJobTemplate)
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
