package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDrdsInstanceRequest struct {
	requests.RpcRequest
	VswitchId      string `position:"Query" name:"VswitchId"`
	Quantity       int    `position:"Query" name:"Quantity"`
	InstanceSeries string `position:"Query" name:"InstanceSeries"`
	VpcId          string `position:"Query" name:"VpcId"`
	Description    string `position:"Query" name:"Description"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	Specification  string `position:"Query" name:"Specification"`
	Type           string `position:"Query" name:"Type"`
	PayType        string `position:"Query" name:"PayType"`
}

func (req *CreateDrdsInstanceRequest) Invoke(client *sdk.Client) (resp *CreateDrdsInstanceResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsInstance", "", "")
	resp = &CreateDrdsInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDrdsInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      CreateDrdsInstanceData
}

type CreateDrdsInstanceData struct {
	OrderId            common.Long
	DrdsInstanceIdList CreateDrdsInstanceDrdsInstanceIdListList
}

type CreateDrdsInstanceDrdsInstanceIdListList []common.String

func (list *CreateDrdsInstanceDrdsInstanceIdListList) UnmarshalJSON(data []byte) error {
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
