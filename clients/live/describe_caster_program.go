package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	CasterId      common.String
	ProgramName   common.String
	ProgramEffect common.Integer
	Total         common.Integer
	Episodes      DescribeCasterProgramEpisodeList
}

type DescribeCasterProgramEpisode struct {
	EpisodeId    common.String
	EpisodeType  common.String
	EpisodeName  common.String
	ResourceId   common.String
	StartTime    common.String
	EndTime      common.String
	SwitchType   common.String
	Status       common.Integer
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

type DescribeCasterProgramComponentIdList []common.String

func (list *DescribeCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
