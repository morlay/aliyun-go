package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBatchDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

func (r DeleteBatchDomainRecordsRequest) Invoke(client *sdk.Client) (response *DeleteBatchDomainRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBatchDomainRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomainRecords", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBatchDomainRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBatchDomainRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}
