package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	ApiId   string `position:"Query" name:"ApiId"`
}

func (req *DescribeApiRequest) Invoke(client *sdk.Client) (resp *DescribeApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApi", "apigateway", "")
	resp = &DescribeApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiResponse struct {
	responses.BaseResponse
	RequestId              common.String
	RegionId               common.String
	ApiId                  common.String
	ApiName                common.String
	GroupId                common.String
	GroupName              common.String
	Visibility             common.String
	AuthType               common.String
	ResultType             common.String
	ResultSample           common.String
	FailResultSample       common.String
	CreatedTime            common.String
	ModifiedTime           common.String
	Description            common.String
	AllowSignatureMethod   common.String
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
	ParameterName        common.String
	ServiceParameterName common.String
	Location             common.String
	DemoValue            common.String
	Description          common.String
}

type DescribeApiConstantParameter struct {
	ServiceParameterName common.String
	ConstantValue        common.String
	Location             common.String
	Description          common.String
}

type DescribeApiRequestParameter struct {
	ApiParameterName  common.String
	Location          common.String
	ParameterType     common.String
	Required          common.String
	DefaultValue      common.String
	DemoValue         common.String
	MaxValue          common.Long
	MinValue          common.Long
	MaxLength         common.Long
	MinLength         common.Long
	RegularExpression common.String
	JsonScheme        common.String
	EnumValue         common.String
	DocShow           common.String
	DocOrder          common.Integer
	Description       common.String
}

type DescribeApiServiceParameter struct {
	ServiceParameterName common.String
	Location             common.String
	ParameterType        common.String
}

type DescribeApiServiceParameterMap struct {
	ServiceParameterName common.String
	RequestParameterName common.String
}

type DescribeApiDeployedInfo struct {
	StageName        common.String
	EffectiveVersion common.String
	DeployedStatus   common.String
}

type DescribeApiErrorCodeSample struct {
	Code        common.String
	Message     common.String
	Description common.String
}

type DescribeApiRequestConfig struct {
	RequestProtocol     common.String
	RequestHttpMethod   common.String
	RequestPath         common.String
	BodyFormat          common.String
	PostBodyDescription common.String
	RequestMode         common.String
}

type DescribeApiServiceConfig struct {
	ServiceProtocol     common.String
	ServiceAddress      common.String
	ServiceHttpMethod   common.String
	ServicePath         common.String
	ServiceTimeout      common.Integer
	Mock                common.String
	MockResult          common.String
	ServiceVpcEnable    bool
	VpcConfig           DescribeApiVpcConfig
	ContentTypeCatagory common.String
	ContentTypeValue    common.String
	AoneAppName         common.String
}

type DescribeApiOpenIdConnectConfig struct {
	OpenIdApiType    common.String
	IdTokenParamName common.String
	PublicKeyId      common.String
	PublicKey        common.String
}

type DescribeApiVpcConfig struct {
	VpcId      common.String
	InstanceId common.String
	Port       common.Integer
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
