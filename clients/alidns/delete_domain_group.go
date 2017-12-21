package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (r DeleteDomainGroupRequest) Invoke(client *sdk.Client) (response *DeleteDomainGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDomainGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDomainGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDomainGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDomainGroupResponse struct {
	RequestId string
	GroupName string
}
