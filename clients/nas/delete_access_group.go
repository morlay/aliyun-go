package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccessGroupRequest struct {
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (r DeleteAccessGroupRequest) Invoke(client *sdk.Client) (response *DeleteAccessGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAccessGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteAccessGroup", "nas", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAccessGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAccessGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAccessGroupResponse struct {
	RequestId string
}
