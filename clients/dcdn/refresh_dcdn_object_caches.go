package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshDcdnObjectCachesRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ObjectType    string `position:"Query" name:"ObjectType"`
}

func (req *RefreshDcdnObjectCachesRequest) Invoke(client *sdk.Client) (resp *RefreshDcdnObjectCachesResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "RefreshDcdnObjectCaches", "dcdn", "")
	resp = &RefreshDcdnObjectCachesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshDcdnObjectCachesResponse struct {
	responses.BaseResponse
	RequestId     string
	RefreshTaskId string
}
