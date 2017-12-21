package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

func (r UpdateBatchDomainRecordsRequest) Invoke(client *sdk.Client) (response *UpdateBatchDomainRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateBatchDomainRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateBatchDomainRecords", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateBatchDomainRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateBatchDomainRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}
