package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetSignatureApisRequest struct {
	requests.RpcRequest
	SignatureId string `position:"Query" name:"SignatureId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	StageName   string `position:"Query" name:"StageName"`
}

func (req *SetSignatureApisRequest) Invoke(client *sdk.Client) (resp *SetSignatureApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetSignatureApis", "apigateway", "")
	resp = &SetSignatureApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetSignatureApisResponse struct {
	responses.BaseResponse
	RequestId common.String
}
