package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRdsReadWeightRequest struct {
	requests.RpcRequest
	InstanceNames  string `position:"Query" name:"InstanceNames"`
	DbName         string `position:"Query" name:"DbName"`
	Weights        string `position:"Query" name:"Weights"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *ModifyRdsReadWeightRequest) Invoke(client *sdk.Client) (resp *ModifyRdsReadWeightResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyRdsReadWeight", "", "")
	resp = &ModifyRdsReadWeightResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyRdsReadWeightResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
