package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyTrafficControlRequest struct {
	TrafficControlId   string `position:"Query" name:"TrafficControlId"`
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	TrafficControlUnit string `position:"Query" name:"TrafficControlUnit"`
	ApiDefault         int    `position:"Query" name:"ApiDefault"`
	UserDefault        int    `position:"Query" name:"UserDefault"`
	AppDefault         int    `position:"Query" name:"AppDefault"`
	Description        string `position:"Query" name:"Description"`
}

func (r ModifyTrafficControlRequest) Invoke(client *sdk.Client) (response *ModifyTrafficControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyTrafficControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyTrafficControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ModifyTrafficControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyTrafficControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyTrafficControlResponse struct {
	RequestId string
}
