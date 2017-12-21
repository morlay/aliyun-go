package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *GetUserRequest) Invoke(client *sdk.Client) (resp *GetUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetUser", "", "")
	resp = &GetUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserResponse struct {
	responses.BaseResponse
	RequestId string
	User      GetUserUser
}

type GetUserUser struct {
	UserId        string
	UserName      string
	DisplayName   string
	MobilePhone   string
	Email         string
	Comments      string
	CreateDate    string
	UpdateDate    string
	LastLoginDate string
}
