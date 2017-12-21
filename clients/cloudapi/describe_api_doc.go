package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	RegionId          string
	GroupId           string
	GroupName         string
	StageName         string
	ApiId             string
	ApiName           string
	Description       string
	Visibility        string
	AuthType          string
	ResultType        string
	ResultSample      string
	FailResultSample  string
	DeployedTime      string
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
