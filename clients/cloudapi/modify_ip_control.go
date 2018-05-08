package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyIpControlRequest struct {
	requests.RpcRequest
	IpControlId   string `position:"Query" name:"IpControlId"`
	IpControlName string `position:"Query" name:"IpControlName"`
	Description   string `position:"Query" name:"Description"`
}

func (req *ModifyIpControlRequest) Invoke(client *sdk.Client) (resp *ModifyIpControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyIpControl", "apigateway", "")
	resp = &ModifyIpControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyIpControlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
