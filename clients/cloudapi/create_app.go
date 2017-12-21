package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAppRequest struct {
	AppName     string `position:"Query" name:"AppName"`
	Description string `position:"Query" name:"Description"`
}

func (r CreateAppRequest) Invoke(client *sdk.Client) (response *CreateAppResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAppRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApp", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateAppResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAppResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAppResponse struct {
	RequestId string
	AppId     int64
}
