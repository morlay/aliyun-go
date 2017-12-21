package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeSecurityGroupRequest struct {
	requests.RpcRequest
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
}

func (req *RevokeSecurityGroupRequest) Invoke(client *sdk.Client) (resp *RevokeSecurityGroupResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "RevokeSecurityGroup", "hpc", "")
	resp = &RevokeSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type RevokeSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
