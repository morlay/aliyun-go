package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveUserFromGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
}

func (r RemoveUserFromGroupRequest) Invoke(client *sdk.Client) (response *RemoveUserFromGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveUserFromGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "RemoveUserFromGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveUserFromGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveUserFromGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveUserFromGroupResponse struct {
	RequestId string
}
