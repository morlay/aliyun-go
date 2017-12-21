package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAllTrafficSpecialControlRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
}

func (r DeleteAllTrafficSpecialControlRequest) Invoke(client *sdk.Client) (response *DeleteAllTrafficSpecialControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAllTrafficSpecialControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteAllTrafficSpecialControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAllTrafficSpecialControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAllTrafficSpecialControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAllTrafficSpecialControlResponse struct {
	RequestId string
}
