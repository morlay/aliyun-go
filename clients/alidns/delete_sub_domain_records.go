package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSubDomainRecordsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
}

func (req *DeleteSubDomainRecordsRequest) Invoke(client *sdk.Client) (resp *DeleteSubDomainRecordsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteSubDomainRecords", "", "")
	resp = &DeleteSubDomainRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSubDomainRecordsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	RR         common.String
	TotalCount common.String
}
