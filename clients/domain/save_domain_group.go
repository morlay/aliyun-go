package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveDomainGroupRequest struct {
	requests.RpcRequest
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	DomainGroupName string `position:"Query" name:"DomainGroupName"`
	Lang            string `position:"Query" name:"Lang"`
	DomainGroupId   int64  `position:"Query" name:"DomainGroupId"`
}

func (req *SaveDomainGroupRequest) Invoke(client *sdk.Client) (resp *SaveDomainGroupResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveDomainGroup", "", "")
	resp = &SaveDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveDomainGroupResponse struct {
	responses.BaseResponse
	RequestId         common.String
	DomainGroupId     common.Long
	DomainGroupName   common.String
	TotalNumber       common.Integer
	CreationDate      common.String
	ModificationDate  common.String
	DomainGroupStatus common.String
	BeingDeleted      bool
}
