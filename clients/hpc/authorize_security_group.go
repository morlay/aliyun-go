package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AuthorizeSecurityGroupRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Policy   string `position:"Query" name:"Policy"`
	Priority string `position:"Query" name:"Priority"`
	NicType  string `position:"Query" name:"NicType"`
}

func (r AuthorizeSecurityGroupRequest) Invoke(client *sdk.Client) (response *AuthorizeSecurityGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AuthorizeSecurityGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "AuthorizeSecurityGroup", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		AuthorizeSecurityGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AuthorizeSecurityGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type AuthorizeSecurityGroupResponse struct {
	RequestId string
}
