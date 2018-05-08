package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiSignaturesRequest struct {
	requests.RpcRequest
	StageName  string `position:"Query" name:"StageName"`
	GroupId    string `position:"Query" name:"GroupId"`
	ApiIds     string `position:"Query" name:"ApiIds"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeApiSignaturesRequest) Invoke(client *sdk.Client) (resp *DescribeApiSignaturesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiSignatures", "apigateway", "")
	resp = &DescribeApiSignaturesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiSignaturesResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Integer
	PageSize      common.Integer
	PageNumber    common.Integer
	ApiSignatures DescribeApiSignaturesApiSignatureItemList
}

type DescribeApiSignaturesApiSignatureItem struct {
	ApiId         common.String
	ApiName       common.String
	SignatureId   common.String
	SignatureName common.String
	BoundTime     common.String
}

type DescribeApiSignaturesApiSignatureItemList []DescribeApiSignaturesApiSignatureItem

func (list *DescribeApiSignaturesApiSignatureItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSignaturesApiSignatureItem)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
