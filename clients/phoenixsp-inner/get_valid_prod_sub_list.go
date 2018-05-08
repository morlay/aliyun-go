package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetValidProdSubListRequest struct {
	requests.RpcRequest
	Filter string `position:"Query" name:"Filter"`
	AliUid int64  `position:"Query" name:"AliUid"`
}

func (req *GetValidProdSubListRequest) Invoke(client *sdk.Client) (resp *GetValidProdSubListResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "GetValidProdSubList", "", "")
	resp = &GetValidProdSubListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetValidProdSubListResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Message   common.String
	Success   bool
	Data      common.String
}
