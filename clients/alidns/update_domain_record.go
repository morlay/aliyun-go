package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDomainRecordRequest struct {
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

func (r UpdateDomainRecordRequest) Invoke(client *sdk.Client) (response *UpdateDomainRecordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateDomainRecordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainRecord", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateDomainRecordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateDomainRecordResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateDomainRecordResponse struct {
	RequestId string
	RecordId  string
}
