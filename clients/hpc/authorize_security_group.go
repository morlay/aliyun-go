package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AuthorizeSecurityGroupRequest struct {
	requests.RpcRequest
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
}

func (req *AuthorizeSecurityGroupRequest) Invoke(client *sdk.Client) (resp *AuthorizeSecurityGroupResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "AuthorizeSecurityGroup", "hpc", "")
	resp = &AuthorizeSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type AuthorizeSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
