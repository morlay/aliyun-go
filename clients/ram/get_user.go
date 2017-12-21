package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r GetUserRequest) Invoke(client *sdk.Client) (response *GetUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetUser", "", "")

	resp := struct {
		*responses.BaseResponse
		GetUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserResponse struct {
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
