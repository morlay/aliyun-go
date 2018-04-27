package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterEpisodeRequest struct {
	requests.RpcRequest
	ResourceId   string                              `position:"Query" name:"ResourceId"`
	ComponentIds *ModifyCasterEpisodeComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SwitchType   string                              `position:"Query" name:"SwitchType"`
	CasterId     string                              `position:"Query" name:"CasterId"`
	EpisodeName  string                              `position:"Query" name:"EpisodeName"`
	EndTime      string                              `position:"Query" name:"EndTime"`
	StartTime    string                              `position:"Query" name:"StartTime"`
	OwnerId      int64                               `position:"Query" name:"OwnerId"`
	EpisodeId    string                              `position:"Query" name:"EpisodeId"`
}

func (req *ModifyCasterEpisodeRequest) Invoke(client *sdk.Client) (resp *ModifyCasterEpisodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterEpisode", "live", "")
	resp = &ModifyCasterEpisodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCasterEpisodeResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
	EpisodeId string
}

type ModifyCasterEpisodeComponentIdList []string

func (list *ModifyCasterEpisodeComponentIdList) UnmarshalJSON(data []byte) error {
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
