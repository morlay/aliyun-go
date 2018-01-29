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
	AllowSignatureMethod   string
	ErrorCodeSamples       DescribeDeployedApiErrorCodeSampleList
	SystemParameters       DescribeDeployedApiSystemParameterList
	CustomSystemParameters DescribeDeployedApiSystemParameterList
	ConstantParameters     DescribeDeployedApiConstantParameterList
	RequestParameters      DescribeDeployedApiRequestParameterList
	ServiceParameters      DescribeDeployedApiServiceParameterList
	ServiceParametersMap   DescribeDeployedApiServiceParameterMapList
	RequestConfig          DescribeDeployedApiRequestConfig
	ServiceConfig          DescribeDeployedApiServiceConfig
	OpenIdConnectConfig    DescribeDeployedApiOpenIdConnectConfig
}

type DescribeDeployedApiErrorCodeSampleList []DescribeDeployedApiErrorCodeSample

func (list *DescribeDeployedApiErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiErrorCodeSample)
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

type DescribeDeployedApiSystemParameterList []DescribeDeployedApiSystemParameter

func (list *DescribeDeployedApiSystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiSystemParameter)
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

type DescribeDeployedApiConstantParameterList []DescribeDeployedApiConstantParameter

func (list *DescribeDeployedApiConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiConstantParameter)
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

type DescribeDeployedApiRequestParameterList []DescribeDeployedApiRequestParameter

func (list *DescribeDeployedApiRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiRequestParameter)
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

type DescribeDeployedApiServiceParameterList []DescribeDeployedApiServiceParameter

func (list *DescribeDeployedApiServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiServiceParameter)
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

type DescribeDeployedApiServiceParameterMapList []DescribeDeployedApiServiceParameterMap

func (list *DescribeDeployedApiServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApiServiceParameterMap)
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
