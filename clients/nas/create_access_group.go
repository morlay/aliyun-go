package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAccessGroupRequest struct {
	Description     string `position:"Query" name:"Description"`
	AccessGroupType string `position:"Query" name:"AccessGroupType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (r CreateAccessGroupRequest) Invoke(client *sdk.Client) (response *CreateAccessGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAccessGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessGroup", "nas", "")

	resp := struct {
		*responses.BaseResponse
		CreateAccessGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAccessGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAccessGroupResponse struct {
	RequestId       string
	AccessGroupName string
}
