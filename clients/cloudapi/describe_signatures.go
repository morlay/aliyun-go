package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSignaturesRequest struct {
	SignatureId   string `position:"Query" name:"SignatureId"`
	SignatureName string `position:"Query" name:"SignatureName"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (r DescribeSignaturesRequest) Invoke(client *sdk.Client) (response *DescribeSignaturesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSignaturesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSignatures", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSignaturesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSignaturesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSignaturesResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	SignatureInfos DescribeSignaturesSignatureInfoList
}

type DescribeSignaturesSignatureInfo struct {
	RegionId        string
	SignatureId     string
	SignatureName   string
	SignatureKey    string
	SignatureSecret string
	CreatedTime     string
	ModifiedTime    string
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
