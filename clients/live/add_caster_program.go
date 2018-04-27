package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCasterProgramRequest struct {
	requests.RpcRequest
	CasterId string                       `position:"Query" name:"CasterId"`
	Episodes *AddCasterProgramEpisodeList `position:"Query" type:"Repeated" name:"Episode"`
	OwnerId  int64                        `position:"Query" name:"OwnerId"`
}

func (req *AddCasterProgramRequest) Invoke(client *sdk.Client) (resp *AddCasterProgramResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterProgram", "live", "")
	resp = &AddCasterProgramResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterProgramEpisode struct {
	EpisodeType  string                           `name:"EpisodeType"`
	EpisodeName  string                           `name:"EpisodeName"`
	ResourceId   string                           `name:"ResourceId"`
	ComponentIds *AddCasterProgramComponentIdList `type:"Repeated" name:"ComponentId"`
	StartTime    string                           `name:"StartTime"`
	EndTime      string                           `name:"EndTime"`
	SwitchType   string                           `name:"SwitchType"`
}

type AddCasterProgramResponse struct {
	responses.BaseResponse
	RequestId  string
	EpisodeIds AddCasterProgramEpisodeIdList
}

type AddCasterProgramEpisodeId struct {
	EpisodeId string
}

type AddCasterProgramEpisodeList []AddCasterProgramEpisode

func (list *AddCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterProgramEpisode)
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

type AddCasterProgramComponentIdList []string

func (list *AddCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
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

type AddCasterProgramEpisodeIdList []AddCasterProgramEpisodeId

func (list *AddCasterProgramEpisodeIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterProgramEpisodeId)
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
