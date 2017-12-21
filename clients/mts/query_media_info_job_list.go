package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaInfoJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaInfoJobIds      string `position:"Query" name:"MediaInfoJobIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryMediaInfoJobListRequest) Invoke(client *sdk.Client) (response *QueryMediaInfoJobListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryMediaInfoJobListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaInfoJobList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryMediaInfoJobListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryMediaInfoJobListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryMediaInfoJobListResponse struct {
	RequestId               string
	MediaInfoJobList        QueryMediaInfoJobListMediaInfoJobList
	NonExistMediaInfoJobIds QueryMediaInfoJobListNonExistMediaInfoJobIdList
}

type QueryMediaInfoJobListMediaInfoJob struct {
	JobId            string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            QueryMediaInfoJobListInput
	Properties       QueryMediaInfoJobListProperties
	MNSMessageResult QueryMediaInfoJobListMNSMessageResult
}

type QueryMediaInfoJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryMediaInfoJobListProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    QueryMediaInfoJobListStreams
	Format     QueryMediaInfoJobListFormat
}

type QueryMediaInfoJobListStreams struct {
	VideoStreamList    QueryMediaInfoJobListVideoStreamList
	AudioStreamList    QueryMediaInfoJobListAudioStreamList
	SubtitleStreamList QueryMediaInfoJobListSubtitleStreamList
}

type QueryMediaInfoJobListVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	Rotate         string
	NetworkCost    QueryMediaInfoJobListNetworkCost
}

type QueryMediaInfoJobListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaInfoJobListAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryMediaInfoJobListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaInfoJobListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryMediaInfoJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type QueryMediaInfoJobListMediaInfoJobList []QueryMediaInfoJobListMediaInfoJob

func (list *QueryMediaInfoJobListMediaInfoJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListMediaInfoJob)
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

type QueryMediaInfoJobListNonExistMediaInfoJobIdList []string

func (list *QueryMediaInfoJobListNonExistMediaInfoJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaInfoJobListVideoStreamList []QueryMediaInfoJobListVideoStream

func (list *QueryMediaInfoJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListVideoStream)
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

type QueryMediaInfoJobListAudioStreamList []QueryMediaInfoJobListAudioStream

func (list *QueryMediaInfoJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListAudioStream)
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

type QueryMediaInfoJobListSubtitleStreamList []QueryMediaInfoJobListSubtitleStream

func (list *QueryMediaInfoJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListSubtitleStream)
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
