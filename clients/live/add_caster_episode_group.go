package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCasterEpisodeGroupRequest struct {
	requests.RpcRequest
	SideOutputUrl string                         `position:"Query" name:"SideOutputUrl"`
	Items         *AddCasterEpisodeGroupItemList `position:"Query" type:"Repeated" name:"Item"`
	ClientToken   string                         `position:"Query" name:"ClientToken"`
	DomainName    string                         `position:"Query" name:"DomainName"`
	StartTime     string                         `position:"Query" name:"StartTime"`
	RepeatNum     int                            `position:"Query" name:"RepeatNum"`
	CallbackUrl   string                         `position:"Query" name:"CallbackUrl"`
	OwnerId       int64                          `position:"Query" name:"OwnerId"`
}

func (req *AddCasterEpisodeGroupRequest) Invoke(client *sdk.Client) (resp *AddCasterEpisodeGroupResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterEpisodeGroup", "live", "")
	resp = &AddCasterEpisodeGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCasterEpisodeGroupItem struct {
	ItemName string `name:"ItemName"`
	VodUrl   string `name:"VodUrl"`
}

type AddCasterEpisodeGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	ProgramId common.String
	ItemIds   AddCasterEpisodeGroupItemIdList
}

type AddCasterEpisodeGroupItemList []AddCasterEpisodeGroupItem

func (list *AddCasterEpisodeGroupItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterEpisodeGroupItem)
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

type AddCasterEpisodeGroupItemIdList []common.String

func (list *AddCasterEpisodeGroupItemIdList) UnmarshalJSON(data []byte) error {
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
