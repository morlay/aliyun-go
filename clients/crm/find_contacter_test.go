package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindContacterTestRequest struct {
	requests.RpcRequest
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (req *FindContacterTestRequest) Invoke(client *sdk.Client) (resp *FindContacterTestResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "FindContacterTest", "", "")
	resp = &FindContacterTestResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindContacterTestResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    common.String
	ResultMessage common.String
	Data          FindContacterTestData
}

type FindContacterTestData struct {
	ContacterId       common.Long
	KpId              common.Long
	CustomerId        common.Long
	ContacterName     common.String
	ContacterEmail    common.String
	ContacterMobile   common.String
	ContacterPosition common.String
}
