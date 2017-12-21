package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAppAttributesRequest struct {
	requests.RpcRequest
	AppId      int64 `position:"Query" name:"AppId"`
	PageNumber int   `position:"Query" name:"PageNumber"`
	PageSize   int   `position:"Query" name:"PageSize"`
}

func (req *DescribeAppAttributesRequest) Invoke(client *sdk.Client) (resp *DescribeAppAttributesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppAttributes", "apigateway", "")
	resp = &DescribeAppAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAppAttributesResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	Apps       DescribeAppAttributesAppAttributeList
}

type DescribeAppAttributesAppAttribute struct {
	AppId        int64
	AppName      string
	Description  string
	CreatedTime  string
	ModifiedTime string
}

type DescribeAppAttributesAppAttributeList []DescribeAppAttributesAppAttribute

func (list *DescribeAppAttributesAppAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAppAttributesAppAttribute)
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
