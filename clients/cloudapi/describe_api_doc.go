package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiDocRequest struct {
	requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	StageName string `position:"Query" name:"StageName"`
	ApiId     string `position:"Query" name:"ApiId"`
}

func (req *DescribeApiDocRequest) Invoke(client *sdk.Client) (resp *DescribeApiDocResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiDoc", "apigateway", "")
	resp = &DescribeApiDocResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiDocResponse struct {
	responses.BaseResponse
	RequestId         common.String
	RegionId          common.String
	GroupId           common.String
	GroupName         common.String
	StageName         common.String
	ApiId             common.String
	ApiName           common.String
	Description       common.String
	Visibility        common.String
	AuthType          common.String
	ResultType        common.String
	ResultSample      common.String
	FailResultSample  common.String
	DeployedTime      common.String
	ErrorCodeSamples  DescribeApiDocErrorCodeSampleList
	RequestParameters DescribeApiDocRequestParameterList
	RequestConfig     DescribeApiRequestConfig
}

type DescribeApiDocErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiDocErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeApiDocRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiDocRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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
