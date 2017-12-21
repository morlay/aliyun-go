package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (r AddDomainGroupRequest) Invoke(client *sdk.Client) (response *AddDomainGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddDomainGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddDomainGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		AddDomainGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddDomainGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}
