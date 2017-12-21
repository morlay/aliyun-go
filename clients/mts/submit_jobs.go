package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitJobsRequest struct {
	Outputs              string `position:"Query" name:"Outputs"`
	Input                string `position:"Query" name:"Input"`
	OutputBucket         string `position:"Query" name:"OutputBucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OutputLocation       string `position:"Query" name:"OutputLocation"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r SubmitJobsRequest) Invoke(client *sdk.Client) (response *SubmitJobsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitJobsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitJobs", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitJobsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitJobsResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitJobsResponse struct {
	RequestId     string
	JobResultList SubmitJobsJobResultList
}

type SubmitJobsJobResult struct {
	Success bool
	Code    string
	Message string
	Job     SubmitJobsJob
}

type SubmitJobsJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            SubmitJobsInput
	Output           SubmitJobsOutput
	MNSMessageResult SubmitJobsMNSMessageResult
}

type SubmitJobsInput struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          SubmitJobsWaterMarkList
	MergeList              SubmitJobsMergeList
	OpeningList            SubmitJobsOpeningList
	TailSlateList          SubmitJobsTailSlateList
	DigiWaterMark          SubmitJobsDigiWaterMark
	OutputFile             SubmitJobsOutputFile
	M3U8NonStandardSupport SubmitJobsM3U8NonStandardSupport
	Properties             SubmitJobsProperties
	Clip                   SubmitJobsClip
	SuperReso              SubmitJobsSuperReso
	SubtitleConfig         SubmitJobsSubtitleConfig
	TransConfig            SubmitJobsTransConfig
	MuxConfig              SubmitJobsMuxConfig
	Audio                  SubmitJobsAudio
	Video                  SubmitJobsVideo
	Container              SubmitJobsContainer
	Encryption             SubmitJobsEncryption
}

type SubmitJobsWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           SubmitJobsInputFile
}

type SubmitJobsInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type SubmitJobsOpening struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type SubmitJobsTailSlate struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  bool
	BgColor       string
}

type SubmitJobsDigiWaterMark struct {
	Type       string
	Alpha      string
	InputFile1 SubmitJobsInputFile1
}

type SubmitJobsInputFile1 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitJobsM3U8NonStandardSupport struct {
	TS SubmitJobsTS
}

type SubmitJobsTS struct {
	Md5Support  bool
	SizeSupport bool
}

type SubmitJobsProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    SubmitJobsStreams
	Format     SubmitJobsFormat
}

type SubmitJobsStreams struct {
	VideoStreamList    SubmitJobsVideoStreamList
	AudioStreamList    SubmitJobsAudioStreamList
	SubtitleStreamList SubmitJobsSubtitleStreamList
}

type SubmitJobsVideoStream struct {
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
	NetworkCost    SubmitJobsNetworkCost
}

type SubmitJobsNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type SubmitJobsAudioStream struct {
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

type SubmitJobsSubtitleStream struct {
	Index string
	Lang  string
}

type SubmitJobsFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type SubmitJobsClip struct {
	TimeSpan SubmitJobsTimeSpan
}

type SubmitJobsTimeSpan struct {
	Seek     string
	Duration string
}

type SubmitJobsSuperReso struct {
	IsHalfSample string
}

type SubmitJobsSubtitleConfig struct {
	SubtitleList    SubmitJobsSubtitleList
	ExtSubtitleList SubmitJobsExtSubtitleList
}

type SubmitJobsSubtitle struct {
	Map string
}

type SubmitJobsExtSubtitle struct {
	FontName string
	CharEnc  string
	Input2   SubmitJobsInput2
}

type SubmitJobsInput2 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type SubmitJobsMuxConfig struct {
	Segment SubmitJobsSegment
	Gif     SubmitJobsGif
}

type SubmitJobsSegment struct {
	Duration string
}

type SubmitJobsGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type SubmitJobsAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     SubmitJobsVolume
}

type SubmitJobsVolume struct {
	Level  string
	Method string
}

type SubmitJobsVideo struct {
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
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd SubmitJobsBitrateBnd
}

type SubmitJobsBitrateBnd struct {
	Max string
	Min string
}

type SubmitJobsContainer struct {
	Format string
}

type SubmitJobsEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type SubmitJobsMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type SubmitJobsJobResultList []SubmitJobsJobResult

func (list *SubmitJobsJobResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsJobResult)
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

type SubmitJobsWaterMarkList []SubmitJobsWaterMark

func (list *SubmitJobsWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsWaterMark)
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

type SubmitJobsMergeList []SubmitJobsMerge

func (list *SubmitJobsMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsMerge)
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

type SubmitJobsOpeningList []SubmitJobsOpening

func (list *SubmitJobsOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsOpening)
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

type SubmitJobsTailSlateList []SubmitJobsTailSlate

func (list *SubmitJobsTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsTailSlate)
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

type SubmitJobsVideoStreamList []SubmitJobsVideoStream

func (list *SubmitJobsVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsVideoStream)
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

type SubmitJobsAudioStreamList []SubmitJobsAudioStream

func (list *SubmitJobsAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsAudioStream)
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

type SubmitJobsSubtitleStreamList []SubmitJobsSubtitleStream

func (list *SubmitJobsSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitleStream)
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

type SubmitJobsSubtitleList []SubmitJobsSubtitle

func (list *SubmitJobsSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitle)
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

type SubmitJobsExtSubtitleList []SubmitJobsExtSubtitle

func (list *SubmitJobsExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsExtSubtitle)
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
