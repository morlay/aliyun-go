package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	DomainGroupId     int64
	DomainGroupName   string
	TotalNumber       int
	CreationDate      string
	ModificationDate  string
	DomainGroupStatus string
	BeingDeleted      bool
}
