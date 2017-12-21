package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDeployedApiRequest struct {
	requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
}

func (req *DescribeDeployedApiRequest) Invoke(client *sdk.Client) (resp *DescribeDeployedApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDeployedApi", "apigateway", "")
	resp = &DescribeDeployedApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDeployedApiResponse struct {
	responses.BaseResponse
	RequestId              string
	RegionId               string
	GroupId                string
	GroupName              string
	StageName              string
	ApiId                  string
	ApiName                string
	Description            string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	DeployedTime           string
	ErrorCodeSamples       DescribeDeployedApiErrorCodeSampleList
	SystemParameters       DescribeDeployedApiSystemParameterList
	CustomSystemParameters DescribeDeployedApiSystemParameterList
	ConstantParameters     DescribeDeployedApiConstantParameterList
	RequestParameters      DescribeDeployedApiRequestParameterList
	ServiceParameters      DescribeDeployedApiServiceParameterList
	ServiceParametersMap   DescribeDeployedApiServiceParameterMapList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeDeployedApiErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeDeployedApiErrorCodeSampleList) UnmarshalJSON(data []byte) error {
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

type DescribeDeployedApiSystemParameterList []DescribeApiSystemParameter

func (list *DescribeDeployedApiSystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSystemParameter)
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

type DescribeDeployedApiConstantParameterList []DescribeApiConstantParameter

func (list *DescribeDeployedApiConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiConstantParameter)
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

type DescribeDeployedApiRequestParameterList []DescribeApiRequestParameter

func (list *DescribeDeployedApiRequestParameterList) UnmarshalJSON(data []byte) error {
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

type DescribeDeployedApiServiceParameterList []DescribeApiServiceParameter

func (list *DescribeDeployedApiServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameter)
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

type DescribeDeployedApiServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeDeployedApiServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameterMap)
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
