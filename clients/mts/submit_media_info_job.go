package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitMediaInfoJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitMediaInfoJobRequest) Invoke(client *sdk.Client) (resp *SubmitMediaInfoJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitMediaInfoJob", "mts", "")
	resp = &SubmitMediaInfoJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitMediaInfoJobResponse struct {
	responses.BaseResponse
	RequestId    string
	MediaInfoJob SubmitMediaInfoJobMediaInfoJob
}

type SubmitMediaInfoJobMediaInfoJob struct {
	JobId            string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            SubmitMediaInfoJobInput
	Properties       SubmitMediaInfoJobProperties
	MNSMessageResult SubmitMediaInfoJobMNSMessageResult
}

type SubmitMediaInfoJobInput struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitMediaInfoJobProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    SubmitMediaInfoJobStreams
	Format     SubmitMediaInfoJobFormat
}

type SubmitMediaInfoJobStreams struct {
	VideoStreamList    SubmitMediaInfoJobVideoStreamList
	AudioStreamList    SubmitMediaInfoJobAudioStreamList
	SubtitleStreamList SubmitMediaInfoJobSubtitleStreamList
}

type SubmitMediaInfoJobVideoStream struct {
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
	NetworkCost    SubmitMediaInfoJobNetworkCost
}

type SubmitMediaInfoJobNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type SubmitMediaInfoJobAudioStream struct {
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

type SubmitMediaInfoJobSubtitleStream struct {
	Index string
	Lang  string
}

type SubmitMediaInfoJobFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type SubmitMediaInfoJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type SubmitMediaInfoJobVideoStreamList []SubmitMediaInfoJobVideoStream

func (list *SubmitMediaInfoJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobVideoStream)
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

type SubmitMediaInfoJobAudioStreamList []SubmitMediaInfoJobAudioStream

func (list *SubmitMediaInfoJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobAudioStream)
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

type SubmitMediaInfoJobSubtitleStreamList []SubmitMediaInfoJobSubtitleStream

func (list *SubmitMediaInfoJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobSubtitleStream)
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
