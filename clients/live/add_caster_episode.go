package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCasterEpisodeRequest struct {
	requests.RpcRequest
	ResourceId   string                           `position:"Query" name:"ResourceId"`
	ComponentIds *AddCasterEpisodeComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SwitchType   string                           `position:"Query" name:"SwitchType"`
	CasterId     string                           `position:"Query" name:"CasterId"`
	EpisodeType  string                           `position:"Query" name:"EpisodeType"`
	EpisodeName  string                           `position:"Query" name:"EpisodeName"`
	EndTime      string                           `position:"Query" name:"EndTime"`
	StartTime    string                           `position:"Query" name:"StartTime"`
	OwnerId      int64                            `position:"Query" name:"OwnerId"`
}

func (req *AddCasterEpisodeRequest) Invoke(client *sdk.Client) (resp *AddCasterEpisodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterEpisode", "live", "")
	resp = &AddCasterEpisodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterEpisodeResponse struct {
	responses.BaseResponse
	RequestId string
	EpisodeId string
}

type AddCasterEpisodeComponentIdList []string

func (list *AddCasterEpisodeComponentIdList) UnmarshalJSON(data []byte) error {
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
