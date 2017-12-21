package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRecordInfoRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
}

func (req *DescribeDomainRecordInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRecordInfoResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainRecordInfo", "", "")
	resp = &DescribeDomainRecordInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRecordInfoResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainId   string
	DomainName string
	PunyCode   string
	GroupId    string
	GroupName  string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
	Status     string
	Locked     bool
}
