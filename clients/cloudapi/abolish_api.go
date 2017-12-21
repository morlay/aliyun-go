package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AbolishApiRequest struct {
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
}

func (r AbolishApiRequest) Invoke(client *sdk.Client) (response *AbolishApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AbolishApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AbolishApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		AbolishApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AbolishApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type AbolishApiResponse struct {
	RequestId string
}
