package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAnalysisJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AnalysisJobIds       string `position:"Query" name:"AnalysisJobIds"`
}

func (req *QueryAnalysisJobListRequest) Invoke(client *sdk.Client) (resp *QueryAnalysisJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAnalysisJobList", "mts", "")
	resp = &QueryAnalysisJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAnalysisJobListResponse struct {
	responses.BaseResponse
	RequestId              string
	AnalysisJobList        QueryAnalysisJobListAnalysisJobList
	NonExistAnalysisJobIds QueryAnalysisJobListNonExistAnalysisJobIdList
}

type QueryAnalysisJobListAnalysisJob struct {
	Id               string
	UserData         string
	State            string
	Code             string
	Message          string
	Percent          int64
	CreationTime     string
	PipelineId       string
	Priority         string
	TemplateList     QueryAnalysisJobListTemplateList
	InputFile        QueryAnalysisJobListInputFile
	AnalysisConfig   QueryAnalysisJobListAnalysisConfig
	MNSMessageResult QueryAnalysisJobListMNSMessageResult
}

type QueryAnalysisJobListTemplate struct {
	Id          string
	Name        string
	State       string
	Container   QueryAnalysisJobListContainer
	Video       QueryAnalysisJobListVideo
	Audio       QueryAnalysisJobListAudio
	TransConfig QueryAnalysisJobListTransConfig
	MuxConfig   QueryAnalysisJobListMuxConfig
}

type QueryAnalysisJobListContainer struct {
	Format string
}

type QueryAnalysisJobListVideo struct {
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
	BitrateBnd QueryAnalysisJobListBitrateBnd
}

type QueryAnalysisJobListBitrateBnd struct {
	Max string
	Min string
}

type QueryAnalysisJobListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
}

type QueryAnalysisJobListTransConfig struct {
	TransMode string
}

type QueryAnalysisJobListMuxConfig struct {
	Segment QueryAnalysisJobListSegment
	Gif     QueryAnalysisJobListGif
}

type QueryAnalysisJobListSegment struct {
	Duration string
}

type QueryAnalysisJobListGif struct {
	Loop       string
	FinalDelay string
}

type QueryAnalysisJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAnalysisJobListAnalysisConfig struct {
	QualityControl    QueryAnalysisJobListQualityControl
	PropertiesControl QueryAnalysisJobListPropertiesControl
}

type QueryAnalysisJobListQualityControl struct {
	RateQuality     string
	MethodStreaming string
}

type QueryAnalysisJobListPropertiesControl struct {
	Deinterlace string
	Crop        QueryAnalysisJobListCrop
}

type QueryAnalysisJobListCrop struct {
	Mode   string
	Width  string
	Height string
	Top    string
	Left   string
}

type QueryAnalysisJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type QueryAnalysisJobListAnalysisJobList []QueryAnalysisJobListAnalysisJob

func (list *QueryAnalysisJobListAnalysisJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListAnalysisJob)
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

type QueryAnalysisJobListNonExistAnalysisJobIdList []string

func (list *QueryAnalysisJobListNonExistAnalysisJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryAnalysisJobListTemplateList []QueryAnalysisJobListTemplate

func (list *QueryAnalysisJobListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListTemplate)
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
