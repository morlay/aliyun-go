package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r DeleteUserRequest) Invoke(client *sdk.Client) (response *DeleteUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteUser", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteUserResponse struct {
	RequestId string
}
