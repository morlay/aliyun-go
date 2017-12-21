package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteContacterRequest struct {
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (r DeleteContacterRequest) Invoke(client *sdk.Client) (response *DeleteContacterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteContacterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "DeleteContacter", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteContacterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteContacterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteContacterResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
}
