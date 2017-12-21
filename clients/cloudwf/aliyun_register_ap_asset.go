package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AliyunRegisterApAssetRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Mac       string `position:"Query" name:"Mac"`
	SerialNo  string `position:"Query" name:"SerialNo"`
}

func (r AliyunRegisterApAssetRequest) Invoke(client *sdk.Client) (response *AliyunRegisterApAssetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AliyunRegisterApAssetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AliyunRegisterApAsset", "", "")

	resp := struct {
		*responses.BaseResponse
		AliyunRegisterApAssetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AliyunRegisterApAssetResponse

	err = client.DoAction(&req, &resp)
	return
}

type AliyunRegisterApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
