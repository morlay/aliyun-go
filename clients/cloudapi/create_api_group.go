package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateApiGroupRequest struct {
	GroupName   string `position:"Query" name:"GroupName"`
	Description string `position:"Query" name:"Description"`
}

func (r CreateApiGroupRequest) Invoke(client *sdk.Client) (response *CreateApiGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateApiGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiGroup", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateApiGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateApiGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateApiGroupResponse struct {
	RequestId   string
	GroupId     string
	GroupName   string
	SubDomain   string
	Description string
}
