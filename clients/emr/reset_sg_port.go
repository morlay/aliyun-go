package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetSgPortRequest struct {
	requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	OperateType string `position:"Query" name:"OperateType"`
}

func (req *ResetSgPortRequest) Invoke(client *sdk.Client) (resp *ResetSgPortResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResetSgPort", "", "")
	resp = &ResetSgPortResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetSgPortResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   common.String
	ErrCode   common.String
	ErrMsg    common.String
}
