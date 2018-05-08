package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddDomainGroupRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (req *AddDomainGroupRequest) Invoke(client *sdk.Client) (resp *AddDomainGroupResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddDomainGroup", "", "")
	resp = &AddDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddDomainGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	GroupId   common.String
	GroupName common.String
}
