package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
