package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApisBySignatureRequest struct {
	SignatureId string `position:"Query" name:"SignatureId"`
	PageSize    int    `position:"Query" name:"PageSize"`
	PageNumber  int    `position:"Query" name:"PageNumber"`
}

func (r DescribeApisBySignatureRequest) Invoke(client *sdk.Client) (response *DescribeApisBySignatureResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApisBySignatureRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisBySignature", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApisBySignatureResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApisBySignatureResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApisBySignatureResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	ApiInfos   DescribeApisBySignatureApiInfoList
}

type DescribeApisBySignatureApiInfo struct {
	RegionId    string
	GroupId     string
	GroupName   string
	StageName   string
	ApiId       string
	ApiName     string
	Description string
	Visibility  string
	BoundTime   string
}

type DescribeApisBySignatureApiInfoList []DescribeApisBySignatureApiInfo

func (list *DescribeApisBySignatureApiInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisBySignatureApiInfo)
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
