package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AbolishApiRequest struct {
	requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
}

func (req *AbolishApiRequest) Invoke(client *sdk.Client) (resp *AbolishApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AbolishApi", "apigateway", "")
	resp = &AbolishApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type AbolishApiResponse struct {
	responses.BaseResponse
	RequestId common.String
}
