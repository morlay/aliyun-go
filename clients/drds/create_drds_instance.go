package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDrdsInstanceRequest struct {
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

func (r CreateDrdsInstanceRequest) Invoke(client *sdk.Client) (response *CreateDrdsInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDrdsInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateDrdsInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDrdsInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDrdsInstanceResponse struct {
	RequestId string
	Success   bool
	Data      CreateDrdsInstanceData
}

type CreateDrdsInstanceData struct {
	OrderId            int64
	DrdsInstanceIdList CreateDrdsInstanceDrdsInstanceIdListList
}

type CreateDrdsInstanceDrdsInstanceIdListList []string

func (list *CreateDrdsInstanceDrdsInstanceIdListList) UnmarshalJSON(data []byte) error {
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
