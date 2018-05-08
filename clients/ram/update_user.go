package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	User      UpdateUserUser
}

type UpdateUserUser struct {
	UserId      common.String
	UserName    common.String
	DisplayName common.String
	MobilePhone common.String
	Email       common.String
	Comments    common.String
	CreateDate  common.String
	UpdateDate  common.String
}
