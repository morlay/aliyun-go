package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId              common.String
	AnalysisJobList        QueryAnalysisJobListAnalysisJobList
	NonExistAnalysisJobIds QueryAnalysisJobListNonExistAnalysisJobIdList
}

type QueryAnalysisJobListAnalysisJob struct {
	Id               common.String
	UserData         common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	CreationTime     common.String
	PipelineId       common.String
	Priority         common.String
	TemplateList     QueryAnalysisJobListTemplateList
	InputFile        QueryAnalysisJobListInputFile
	AnalysisConfig   QueryAnalysisJobListAnalysisConfig
	MNSMessageResult QueryAnalysisJobListMNSMessageResult
}

type QueryAnalysisJobListTemplate struct {
	Id          common.String
	Name        common.String
	State       common.String
	Container   QueryAnalysisJobListContainer
	Video       QueryAnalysisJobListVideo
	Audio       QueryAnalysisJobListAudio
	TransConfig QueryAnalysisJobListTransConfig
	MuxConfig   QueryAnalysisJobListMuxConfig
}

type QueryAnalysisJobListContainer struct {
	Format common.String
}

type QueryAnalysisJobListVideo struct {
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
	BitrateBnd QueryAnalysisJobListBitrateBnd
}

type QueryAnalysisJobListBitrateBnd struct {
	Max common.String
	Min common.String
}

type QueryAnalysisJobListAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
}

type QueryAnalysisJobListTransConfig struct {
	TransMode common.String
}

type QueryAnalysisJobListMuxConfig struct {
	Segment QueryAnalysisJobListSegment
	Gif     QueryAnalysisJobListGif
}

type QueryAnalysisJobListSegment struct {
	Duration common.String
}

type QueryAnalysisJobListGif struct {
	Loop       common.String
	FinalDelay common.String
}

type QueryAnalysisJobListInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryAnalysisJobListAnalysisConfig struct {
	QualityControl    QueryAnalysisJobListQualityControl
	PropertiesControl QueryAnalysisJobListPropertiesControl
}

type QueryAnalysisJobListQualityControl struct {
	RateQuality     common.String
	MethodStreaming common.String
}

type QueryAnalysisJobListPropertiesControl struct {
	Deinterlace common.String
	Crop        QueryAnalysisJobListCrop
}

type QueryAnalysisJobListCrop struct {
	Mode   common.String
	Width  common.String
	Height common.String
	Top    common.String
	Left   common.String
}

type QueryAnalysisJobListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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

type QueryAnalysisJobListNonExistAnalysisJobIdList []common.String

func (list *QueryAnalysisJobListNonExistAnalysisJobIdList) UnmarshalJSON(data []byte) error {
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
