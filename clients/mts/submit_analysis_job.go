package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAnalysisJobRequest struct {
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

func (r SubmitAnalysisJobRequest) Invoke(client *sdk.Client) (response *SubmitAnalysisJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAnalysisJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitAnalysisJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAnalysisJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAnalysisJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAnalysisJobResponse struct {
	RequestId   string
	AnalysisJob SubmitAnalysisJobAnalysisJob
}

type SubmitAnalysisJobAnalysisJob struct {
	Id               string
	UserData         string
	State            string
	Code             string
	Message          string
	Percent          int64
	CreationTime     string
	PipelineId       string
	Priority         string
	TemplateList     SubmitAnalysisJobTemplateList
	InputFile        SubmitAnalysisJobInputFile
	AnalysisConfig   SubmitAnalysisJobAnalysisConfig
	MNSMessageResult SubmitAnalysisJobMNSMessageResult
}

type SubmitAnalysisJobTemplate struct {
	Id          string
	Name        string
	State       string
	Container   SubmitAnalysisJobContainer
	Video       SubmitAnalysisJobVideo
	Audio       SubmitAnalysisJobAudio
	TransConfig SubmitAnalysisJobTransConfig
	MuxConfig   SubmitAnalysisJobMuxConfig
}

type SubmitAnalysisJobContainer struct {
	Format string
}

type SubmitAnalysisJobVideo struct {
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
	BitrateBnd SubmitAnalysisJobBitrateBnd
}

type SubmitAnalysisJobBitrateBnd struct {
	Max string
	Min string
}

type SubmitAnalysisJobAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
}

type SubmitAnalysisJobTransConfig struct {
	TransMode string
}

type SubmitAnalysisJobMuxConfig struct {
	Segment SubmitAnalysisJobSegment
	Gif     SubmitAnalysisJobGif
}

type SubmitAnalysisJobSegment struct {
	Duration string
}

type SubmitAnalysisJobGif struct {
	Loop       string
	FinalDelay string
}

type SubmitAnalysisJobInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitAnalysisJobAnalysisConfig struct {
	QualityControl    SubmitAnalysisJobQualityControl
	PropertiesControl SubmitAnalysisJobPropertiesControl
}

type SubmitAnalysisJobQualityControl struct {
	RateQuality     string
	MethodStreaming string
}

type SubmitAnalysisJobPropertiesControl struct {
	Deinterlace string
	Crop        SubmitAnalysisJobCrop
}

type SubmitAnalysisJobCrop struct {
	Mode   string
	Width  string
	Height string
	Top    string
	Left   string
}

type SubmitAnalysisJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
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
