package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeployApiRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	ApiId       string `position:"Query" name:"ApiId"`
	StageName   string `position:"Query" name:"StageName"`
	Description string `position:"Query" name:"Description"`
}

func (req *DeployApiRequest) Invoke(client *sdk.Client) (resp *DeployApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeployApi", "apigateway", "")
	resp = &DeployApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeployApiResponse struct {
	responses.BaseResponse
	RequestId string
}
