package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	NextPageToken common.String
	JobList       ListJobJobList
}

type ListJobJob struct {
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	PipelineId       common.String
	CreationTime     common.String
	FinishTime       common.String
	Input            ListJobInput
	Output           ListJobOutput
	MNSMessageResult ListJobMNSMessageResult
}

type ListJobInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type ListJobOutput struct {
	TemplateId             common.String
	UserData               common.String
	Rotate                 common.String
	VideoStreamMap         common.String
	AudioStreamMap         common.String
	DeWatermark            common.String
	Priority               common.String
	WaterMarkConfigUrl     common.String
	MergeConfigUrl         common.String
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
	WaterMarkTemplateId common.String
	Width               common.String
	Height              common.String
	Dx                  common.String
	Dy                  common.String
	ReferPos            common.String
	Type                common.String
	InputFile           ListJobInputFile
}

type ListJobInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type ListJobMerge struct {
	MergeURL common.String
	Start    common.String
	Duration common.String
	RoleArn  common.String
}

type ListJobMerge1 struct {
	OpenUrl common.String
	Start   common.String
	Width   common.String
	Height  common.String
}

type ListJobMerge2 struct {
	TailUrl       common.String
	Start         common.String
	BlendDuration common.String
	Width         common.String
	Height        common.String
	IsMergeAudio  bool
	BgColor       common.String
}

type ListJobOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type ListJobM3U8NonStandardSupport struct {
	TS ListJobTS
}

type ListJobTS struct {
	Md5Support  bool
	SizeSupport bool
}

type ListJobProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    ListJobStreams
	Format     ListJobFormat
}

type ListJobStreams struct {
	VideoStreamList    ListJobVideoStreamList
	AudioStreamList    ListJobAudioStreamList
	SubtitleStreamList ListJobSubtitleStreamList
}

type ListJobVideoStream struct {
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
	NetworkCost    ListJobNetworkCost
}

type ListJobNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type ListJobAudioStream struct {
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

type ListJobSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type ListJobFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type ListJobClip struct {
	TimeSpan ListJobTimeSpan
}

type ListJobTimeSpan struct {
	Seek     common.String
	Duration common.String
}

type ListJobSuperReso struct {
	IsHalfSample common.String
}

type ListJobSubtitleConfig struct {
	SubtitleList    ListJobSubtitleList
	ExtSubtitleList ListJobExtSubtitleList
}

type ListJobSubtitle struct {
	Map common.String
}

type ListJobExtSubtitle struct {
	FontName common.String
	CharEnc  common.String
	Input3   ListJobInput3
}

type ListJobInput3 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type ListJobTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type ListJobMuxConfig struct {
	Segment ListJobSegment
	Gif     ListJobGif
}

type ListJobSegment struct {
	Duration common.String
}

type ListJobGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type ListJobAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Volume     ListJobVolume
}

type ListJobVolume struct {
	Level  common.String
	Method common.String
}

type ListJobVideo struct {
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
	BitrateBnd ListJobBitrateBnd
}

type ListJobBitrateBnd struct {
	Max common.String
	Min common.String
}

type ListJobContainer struct {
	Format common.String
}

type ListJobEncryption struct {
	Type    common.String
	Id      common.String
	Key     common.String
	KeyUri  common.String
	KeyType common.String
	SkipCnt common.String
}

type ListJobMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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
