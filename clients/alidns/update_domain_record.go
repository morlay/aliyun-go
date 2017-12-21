package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDomainRecordRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
	TTL          int64  `position:"Query" name:"TTL"`
	Priority     int64  `position:"Query" name:"Priority"`
	Line         string `position:"Query" name:"Line"`
}

func (req *UpdateDomainRecordRequest) Invoke(client *sdk.Client) (resp *UpdateDomainRecordResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainRecord", "", "")
	resp = &UpdateDomainRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDomainRecordResponse struct {
	responses.BaseResponse
	RequestId string
	RecordId  string
}
