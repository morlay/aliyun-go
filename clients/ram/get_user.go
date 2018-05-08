package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	User      GetUserUser
}

type GetUserUser struct {
	UserId        common.String
	UserName      common.String
	DisplayName   common.String
	MobilePhone   common.String
	Email         common.String
	Comments      common.String
	CreateDate    common.String
	UpdateDate    common.String
	LastLoginDate common.String
}
