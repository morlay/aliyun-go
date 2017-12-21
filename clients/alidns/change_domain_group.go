package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ChangeDomainGroupRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (r ChangeDomainGroupRequest) Invoke(client *sdk.Client) (response *ChangeDomainGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ChangeDomainGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "ChangeDomainGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		ChangeDomainGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ChangeDomainGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ChangeDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}
