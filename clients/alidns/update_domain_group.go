package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (r UpdateDomainGroupRequest) Invoke(client *sdk.Client) (response *UpdateDomainGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateDomainGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateDomainGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateDomainGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}
