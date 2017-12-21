package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteTrafficControlRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
}

func (r DeleteTrafficControlRequest) Invoke(client *sdk.Client) (response *DeleteTrafficControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteTrafficControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteTrafficControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteTrafficControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteTrafficControlResponse struct {
	RequestId string
}
