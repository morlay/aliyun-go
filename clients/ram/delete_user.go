package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteUserRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *DeleteUserRequest) Invoke(client *sdk.Client) (resp *DeleteUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteUser", "", "")
	resp = &DeleteUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteUserResponse struct {
	responses.BaseResponse
	RequestId string
}
