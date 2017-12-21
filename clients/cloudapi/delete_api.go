package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApiRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	ApiId   string `position:"Query" name:"ApiId"`
}

func (req *DeleteApiRequest) Invoke(client *sdk.Client) (resp *DeleteApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApi", "apigateway", "")
	resp = &DeleteApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApiResponse struct {
	responses.BaseResponse
	RequestId string
}
