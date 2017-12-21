package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAppRequest struct {
	AppId int64 `position:"Query" name:"AppId"`
}

func (r DeleteAppRequest) Invoke(client *sdk.Client) (response *DeleteAppResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAppRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApp", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAppResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAppResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAppResponse struct {
	RequestId string
}
