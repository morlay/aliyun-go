package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindImagePackageRequest struct {
	EcsInstanceId          string `position:"Query" name:"EcsInstanceId"`
	ImagePackageInstanceId string `position:"Query" name:"ImagePackageInstanceId"`
}

func (r BindImagePackageRequest) Invoke(client *sdk.Client) (response *BindImagePackageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindImagePackageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "BindImagePackage", "", "")

	resp := struct {
		*responses.BaseResponse
		BindImagePackageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindImagePackageResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindImagePackageResponse struct {
	RequestId string
	Success   bool
}
