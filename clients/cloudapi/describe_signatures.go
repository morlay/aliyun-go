package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSignaturesRequest struct {
	requests.RpcRequest
	SignatureId   string `position:"Query" name:"SignatureId"`
	SignatureName string `position:"Query" name:"SignatureName"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (req *DescribeSignaturesRequest) Invoke(client *sdk.Client) (resp *DescribeSignaturesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSignatures", "apigateway", "")
	resp = &DescribeSignaturesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSignaturesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalCount     common.Integer
	PageSize       common.Integer
	PageNumber     common.Integer
	SignatureInfos DescribeSignaturesSignatureInfoList
}

type DescribeSignaturesSignatureInfo struct {
	RegionId        common.String
	SignatureId     common.String
	SignatureName   common.String
	SignatureKey    common.String
	SignatureSecret common.String
	CreatedTime     common.String
	ModifiedTime    common.String
}

type DescribeSignaturesSignatureInfoList []DescribeSignaturesSignatureInfo

func (list *DescribeSignaturesSignatureInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSignaturesSignatureInfo)
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
