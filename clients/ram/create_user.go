package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUserRequest struct {
	Comments    string `position:"Query" name:"Comments"`
	DisplayName string `position:"Query" name:"DisplayName"`
	MobilePhone string `position:"Query" name:"MobilePhone"`
	Email       string `position:"Query" name:"Email"`
	UserName    string `position:"Query" name:"UserName"`
}

func (r CreateUserRequest) Invoke(client *sdk.Client) (response *CreateUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateUser", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateUserResponse struct {
	RequestId string
	User      CreateUserUser
}

type CreateUserUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
}
