package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerSubmitImageSecurityAuditResultRequest struct {
	requests.RpcRequest
	ProductCode  string `position:"Query" name:"ProductCode"`
	Pass         int    `position:"Query" name:"Pass"`
	ImageVersion string `position:"Query" name:"ImageVersion"`
	Channel      string `position:"Query" name:"Channel"`
	ImageNo      string `position:"Query" name:"ImageNo"`
	Remark       string `position:"Query" name:"Remark"`
	Operator     string `position:"Query" name:"Operator"`
	RegionNo     string `position:"Query" name:"RegionNo"`
}

func (req *InnerSubmitImageSecurityAuditResultRequest) Invoke(client *sdk.Client) (resp *InnerSubmitImageSecurityAuditResultResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerSubmitImageSecurityAuditResult", "yunmarket", "")
	resp = &InnerSubmitImageSecurityAuditResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerSubmitImageSecurityAuditResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
