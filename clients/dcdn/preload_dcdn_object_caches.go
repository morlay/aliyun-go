package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PreloadDcdnObjectCachesRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *PreloadDcdnObjectCachesRequest) Invoke(client *sdk.Client) (resp *PreloadDcdnObjectCachesResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "PreloadDcdnObjectCaches", "dcdn", "")
	resp = &PreloadDcdnObjectCachesResponse{}
	err = client.DoAction(req, resp)
	return
}

type PreloadDcdnObjectCachesResponse struct {
	responses.BaseResponse
	RequestId     common.String
	PreloadTaskId common.String
}
