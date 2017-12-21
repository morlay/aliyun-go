package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AliyunRegisterApAssetRequest struct {
	requests.RpcRequest
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Mac       string `position:"Query" name:"Mac"`
	SerialNo  string `position:"Query" name:"SerialNo"`
}

func (req *AliyunRegisterApAssetRequest) Invoke(client *sdk.Client) (resp *AliyunRegisterApAssetResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AliyunRegisterApAsset", "", "")
	resp = &AliyunRegisterApAssetResponse{}
	err = client.DoAction(req, resp)
	return
}

type AliyunRegisterApAssetResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
