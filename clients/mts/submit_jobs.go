package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitJobsRequest struct {
	requests.RpcRequest
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

func (req *SubmitJobsRequest) Invoke(client *sdk.Client) (resp *SubmitJobsResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitJobs", "mts", "")
	resp = &SubmitJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitJobsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	JobResultList SubmitJobsJobResultList
}

type SubmitJobsJobResult struct {
	Success bool
	Code    common.String
	Message common.String
	Job     SubmitJobsJob
}

type SubmitJobsJob struct {
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	PipelineId       common.String
	CreationTime     common.String
	FinishTime       common.String
	Input            SubmitJobsInput
	Output           SubmitJobsOutput
	MNSMessageResult SubmitJobsMNSMessageResult
}

type SubmitJobsInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitJobsOutput struct {
	TemplateId             common.String
	UserData               common.String
	Rotate                 common.String
	VideoStreamMap         common.String
	AudioStreamMap         common.String
	DeWatermark            common.String
	Priority               common.String
	WaterMarkConfigUrl     common.String
	MergeConfigUrl         common.String
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
	WaterMarkTemplateId common.String
	Width               common.String
	Height              common.String
	Dx                  common.String
	Dy                  common.String
	ReferPos            common.String
	Type                common.String
	InputFile           SubmitJobsInputFile
}

type SubmitJobsInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitJobsMerge struct {
	MergeURL common.String
	Start    common.String
	Duration common.String
	RoleArn  common.String
}

type SubmitJobsOpening struct {
	OpenUrl common.String
	Start   common.String
	Width   common.String
	Height  common.String
}

type SubmitJobsTailSlate struct {
	TailUrl       common.String
	Start         common.String
	BlendDuration common.String
	Width         common.String
	Height        common.String
	IsMergeAudio  bool
	BgColor       common.String
}

type SubmitJobsDigiWaterMark struct {
	Type       common.String
	Alpha      common.String
	InputFile1 SubmitJobsInputFile1
}

type SubmitJobsInputFile1 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitJobsOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type SubmitJobsM3U8NonStandardSupport struct {
	TS SubmitJobsTS
}

type SubmitJobsTS struct {
	Md5Support  bool
	SizeSupport bool
}

type SubmitJobsProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    SubmitJobsStreams
	Format     SubmitJobsFormat
}

type SubmitJobsStreams struct {
	VideoStreamList    SubmitJobsVideoStreamList
	AudioStreamList    SubmitJobsAudioStreamList
	SubtitleStreamList SubmitJobsSubtitleStreamList
}

type SubmitJobsVideoStream struct {
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
	NetworkCost    SubmitJobsNetworkCost
}

type SubmitJobsNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type SubmitJobsAudioStream struct {
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

type SubmitJobsSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type SubmitJobsFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type SubmitJobsClip struct {
	TimeSpan SubmitJobsTimeSpan
}

type SubmitJobsTimeSpan struct {
	Seek     common.String
	Duration common.String
}

type SubmitJobsSuperReso struct {
	IsHalfSample common.String
}

type SubmitJobsSubtitleConfig struct {
	SubtitleList    SubmitJobsSubtitleList
	ExtSubtitleList SubmitJobsExtSubtitleList
}

type SubmitJobsSubtitle struct {
	Map common.String
}

type SubmitJobsExtSubtitle struct {
	FontName common.String
	CharEnc  common.String
	Input2   SubmitJobsInput2
}

type SubmitJobsInput2 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitJobsTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type SubmitJobsMuxConfig struct {
	Segment SubmitJobsSegment
	Gif     SubmitJobsGif
}

type SubmitJobsSegment struct {
	Duration common.String
}

type SubmitJobsGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type SubmitJobsAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Volume     SubmitJobsVolume
}

type SubmitJobsVolume struct {
	Level  common.String
	Method common.String
}

type SubmitJobsVideo struct {
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
	Crop       common.String
	Pad        common.String
	MaxFps     common.String
	BitrateBnd SubmitJobsBitrateBnd
}

type SubmitJobsBitrateBnd struct {
	Max common.String
	Min common.String
}

type SubmitJobsContainer struct {
	Format common.String
}

type SubmitJobsEncryption struct {
	Type    common.String
	Id      common.String
	Key     common.String
	KeyUri  common.String
	KeyType common.String
	SkipCnt common.String
}

type SubmitJobsMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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
