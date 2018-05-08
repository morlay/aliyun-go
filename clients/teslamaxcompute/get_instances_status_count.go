package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetInstancesStatusCountRequest struct {
	requests.RpcRequest
	Cluster string `position:"Query" name:"Cluster"`
	Region  string `position:"Query" name:"Region"`
}

func (req *GetInstancesStatusCountRequest) Invoke(client *sdk.Client) (resp *GetInstancesStatusCountResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetInstancesStatusCount", "", "")
	resp = &GetInstancesStatusCountResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstancesStatusCountResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetInstancesStatusCountDataItemList
}

type GetInstancesStatusCountDataItem struct {
	Status common.String
	Size   common.Integer
}

type GetInstancesStatusCountDataItemList []GetInstancesStatusCountDataItem

func (list *GetInstancesStatusCountDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetInstancesStatusCountDataItem)
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
