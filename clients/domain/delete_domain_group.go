package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainGroupRequest struct {
	requests.RpcRequest
	UserClientIp  string `position:"Query" name:"UserClientIp"`
	Lang          string `position:"Query" name:"Lang"`
	DomainGroupId int64  `position:"Query" name:"DomainGroupId"`
}

func (req *DeleteDomainGroupRequest) Invoke(client *sdk.Client) (resp *DeleteDomainGroupResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "DeleteDomainGroup", "", "")
	resp = &DeleteDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
