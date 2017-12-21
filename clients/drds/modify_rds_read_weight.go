package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRdsReadWeightRequest struct {
	InstanceNames  string `position:"Query" name:"InstanceNames"`
	DbName         string `position:"Query" name:"DbName"`
	Weights        string `position:"Query" name:"Weights"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r ModifyRdsReadWeightRequest) Invoke(client *sdk.Client) (response *ModifyRdsReadWeightResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyRdsReadWeightRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyRdsReadWeight", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyRdsReadWeightResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyRdsReadWeightResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyRdsReadWeightResponse struct {
	RequestId string
	Success   bool
}
