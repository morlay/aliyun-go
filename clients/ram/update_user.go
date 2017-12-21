package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateUserRequest struct {
	NewUserName    string `position:"Query" name:"NewUserName"`
	NewDisplayName string `position:"Query" name:"NewDisplayName"`
	NewMobilePhone string `position:"Query" name:"NewMobilePhone"`
	NewComments    string `position:"Query" name:"NewComments"`
	NewEmail       string `position:"Query" name:"NewEmail"`
	UserName       string `position:"Query" name:"UserName"`
}

func (r UpdateUserRequest) Invoke(client *sdk.Client) (response *UpdateUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateUser", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateUserResponse struct {
	RequestId string
	User      UpdateUserUser
}

type UpdateUserUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
	UpdateDate  string
}
