package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateUserRequest struct {
	requests.RpcRequest
	Comments    string `position:"Query" name:"Comments"`
	DisplayName string `position:"Query" name:"DisplayName"`
	MobilePhone string `position:"Query" name:"MobilePhone"`
	Email       string `position:"Query" name:"Email"`
	UserName    string `position:"Query" name:"UserName"`
}

func (req *CreateUserRequest) Invoke(client *sdk.Client) (resp *CreateUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateUser", "", "")
	resp = &CreateUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUserResponse struct {
	responses.BaseResponse
	RequestId common.String
	User      CreateUserUser
}

type CreateUserUser struct {
	UserId      common.String
	UserName    common.String
	DisplayName common.String
	MobilePhone common.String
	Email       common.String
	Comments    common.String
	CreateDate  common.String
}
