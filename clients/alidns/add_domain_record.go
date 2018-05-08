package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddDomainRecordRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
	TTL          int64  `position:"Query" name:"TTL"`
	Priority     int64  `position:"Query" name:"Priority"`
	Line         string `position:"Query" name:"Line"`
}

func (req *AddDomainRecordRequest) Invoke(client *sdk.Client) (resp *AddDomainRecordResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddDomainRecord", "", "")
	resp = &AddDomainRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddDomainRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	RecordId  common.String
}
