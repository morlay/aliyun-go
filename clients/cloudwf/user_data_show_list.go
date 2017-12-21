package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserDataShowListRequest struct {
	Iid  int64  `position:"Query" name:"Iid"`
	Name string `position:"Query" name:"Name"`
	Page int    `position:"Query" name:"Page"`
	Bid  int64  `position:"Query" name:"Bid"`
	Per  int    `position:"Query" name:"Per"`
}

func (r UserDataShowListRequest) Invoke(client *sdk.Client) (response *UserDataShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UserDataShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		UserDataShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UserDataShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type UserDataShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
