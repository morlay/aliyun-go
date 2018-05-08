package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetStatisticsByUuidRequest struct {
	requests.RpcRequest
	Uuid string `position:"Query" name:"Uuid"`
}

func (req *GetStatisticsByUuidRequest) Invoke(client *sdk.Client) (resp *GetStatisticsByUuidResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetStatisticsByUuid", "vipaegis", "")
	resp = &GetStatisticsByUuidResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStatisticsByUuidResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      GetStatisticsByUuidEntityList
}

type GetStatisticsByUuidEntity struct {
	Uuid    common.String
	Account common.Integer
	Health  common.Integer
	Patch   common.Integer
	Trojan  common.Integer
	Online  bool
}

type GetStatisticsByUuidEntityList []GetStatisticsByUuidEntity

func (list *GetStatisticsByUuidEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetStatisticsByUuidEntity)
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
