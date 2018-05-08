package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForUpdateProhibitionLockRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SaveSingleTaskForUpdateProhibitionLockRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForUpdateProhibitionLockResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForUpdateProhibitionLock", "", "")
	resp = &SaveSingleTaskForUpdateProhibitionLockResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForUpdateProhibitionLockResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
