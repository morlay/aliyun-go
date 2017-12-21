package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCustomerInfoRequest struct {
	requests.RpcRequest
	KpId    int64  `position:"Query" name:"KpId"`
	Website string `position:"Query" name:"Website"`
	Biz     string `position:"Query" name:"Biz"`
}

func (req *ModifyCustomerInfoRequest) Invoke(client *sdk.Client) (resp *ModifyCustomerInfoResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "ModifyCustomerInfo", "", "")
	resp = &ModifyCustomerInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCustomerInfoResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    string
	ResultMessage string
}
