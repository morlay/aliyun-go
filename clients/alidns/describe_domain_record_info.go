package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	DomainId   common.String
	DomainName common.String
	PunyCode   common.String
	GroupId    common.String
	GroupName  common.String
	RecordId   common.String
	RR         common.String
	Type       common.String
	Value      common.String
	TTL        common.Long
	Priority   common.Long
	Line       common.String
	Status     common.String
	Locked     bool
}
