package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDomainGroupRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (req *DeleteDomainGroupRequest) Invoke(client *sdk.Client) (resp *DeleteDomainGroupResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainGroup", "", "")
	resp = &DeleteDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	GroupName common.String
}
