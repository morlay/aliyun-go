package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

func (r AddBatchDomainRecordsRequest) Invoke(client *sdk.Client) (response *AddBatchDomainRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddBatchDomainRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "AddBatchDomainRecords", "", "")

	resp := struct {
		*responses.BaseResponse
		AddBatchDomainRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddBatchDomainRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}
