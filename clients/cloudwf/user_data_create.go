package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataCreateRequest struct {
	UploadFile string `position:"Query" name:"UploadFile"`
	Name       string `position:"Query" name:"Name"`
	Bid        int64  `position:"Query" name:"Bid"`
	Type       string `position:"Query" name:"Type"`
}

func (r UserDataCreateRequest) Invoke(client *sdk.Client) (response *UserDataCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UserDataCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		UserDataCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UserDataCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type UserDataCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
