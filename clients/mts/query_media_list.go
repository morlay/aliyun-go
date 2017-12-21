package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
}

func (r QueryMediaListRequest) Invoke(client *sdk.Client) (response *QueryMediaListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryMediaListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryMediaListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryMediaListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryMediaListResponse struct {
	RequestId        string
	MediaList        QueryMediaListMediaList
	NonExistMediaIds QueryMediaListNonExistMediaIdList
}

type QueryMediaListMedia struct {
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
	PlayList     QueryMediaListPlayList
	SnapshotList QueryMediaListSnapshotList
	Tags         QueryMediaListTagList
	RunIdList    QueryMediaListRunIdListList
	File         QueryMediaListFile
	MediaInfo    QueryMediaListMediaInfo
}

type QueryMediaListPlay struct {
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
	File1             QueryMediaListFile1
}

type QueryMediaListFile1 struct {
	URL   string
	State string
}

type QueryMediaListSnapshot struct {
	Type              string
	MediaWorkflowId   string
	MediaWorkflowName string
	ActivityName      string
	Count             string
	File2             QueryMediaListFile2
}

type QueryMediaListFile2 struct {
	URL   string
	State string
}

type QueryMediaListFile struct {
	URL   string
	State string
}

type QueryMediaListMediaInfo struct {
	Streams QueryMediaListStreams
	Format  QueryMediaListFormat
}

type QueryMediaListStreams struct {
	VideoStreamList    QueryMediaListVideoStreamList
	AudioStreamList    QueryMediaListAudioStreamList
	SubtitleStreamList QueryMediaListSubtitleStreamList
}

type QueryMediaListVideoStream struct {
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
	NetworkCost    QueryMediaListNetworkCost
}

type QueryMediaListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaListAudioStream struct {
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

type QueryMediaListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryMediaListMediaList []QueryMediaListMedia

func (list *QueryMediaListMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListMedia)
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

type QueryMediaListNonExistMediaIdList []string

func (list *QueryMediaListNonExistMediaIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListPlayList []QueryMediaListPlay

func (list *QueryMediaListPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListPlay)
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

type QueryMediaListSnapshotList []QueryMediaListSnapshot

func (list *QueryMediaListSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSnapshot)
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

type QueryMediaListTagList []string

func (list *QueryMediaListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListRunIdListList []string

func (list *QueryMediaListRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListVideoStreamList []QueryMediaListVideoStream

func (list *QueryMediaListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListVideoStream)
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

type QueryMediaListAudioStreamList []QueryMediaListAudioStream

func (list *QueryMediaListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListAudioStream)
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

type QueryMediaListSubtitleStreamList []QueryMediaListSubtitleStream

func (list *QueryMediaListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSubtitleStream)
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
