package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveStaStatusRequest struct {
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
}

func (r SaveStaStatusRequest) Invoke(client *sdk.Client) (response *SaveStaStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveStaStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveStaStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveStaStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveStaStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveStaStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
