package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterProgramRequest struct {
	requests.RpcRequest
	CasterId string                          `position:"Query" name:"CasterId"`
	Episodes *ModifyCasterProgramEpisodeList `position:"Query" type:"Repeated" name:"Episode"`
	OwnerId  int64                           `position:"Query" name:"OwnerId"`
}

func (req *ModifyCasterProgramRequest) Invoke(client *sdk.Client) (resp *ModifyCasterProgramResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterProgram", "live", "")
	resp = &ModifyCasterProgramResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCasterProgramEpisode struct {
	EpisodeId    string                              `name:"EpisodeId"`
	EpisodeType  string                              `name:"EpisodeType"`
	EpisodeName  string                              `name:"EpisodeName"`
	ResourceId   string                              `name:"ResourceId"`
	ComponentIds *ModifyCasterProgramComponentIdList `type:"Repeated" name:"ComponentId"`
	StartTime    string                              `name:"StartTime"`
	EndTime      string                              `name:"EndTime"`
	SwitchType   string                              `name:"SwitchType"`
}

type ModifyCasterProgramResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}

type ModifyCasterProgramEpisodeList []ModifyCasterProgramEpisode

func (list *ModifyCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterProgramEpisode)
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

type ModifyCasterProgramComponentIdList []string

func (list *ModifyCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
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
