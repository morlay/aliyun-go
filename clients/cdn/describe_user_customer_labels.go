package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserCustomerLabelsRequest struct {
	Uid           int64  `position:"Query" name:"Uid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeUserCustomerLabelsRequest) Invoke(client *sdk.Client) (response *DescribeUserCustomerLabelsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserCustomerLabelsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserCustomerLabels", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserCustomerLabelsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserCustomerLabelsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserCustomerLabelsResponse struct {
	RequestId   string
	IsInnerUser bool
}
