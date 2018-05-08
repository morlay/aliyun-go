package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveStaStatusRequest struct {
	requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
}

func (req *SaveStaStatusRequest) Invoke(client *sdk.Client) (resp *SaveStaStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveStaStatus", "", "")
	resp = &SaveStaStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveStaStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
