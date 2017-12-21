package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaListByURLRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	FileURLs             string `position:"Query" name:"FileURLs"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
}

func (req *QueryMediaListByURLRequest) Invoke(client *sdk.Client) (resp *QueryMediaListByURLResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaListByURL", "", "")
	resp = &QueryMediaListByURLResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaListByURLResponse struct {
	responses.BaseResponse
	RequestId        string
	MediaList        QueryMediaListByURLMediaList
	NonExistFileURLs QueryMediaListByURLNonExistFileURLList
}

type QueryMediaListByURLMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	PlayList     QueryMediaListByURLPlayList
	SnapshotList QueryMediaListByURLSnapshotList
	Tags         QueryMediaListByURLTagList
	RunIdList    QueryMediaListByURLRunIdListList
	File         QueryMediaListByURLFile
	MediaInfo    QueryMediaListByURLMediaInfo
}

type QueryMediaListByURLPlay struct {
	ActivityName      string
	MediaWorkflowId   string
	MediaWorkflowName string
	Duration          string
	Format            string
	Size              string
	Bitrate           string
	Width             string
	Height            string
	Fps               string
	Encryption        string
	File1             QueryMediaListByURLFile1
}

type QueryMediaListByURLFile1 struct {
	URL   string
	State string
}

type QueryMediaListByURLSnapshot struct {
	Type              string
	MediaWorkflowId   string
	MediaWorkflowName string
	ActivityName      string
	Count             string
	File2             QueryMediaListByURLFile2
}

type QueryMediaListByURLFile2 struct {
	URL   string
	State string
}

type QueryMediaListByURLFile struct {
	URL   string
	State string
}

type QueryMediaListByURLMediaInfo struct {
	Streams QueryMediaListByURLStreams
	Format  QueryMediaListByURLFormat
}

type QueryMediaListByURLStreams struct {
	VideoStreamList    QueryMediaListByURLVideoStreamList
	AudioStreamList    QueryMediaListByURLAudioStreamList
	SubtitleStreamList QueryMediaListByURLSubtitleStreamList
}

type QueryMediaListByURLVideoStream struct {
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
	NetworkCost    QueryMediaListByURLNetworkCost
}

type QueryMediaListByURLNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaListByURLAudioStream struct {
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

type QueryMediaListByURLSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaListByURLFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryMediaListByURLMediaList []QueryMediaListByURLMedia

func (list *QueryMediaListByURLMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLMedia)
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

type QueryMediaListByURLNonExistFileURLList []string

func (list *QueryMediaListByURLNonExistFileURLList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLPlayList []QueryMediaListByURLPlay

func (list *QueryMediaListByURLPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLPlay)
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

type QueryMediaListByURLSnapshotList []QueryMediaListByURLSnapshot

func (list *QueryMediaListByURLSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSnapshot)
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

type QueryMediaListByURLTagList []string

func (list *QueryMediaListByURLTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLRunIdListList []string

func (list *QueryMediaListByURLRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLVideoStreamList []QueryMediaListByURLVideoStream

func (list *QueryMediaListByURLVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLVideoStream)
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

type QueryMediaListByURLAudioStreamList []QueryMediaListByURLAudioStream

func (list *QueryMediaListByURLAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLAudioStream)
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

type QueryMediaListByURLSubtitleStreamList []QueryMediaListByURLSubtitleStream

func (list *QueryMediaListByURLSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSubtitleStream)
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
