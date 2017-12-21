package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddUserToGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
}

func (r AddUserToGroupRequest) Invoke(client *sdk.Client) (response *AddUserToGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddUserToGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "AddUserToGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		AddUserToGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddUserToGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddUserToGroupResponse struct {
	RequestId string
}
