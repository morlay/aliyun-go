package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterProgramRequest struct {
	requests.RpcRequest
	CasterId    string `position:"Query" name:"CasterId"`
	EpisodeType string `position:"Query" name:"EpisodeType"`
	PageSize    int    `position:"Query" name:"PageSize"`
	EndTime     string `position:"Query" name:"EndTime"`
	StartTime   string `position:"Query" name:"StartTime"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	EpisodeId   string `position:"Query" name:"EpisodeId"`
	PageNum     int    `position:"Query" name:"PageNum"`
	Status      int    `position:"Query" name:"Status"`
}

func (req *DescribeCasterProgramRequest) Invoke(client *sdk.Client) (resp *DescribeCasterProgramResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterProgram", "live", "")
	resp = &DescribeCasterProgramResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterProgramResponse struct {
	responses.BaseResponse
	RequestId     string
	CasterId      string
	ProgramName   string
	ProgramEffect int
	Total         int
	Episodes      DescribeCasterProgramEpisodeList
}

type DescribeCasterProgramEpisode struct {
	EpisodeId    string
	EpisodeType  string
	EpisodeName  string
	ResourceId   string
	StartTime    string
	EndTime      string
	SwitchType   string
	Status       int
	ComponentIds DescribeCasterProgramComponentIdList
}

type DescribeCasterProgramEpisodeList []DescribeCasterProgramEpisode

func (list *DescribeCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterProgramEpisode)
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

type DescribeCasterProgramComponentIdList []string

func (list *DescribeCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
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
