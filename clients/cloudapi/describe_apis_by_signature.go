package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApisBySignatureRequest struct {
	requests.RpcRequest
	SignatureId string `position:"Query" name:"SignatureId"`
	PageSize    int    `position:"Query" name:"PageSize"`
	PageNumber  int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeApisBySignatureRequest) Invoke(client *sdk.Client) (resp *DescribeApisBySignatureResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisBySignature", "apigateway", "")
	resp = &DescribeApisBySignatureResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApisBySignatureResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	ApiInfos   DescribeApisBySignatureApiInfoList
}

type DescribeApisBySignatureApiInfo struct {
	RegionId    common.String
	GroupId     common.String
	GroupName   common.String
	StageName   common.String
	ApiId       common.String
	ApiName     common.String
	Description common.String
	Visibility  common.String
	BoundTime   common.String
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
