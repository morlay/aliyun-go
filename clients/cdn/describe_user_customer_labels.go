package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserCustomerLabelsRequest struct {
	requests.RpcRequest
	Uid           int64  `position:"Query" name:"Uid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeUserCustomerLabelsRequest) Invoke(client *sdk.Client) (resp *DescribeUserCustomerLabelsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserCustomerLabels", "", "")
	resp = &DescribeUserCustomerLabelsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserCustomerLabelsResponse struct {
	responses.BaseResponse
	RequestId   common.String
	IsInnerUser bool
}
