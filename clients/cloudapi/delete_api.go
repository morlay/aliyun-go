package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApiRequest struct {
	GroupId string `position:"Query" name:"GroupId"`
	ApiId   string `position:"Query" name:"ApiId"`
}

func (r DeleteApiRequest) Invoke(client *sdk.Client) (response *DeleteApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApiResponse struct {
	RequestId string
}
