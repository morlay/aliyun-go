package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBatchDomainRecordsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

func (req *DeleteBatchDomainRecordsRequest) Invoke(client *sdk.Client) (resp *DeleteBatchDomainRecordsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomainRecords", "", "")
	resp = &DeleteBatchDomainRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBatchDomainRecordsResponse struct {
	responses.BaseResponse
	RequestId common.String
	TraceId   common.String
}
