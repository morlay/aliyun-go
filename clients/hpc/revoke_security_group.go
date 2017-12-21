package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeSecurityGroupRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
}

func (r RevokeSecurityGroupRequest) Invoke(client *sdk.Client) (response *RevokeSecurityGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RevokeSecurityGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "RevokeSecurityGroup", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		RevokeSecurityGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RevokeSecurityGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type RevokeSecurityGroupResponse struct {
	RequestId string
}
