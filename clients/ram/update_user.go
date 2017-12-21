package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateUserRequest struct {
	requests.RpcRequest
	NewUserName    string `position:"Query" name:"NewUserName"`
	NewDisplayName string `position:"Query" name:"NewDisplayName"`
	NewMobilePhone string `position:"Query" name:"NewMobilePhone"`
	NewComments    string `position:"Query" name:"NewComments"`
	NewEmail       string `position:"Query" name:"NewEmail"`
	UserName       string `position:"Query" name:"UserName"`
}

func (req *UpdateUserRequest) Invoke(client *sdk.Client) (resp *UpdateUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateUser", "", "")
	resp = &UpdateUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateUserResponse struct {
	responses.BaseResponse
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
