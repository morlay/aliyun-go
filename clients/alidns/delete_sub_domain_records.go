package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSubDomainRecordsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
}

func (r DeleteSubDomainRecordsRequest) Invoke(client *sdk.Client) (response *DeleteSubDomainRecordsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSubDomainRecordsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteSubDomainRecords", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSubDomainRecordsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSubDomainRecordsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSubDomainRecordsResponse struct {
	RequestId  string
	RR         string
	TotalCount string
}
