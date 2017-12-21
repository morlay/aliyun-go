package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveSignatureApisRequest struct {
	requests.RpcRequest
	SignatureId string `position:"Query" name:"SignatureId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	StageName   string `position:"Query" name:"StageName"`
}

func (req *RemoveSignatureApisRequest) Invoke(client *sdk.Client) (resp *RemoveSignatureApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveSignatureApis", "apigateway", "")
	resp = &RemoveSignatureApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveSignatureApisResponse struct {
	responses.BaseResponse
	RequestId string
}
