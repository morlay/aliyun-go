package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveIpControlApisRequest struct {
	requests.RpcRequest
	StageName   string `position:"Query" name:"StageName"`
	IpControlId string `position:"Query" name:"IpControlId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
}

func (req *RemoveIpControlApisRequest) Invoke(client *sdk.Client) (resp *RemoveIpControlApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveIpControlApis", "apigateway", "")
	resp = &RemoveIpControlApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveIpControlApisResponse struct {
	responses.BaseResponse
	RequestId common.String
}
