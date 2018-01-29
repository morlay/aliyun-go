package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken              string `position:"Query" name:"NextPageToken"`
	StartOfJobCreatedTimeRange string `position:"Query" name:"StartOfJobCreatedTimeRange"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            int64  `position:"Query" name:"MaximumPageSize"`
	State                      string `position:"Query" name:"State"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	EndOfJobCreatedTimeRange   string `position:"Query" name:"EndOfJobCreatedTimeRange"`
	PipelineId                 string `position:"Query" name:"PipelineId"`
}

func (req *ListJobRequest) Invoke(client *sdk.Client) (resp *ListJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListJob", "mts", "")
	resp = &ListJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobResponse struct {
	responses.BaseResponse
	RequestId     string
	NextPageToken string
	JobList       ListJobJobList
}

type ListJobJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            ListJobInput
	Output           ListJobOutput
	MNSMessageResult ListJobMNSMessageResult
}

type ListJobInput struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          ListJobWaterMarkList
	MergeList              ListJobMergeList
	OpeningList            ListJobMerge1List
	TailSlateList          ListJobMerge2List
	OutputFile             ListJobOutputFile
	M3U8NonStandardSupport ListJobM3U8NonStandardSupport
	Properties             ListJobProperties
	Clip                   ListJobClip
	SuperReso              ListJobSuperReso
	SubtitleConfig         ListJobSubtitleConfig
	TransConfig            ListJobTransConfig
	MuxConfig              ListJobMuxConfig
	Audio                  ListJobAudio
	Video                  ListJobVideo
	Container              ListJobContainer
	Encryption             ListJobEncryption
}

type ListJobWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           ListJobInputFile
}

type ListJobInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type ListJobMerge1 struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type ListJobMerge2 struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  bool
	BgColor       string
}

type ListJobOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type ListJobM3U8NonStandardSupport struct {
	TS ListJobTS
}

type ListJobTS struct {
	Md5Support  bool
	SizeSupport bool
}

type ListJobProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    ListJobStreams
	Format     ListJobFormat
}

type ListJobStreams struct {
	VideoStreamList    ListJobVideoStreamList
	AudioStreamList    ListJobAudioStreamList
	SubtitleStreamList ListJobSubtitleStreamList
}

type ListJobVideoStream struct {
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
	NetworkCost    ListJobNetworkCost
}

type ListJobNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type ListJobAudioStream struct {
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

type ListJobSubtitleStream struct {
	Index string
	Lang  string
}

type ListJobFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type ListJobClip struct {
	TimeSpan ListJobTimeSpan
}

type ListJobTimeSpan struct {
	Seek     string
	Duration string
}

type ListJobSuperReso struct {
	IsHalfSample string
}

type ListJobSubtitleConfig struct {
	SubtitleList    ListJobSubtitleList
	ExtSubtitleList ListJobExtSubtitleList
}

type ListJobSubtitle struct {
	Map string
}

type ListJobExtSubtitle struct {
	FontName string
	CharEnc  string
	Input3   ListJobInput3
}

type ListJobInput3 struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type ListJobMuxConfig struct {
	Segment ListJobSegment
	Gif     ListJobGif
}

type ListJobSegment struct {
	Duration string
}

type ListJobGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type ListJobAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     ListJobVolume
}

type ListJobVolume struct {
	Level  string
	Method string
}

type ListJobVideo struct {
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
	BitrateBnd ListJobBitrateBnd
}

type ListJobBitrateBnd struct {
	Max string
	Min string
}

type ListJobContainer struct {
	Format string
}

type ListJobEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type ListJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type ListJobJobList []ListJobJob

func (list *ListJobJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobJob)
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

type ListJobWaterMarkList []ListJobWaterMark

func (list *ListJobWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobWaterMark)
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

type ListJobMergeList []ListJobMerge

func (list *ListJobMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge)
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

type ListJobMerge1List []ListJobMerge1

func (list *ListJobMerge1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge1)
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

type ListJobMerge2List []ListJobMerge2

func (list *ListJobMerge2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge2)
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

type ListJobVideoStreamList []ListJobVideoStream

func (list *ListJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobVideoStream)
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

type ListJobAudioStreamList []ListJobAudioStream

func (list *ListJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobAudioStream)
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

type ListJobSubtitleStreamList []ListJobSubtitleStream

func (list *ListJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitleStream)
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

type ListJobSubtitleList []ListJobSubtitle

func (list *ListJobSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitle)
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

type ListJobExtSubtitleList []ListJobExtSubtitle

func (list *ListJobExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExtSubtitle)
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
