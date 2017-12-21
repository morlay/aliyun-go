package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAccessGroupRequest struct {
	Description     string `position:"Query" name:"Description"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
}

func (r ModifyAccessGroupRequest) Invoke(client *sdk.Client) (response *ModifyAccessGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyAccessGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyAccessGroup", "nas", "")

	resp := struct {
		*responses.BaseResponse
		ModifyAccessGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyAccessGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyAccessGroupResponse struct {
	RequestId string
}
