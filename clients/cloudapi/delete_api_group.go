package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApiGroupRequest struct {
	GroupId string `position:"Query" name:"GroupId"`
}

func (r DeleteApiGroupRequest) Invoke(client *sdk.Client) (response *DeleteApiGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApiGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiGroup", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApiGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApiGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApiGroupResponse struct {
	RequestId string
}
