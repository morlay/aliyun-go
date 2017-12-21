package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyApiGroupRequest struct {
	GroupId     string `position:"Query" name:"GroupId"`
	GroupName   string `position:"Query" name:"GroupName"`
	Description string `position:"Query" name:"Description"`
}

func (r ModifyApiGroupRequest) Invoke(client *sdk.Client) (response *ModifyApiGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyApiGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApiGroup", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ModifyApiGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyApiGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyApiGroupResponse struct {
	RequestId   string
	GroupId     string
	GroupName   string
	SubDomain   string
	Description string
}
