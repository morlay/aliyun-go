package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiRequest struct {
	GroupId string `position:"Query" name:"GroupId"`
	ApiId   string `position:"Query" name:"ApiId"`
}

func (r DescribeApiRequest) Invoke(client *sdk.Client) (response *DescribeApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiResponse struct {
	RequestId              string
	RegionId               string
	ApiId                  string
	ApiName                string
	GroupId                string
	GroupName              string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	CreatedTime            string
	ModifiedTime           string
	Description            string
	SystemParameters       DescribeApiSystemParameterList
	CustomSystemParameters DescribeApiSystemParameterList
	ConstantParameters     DescribeApiConstantParameterList
	RequestParameters      DescribeApiRequestParameterList
	ServiceParameters      DescribeApiServiceParameterList
	ServiceParametersMap   DescribeApiServiceParameterMapList
	ErrorCodeSamples       DescribeApiErrorCodeSampleList
	DeployedInfos          DescribeApiDeployedInfoList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeApiSystemParameter struct {
	ParameterName        string
	ServiceParameterName string
	Location             string
	DemoValue            string
	Description          string
}

type DescribeApiConstantParameter struct {
	ServiceParameterName string
	ConstantValue        string
	Location             string
	Description          string
}

type DescribeApiRequestParameter struct {
	ApiParameterName  string
	Location          string
	ParameterType     string
	Required          string
	DefaultValue      string
	DemoValue         string
	MaxValue          int64
	MinValue          int64
	MaxLength         int64
	MinLength         int64
	RegularExpression string
	JsonScheme        string
	EnumValue         string
	DocShow           string
	DocOrder          int
	Description       string
}

type DescribeApiServiceParameter struct {
	ServiceParameterName string
	Location             string
	ParameterType        string
}

type DescribeApiServiceParameterMap struct {
	ServiceParameterName string
	RequestParameterName string
}

type DescribeApiDeployedInfo struct {
	StageName        string
	EffectiveVersion string
	DeployedStatus   string
}

type DescribeApiErrorCodeSample struct {
	Code        string
	Message     string
	Description string
}

type DescribeApiRequestConfig struct {
	RequestProtocol     string
	RequestHttpMethod   string
	RequestPath         string
	BodyFormat          string
	PostBodyDescription string
	RequestMode         string
}

type DescribeApiServiceConfig struct {
	ServiceProtocol   string
	ServiceAddress    string
	ServiceHttpMethod string
	ServicePath       string
	ServiceTimeout    int
	Mock              string
	MockResult        string
	ServiceVpcEnable  bool
	VpcConfig         DescribeApiVpcConfig
}

type DescribeApiOpenIdConnectConfig struct {
	OpenIdApiType    string
	IdTokenParamName string
	PublicKeyId      string
	PublicKey        string
}

type DescribeApiVpcConfig struct {
	VpcId      string
	InstanceId string
	Port       int
}

type DescribeApiSystemParameterList []DescribeApiSystemParameter

func (list *DescribeApiSystemParameterList) UnmarshalJSON(data []byte) error {
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

type DescribeApiConstantParameterList []DescribeApiConstantParameter

func (list *DescribeApiConstantParameterList) UnmarshalJSON(data []byte) error {
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

type DescribeApiRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiRequestParameterList) UnmarshalJSON(data []byte) error {
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

type DescribeApiServiceParameterList []DescribeApiServiceParameter

func (list *DescribeApiServiceParameterList) UnmarshalJSON(data []byte) error {
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

type DescribeApiServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeApiServiceParameterMapList) UnmarshalJSON(data []byte) error {
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

type DescribeApiErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiErrorCodeSampleList) UnmarshalJSON(data []byte) error {
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

type DescribeApiDeployedInfoList []DescribeApiDeployedInfo

func (list *DescribeApiDeployedInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiDeployedInfo)
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
