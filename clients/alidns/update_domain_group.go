package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDomainGroupRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (req *UpdateDomainGroupRequest) Invoke(client *sdk.Client) (resp *UpdateDomainGroupResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainGroup", "", "")
	resp = &UpdateDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDomainGroupResponse struct {
	responses.BaseResponse
	RequestId string
	GroupId   string
	GroupName string
}
