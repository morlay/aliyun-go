package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetIpControlApisRequest struct {
	requests.RpcRequest
	StageName   string `position:"Query" name:"StageName"`
	IpControlId string `position:"Query" name:"IpControlId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
}

func (req *SetIpControlApisRequest) Invoke(client *sdk.Client) (resp *SetIpControlApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetIpControlApis", "apigateway", "")
	resp = &SetIpControlApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetIpControlApisResponse struct {
	responses.BaseResponse
	RequestId string
}
