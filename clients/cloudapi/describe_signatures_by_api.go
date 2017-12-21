package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSignaturesByApiRequest struct {
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
}

func (r DescribeSignaturesByApiRequest) Invoke(client *sdk.Client) (response *DescribeSignaturesByApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSignaturesByApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSignaturesByApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSignaturesByApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSignaturesByApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSignaturesByApiResponse struct {
	RequestId  string
	Signatures DescribeSignaturesByApiSignatureItemList
}

type DescribeSignaturesByApiSignatureItem struct {
	SignatureId   string
	SignatureName string
	BoundTime     string
}

type DescribeSignaturesByApiSignatureItemList []DescribeSignaturesByApiSignatureItem

func (list *DescribeSignaturesByApiSignatureItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSignaturesByApiSignatureItem)
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
