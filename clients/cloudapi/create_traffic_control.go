package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTrafficControlRequest struct {
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	TrafficControlUnit string `position:"Query" name:"TrafficControlUnit"`
	ApiDefault         int    `position:"Query" name:"ApiDefault"`
	UserDefault        int    `position:"Query" name:"UserDefault"`
	AppDefault         int    `position:"Query" name:"AppDefault"`
	Description        string `position:"Query" name:"Description"`
}

func (r CreateTrafficControlRequest) Invoke(client *sdk.Client) (response *CreateTrafficControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateTrafficControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateTrafficControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateTrafficControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateTrafficControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateTrafficControlResponse struct {
	RequestId        string
	TrafficControlId string
}
