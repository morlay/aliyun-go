package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeployApiRequest struct {
	GroupId     string `position:"Query" name:"GroupId"`
	ApiId       string `position:"Query" name:"ApiId"`
	StageName   string `position:"Query" name:"StageName"`
	Description string `position:"Query" name:"Description"`
}

func (r DeployApiRequest) Invoke(client *sdk.Client) (response *DeployApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeployApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeployApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeployApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeployApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeployApiResponse struct {
	RequestId string
}
