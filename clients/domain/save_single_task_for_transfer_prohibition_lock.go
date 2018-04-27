package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForTransferProhibitionLockRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SaveSingleTaskForTransferProhibitionLockRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForTransferProhibitionLockResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForTransferProhibitionLock", "", "")
	resp = &SaveSingleTaskForTransferProhibitionLockResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForTransferProhibitionLockResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
