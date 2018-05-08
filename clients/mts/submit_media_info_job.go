package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	MediaInfoJob SubmitMediaInfoJobMediaInfoJob
}

type SubmitMediaInfoJobMediaInfoJob struct {
	JobId            common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Message          common.String
	CreationTime     common.String
	Input            SubmitMediaInfoJobInput
	Properties       SubmitMediaInfoJobProperties
	MNSMessageResult SubmitMediaInfoJobMNSMessageResult
}

type SubmitMediaInfoJobInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitMediaInfoJobProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    SubmitMediaInfoJobStreams
	Format     SubmitMediaInfoJobFormat
}

type SubmitMediaInfoJobStreams struct {
	VideoStreamList    SubmitMediaInfoJobVideoStreamList
	AudioStreamList    SubmitMediaInfoJobAudioStreamList
	SubtitleStreamList SubmitMediaInfoJobSubtitleStreamList
}

type SubmitMediaInfoJobVideoStream struct {
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
	NetworkCost    SubmitMediaInfoJobNetworkCost
}

type SubmitMediaInfoJobNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type SubmitMediaInfoJobAudioStream struct {
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

type SubmitMediaInfoJobSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type SubmitMediaInfoJobFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type SubmitMediaInfoJobMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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
