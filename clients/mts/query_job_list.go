package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryJobListRequest) Invoke(client *sdk.Client) (resp *QueryJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryJobList", "mts", "")
	resp = &QueryJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryJobListResponse struct {
	responses.BaseResponse
	RequestId      string
	JobList        QueryJobListJobList
	NonExistJobIds QueryJobListNonExistJobIdList
}

type QueryJobListJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            QueryJobListInput
	Output           QueryJobListOutput
	MNSMessageResult QueryJobListMNSMessageResult
}

type QueryJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          QueryJobListWaterMarkList
	MergeList              QueryJobListMergeList
	OpeningList            QueryJobListOpeningList
	TailSlateList          QueryJobListTailSlateList
	OutputFile             QueryJobListOutputFile
	M3U8NonStandardSupport QueryJobListM3U8NonStandardSupport
	Properties             QueryJobListProperties
	Clip                   QueryJobListClip
	SuperReso              QueryJobListSuperReso
	SubtitleConfig         QueryJobListSubtitleConfig
	TransConfig            QueryJobListTransConfig
	MuxConfig              QueryJobListMuxConfig
	Audio                  QueryJobListAudio
	Video                  QueryJobListVideo
	Container              QueryJobListContainer
	Encryption             QueryJobListEncryption
}

type QueryJobListWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           QueryJobListInputFile
}

type QueryJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type QueryJobListOpening struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type QueryJobListTailSlate struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  bool
	BgColor       string
}

type QueryJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QueryJobListM3U8NonStandardSupport struct {
	TS QueryJobListTS
}

type QueryJobListTS struct {
	Md5Support  bool
	SizeSupport bool
}

type QueryJobListProperties struct {
	Width       string
	Height      string
	Bitrate     string
	Duration    string
	Fps         string
	FileSize    string
	FileFormat  string
	SourceLogos QueryJobListSourceLogoList
	Streams     QueryJobListStreams
	Format      QueryJobListFormat
}

type QueryJobListSourceLogo struct {
	Source string
}

type QueryJobListStreams struct {
	VideoStreamList    QueryJobListVideoStreamList
	AudioStreamList    QueryJobListAudioStreamList
	SubtitleStreamList QueryJobListSubtitleStreamList
}

type QueryJobListVideoStream struct {
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
	NetworkCost    QueryJobListNetworkCost
}

type QueryJobListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryJobListAudioStream struct {
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

type QueryJobListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryJobListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryJobListClip struct {
	TimeSpan QueryJobListTimeSpan
}

type QueryJobListTimeSpan struct {
	Seek     string
	Duration string
}

type QueryJobListSuperReso struct {
	IsHalfSample string
}

type QueryJobListSubtitleConfig struct {
	SubtitleList    QueryJobListSubtitleList
	ExtSubtitleList QueryJobListExtSubtitleList
}

type QueryJobListSubtitle struct {
	Map string
}

type QueryJobListExtSubtitle struct {
	FontName string
	CharEnc  string
	Input1   QueryJobListInput1
}

type QueryJobListInput1 struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type QueryJobListMuxConfig struct {
	Segment QueryJobListSegment
	Gif     QueryJobListGif
}

type QueryJobListSegment struct {
	Duration string
}

type QueryJobListGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type QueryJobListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     QueryJobListVolume
}

type QueryJobListVolume struct {
	Level  string
	Method string
}

type QueryJobListVideo struct {
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
	BitrateBnd QueryJobListBitrateBnd
}

type QueryJobListBitrateBnd struct {
	Max string
	Min string
}

type QueryJobListContainer struct {
	Format string
}

type QueryJobListEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type QueryJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type QueryJobListJobList []QueryJobListJob

func (list *QueryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListJob)
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

type QueryJobListNonExistJobIdList []string

func (list *QueryJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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

type QueryJobListWaterMarkList []QueryJobListWaterMark

func (list *QueryJobListWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListWaterMark)
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

type QueryJobListMergeList []QueryJobListMerge

func (list *QueryJobListMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListMerge)
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

type QueryJobListOpeningList []QueryJobListOpening

func (list *QueryJobListOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListOpening)
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

type QueryJobListTailSlateList []QueryJobListTailSlate

func (list *QueryJobListTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListTailSlate)
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

type QueryJobListSourceLogoList []QueryJobListSourceLogo

func (list *QueryJobListSourceLogoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSourceLogo)
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

type QueryJobListVideoStreamList []QueryJobListVideoStream

func (list *QueryJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListVideoStream)
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

type QueryJobListAudioStreamList []QueryJobListAudioStream

func (list *QueryJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListAudioStream)
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

type QueryJobListSubtitleStreamList []QueryJobListSubtitleStream

func (list *QueryJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitleStream)
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

type QueryJobListSubtitleList []QueryJobListSubtitle

func (list *QueryJobListSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitle)
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

type QueryJobListExtSubtitleList []QueryJobListExtSubtitle

func (list *QueryJobListExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListExtSubtitle)
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
