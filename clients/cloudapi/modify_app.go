package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAppRequest struct {
	AppId       int64  `position:"Query" name:"AppId"`
	AppName     string `position:"Query" name:"AppName"`
	Description string `position:"Query" name:"Description"`
}

func (r ModifyAppRequest) Invoke(client *sdk.Client) (response *ModifyAppResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyAppRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApp", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ModifyAppResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyAppResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyAppResponse struct {
	RequestId string
}
