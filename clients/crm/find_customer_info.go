package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindCustomerInfoRequest struct {
	requests.RpcRequest
	KpId int64 `position:"Query" name:"KpId"`
}

func (req *FindCustomerInfoRequest) Invoke(client *sdk.Client) (resp *FindCustomerInfoResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "FindCustomerInfo", "", "")
	resp = &FindCustomerInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindCustomerInfoResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          FindCustomerInfoData
}

type FindCustomerInfoData struct {
	Website string
	Biz     string
}
