package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMediaInfoJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaInfoJobIds      string `position:"Query" name:"MediaInfoJobIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryMediaInfoJobListRequest) Invoke(client *sdk.Client) (resp *QueryMediaInfoJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaInfoJobList", "mts", "")
	resp = &QueryMediaInfoJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaInfoJobListResponse struct {
	responses.BaseResponse
	RequestId               common.String
	MediaInfoJobList        QueryMediaInfoJobListMediaInfoJobList
	NonExistMediaInfoJobIds QueryMediaInfoJobListNonExistMediaInfoJobIdList
}

type QueryMediaInfoJobListMediaInfoJob struct {
	JobId            common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Message          common.String
	CreationTime     common.String
	Input            QueryMediaInfoJobListInput
	Properties       QueryMediaInfoJobListProperties
	MNSMessageResult QueryMediaInfoJobListMNSMessageResult
}

type QueryMediaInfoJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryMediaInfoJobListProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    QueryMediaInfoJobListStreams
	Format     QueryMediaInfoJobListFormat
}

type QueryMediaInfoJobListStreams struct {
	VideoStreamList    QueryMediaInfoJobListVideoStreamList
	AudioStreamList    QueryMediaInfoJobListAudioStreamList
	SubtitleStreamList QueryMediaInfoJobListSubtitleStreamList
}

type QueryMediaInfoJobListVideoStream struct {
	Index          common.String
	CodecName      common.String
	CodecLongName  common.String
	Profile        common.String
	CodecTimeBase  common.String
	CodecTagString common.String
	CodecTag       common.String
	Width          common.String
	Height         common.String
	HasBFrames     common.String
	Sar            common.String
	Dar            common.String
	PixFmt         common.String
	Level          common.String
	Fps            common.String
	AvgFPS         common.String
	Timebase       common.String
	StartTime      common.String
	Duration       common.String
	Bitrate        common.String
	NumFrames      common.String
	Lang           common.String
	Rotate         common.String
	NetworkCost    QueryMediaInfoJobListNetworkCost
}

type QueryMediaInfoJobListNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type QueryMediaInfoJobListAudioStream struct {
	Index          common.String
	CodecName      common.String
	CodecTimeBase  common.String
	CodecLongName  common.String
	CodecTagString common.String
	CodecTag       common.String
	SampleFmt      common.String
	Samplerate     common.String
	Channels       common.String
	ChannelLayout  common.String
	Timebase       common.String
	StartTime      common.String
	Duration       common.String
	Bitrate        common.String
	NumFrames      common.String
	Lang           common.String
}

type QueryMediaInfoJobListSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type QueryMediaInfoJobListFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type QueryMediaInfoJobListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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

type QueryMediaInfoJobListNonExistMediaInfoJobIdList []common.String

func (list *QueryMediaInfoJobListNonExistMediaInfoJobIdList) UnmarshalJSON(data []byte) error {
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
