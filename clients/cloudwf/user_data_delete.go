package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataDeleteRequest struct {
	Iid int64 `position:"Query" name:"Iid"`
	Bid int64 `position:"Query" name:"Bid"`
}

func (r UserDataDeleteRequest) Invoke(client *sdk.Client) (response *UserDataDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UserDataDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		UserDataDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UserDataDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type UserDataDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
