package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForTransferProhibitionLockRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SaveSingleTaskForTransferProhibitionLockRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForTransferProhibitionLockResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForTransferProhibitionLock", "domain", "")
	resp = &SaveSingleTaskForTransferProhibitionLockResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForTransferProhibitionLockResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
