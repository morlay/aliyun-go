package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLoginProfileRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r DeleteLoginProfileRequest) Invoke(client *sdk.Client) (response *DeleteLoginProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLoginProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteLoginProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLoginProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLoginProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLoginProfileResponse struct {
	RequestId string
}
