package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ChangeDomainGroupRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (req *ChangeDomainGroupRequest) Invoke(client *sdk.Client) (resp *ChangeDomainGroupResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "ChangeDomainGroup", "", "")
	resp = &ChangeDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ChangeDomainGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	GroupId   common.String
	GroupName common.String
}
