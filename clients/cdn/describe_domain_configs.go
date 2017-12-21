package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainConfigsRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	ConfigList    string `position:"Query" name:"ConfigList"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainConfigsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainConfigsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainConfigs", "", "")
	resp = &DescribeDomainConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainConfigsResponse struct {
	responses.BaseResponse
	RequestId     string
	DomainConfigs DescribeDomainConfigsDomainConfigs
}

type DescribeDomainConfigsDomainConfigs struct {
	CacheExpiredConfigs     DescribeDomainConfigsCacheExpiredConfigList
	HttpErrorPageConfigs    DescribeDomainConfigsHttpErrorPageConfigList
	HttpHeaderConfigs       DescribeDomainConfigsHttpHeaderConfigList
	DynamicConfigs          DescribeDomainConfigsDynamicConfigList
	ReqHeaderConfigs        DescribeDomainConfigsReqHeaderConfigList
	SetVarsConfigs          DescribeDomainConfigsSetVarsConfigList
	CcConfig                DescribeDomainConfigsCcConfig
	ErrorPageConfig         DescribeDomainConfigsErrorPageConfig
	OptimizeConfig          DescribeDomainConfigsOptimizeConfig
	PageCompressConfig      DescribeDomainConfigsPageCompressConfig
	IgnoreQueryStringConfig DescribeDomainConfigsIgnoreQueryStringConfig
	RangeConfig             DescribeDomainConfigsRangeConfig
	RefererConfig           DescribeDomainConfigsRefererConfig
	ReqAuthConfig           DescribeDomainConfigsReqAuthConfig
	SrcHostConfig           DescribeDomainConfigsSrcHostConfig
	VideoSeekConfig         DescribeDomainConfigsVideoSeekConfig
	WafConfig               DescribeDomainConfigsWafConfig
	NotifyUrlConfig         DescribeDomainConfigsNotifyUrlConfig
	RedirectTypeConfig      DescribeDomainConfigsRedirectTypeConfig
	ForwardSchemeConfig     DescribeDomainConfigsForwardSchemeConfig
	RemoveQueryStringConfig DescribeDomainConfigsRemoveQueryStringConfig
	L2OssKeyConfig          DescribeDomainConfigsL2OssKeyConfig
	MacServiceConfig        DescribeDomainConfigsMacServiceConfig
	GreenManagerConfig      DescribeDomainConfigsGreenManagerConfig
	HttpsOptionConfig       DescribeDomainConfigsHttpsOptionConfig
	AliBusinessConfig       DescribeDomainConfigsAliBusinessConfig
	IpAllowListConfig       DescribeDomainConfigsIpAllowListConfig
}

type DescribeDomainConfigsCacheExpiredConfig struct {
	ConfigId     string
	CacheType    string
	CacheContent string
	TTL          string
	Weight       string
	Status       string
}

type DescribeDomainConfigsHttpErrorPageConfig struct {
	ConfigId  string
	ErrorCode string
	PageUrl   string
	Status    string
}

type DescribeDomainConfigsHttpHeaderConfig struct {
	ConfigId    string
	HeaderKey   string
	HeaderValue string
	Status      string
}

type DescribeDomainConfigsDynamicConfig struct {
	ConfigId            string
	DynamicOrigin       string
	StaticType          string
	StaticUri           string
	StaticPath          string
	DynamicCacheControl string
	Status              string
}

type DescribeDomainConfigsReqHeaderConfig struct {
	ConfigId string
	Key      string
	Value    string
	Status   string
}

type DescribeDomainConfigsSetVarsConfig struct {
	ConfigId string
	VarName  string
	VarValue string
	Status   string
}

type DescribeDomainConfigsCcConfig struct {
	ConfigId string
	Enable   string
	AllowIps string
	BlockIps string
	Status   string
}

type DescribeDomainConfigsErrorPageConfig struct {
	ConfigId      string
	ErrorCode     string
	PageType      string
	CustomPageUrl string
	Status        string
}

type DescribeDomainConfigsOptimizeConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsPageCompressConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsIgnoreQueryStringConfig struct {
	ConfigId    string
	HashKeyArgs string
	Enable      string
	Status      string
}

type DescribeDomainConfigsRangeConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsRefererConfig struct {
	ConfigId   string
	ReferType  string
	ReferList  string
	AllowEmpty string
	DisableAst string
	Status     string
}

type DescribeDomainConfigsReqAuthConfig struct {
	ConfigId         string
	AuthType         string
	Key1             string
	Key2             string
	Status           string
	AliAuthWhiteList string
	AuthM3u8         string
	AuthAddr         string
	AuthRemoteDesc   string
	TimeOut          string
}

type DescribeDomainConfigsSrcHostConfig struct {
	ConfigId   string
	DomainName string
	Status     string
}

type DescribeDomainConfigsVideoSeekConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsWafConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsNotifyUrlConfig struct {
	Enable    string
	NotifyUrl string
}

type DescribeDomainConfigsRedirectTypeConfig struct {
	RedirectType string
}

type DescribeDomainConfigsForwardSchemeConfig struct {
	ConfigId         string
	Enable           string
	SchemeOrigin     string
	SchemeOriginPort string
	Status           string
}

type DescribeDomainConfigsRemoveQueryStringConfig struct {
	AliRemoveArgs string
	ConfigId      string
	Status        string
}

type DescribeDomainConfigsL2OssKeyConfig struct {
	PrivateOssAuth string
	ConfigId       string
	Status         string
}

type DescribeDomainConfigsMacServiceConfig struct {
	AppList       string
	Enabled       string
	ProcessResult string
	ConfigId      string
	Status        string
}

type DescribeDomainConfigsGreenManagerConfig struct {
	Enabled  string
	ConfigId string
	Status   string
}

type DescribeDomainConfigsHttpsOptionConfig struct {
	Http2    string
	ConfigId string
	Status   string
}

type DescribeDomainConfigsAliBusinessConfig struct {
	AliBusinessTable string
	AliBusinessType  string
	ConfigId         string
	Status           string
}

type DescribeDomainConfigsIpAllowListConfig struct {
	ConfigId  string
	IpList    string
	IpAclXfwd string
	Status    string
}

type DescribeDomainConfigsCacheExpiredConfigList []DescribeDomainConfigsCacheExpiredConfig

func (list *DescribeDomainConfigsCacheExpiredConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsCacheExpiredConfig)
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

type DescribeDomainConfigsHttpErrorPageConfigList []DescribeDomainConfigsHttpErrorPageConfig

func (list *DescribeDomainConfigsHttpErrorPageConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpErrorPageConfig)
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

type DescribeDomainConfigsHttpHeaderConfigList []DescribeDomainConfigsHttpHeaderConfig

func (list *DescribeDomainConfigsHttpHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpHeaderConfig)
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

type DescribeDomainConfigsDynamicConfigList []DescribeDomainConfigsDynamicConfig

func (list *DescribeDomainConfigsDynamicConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsDynamicConfig)
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

type DescribeDomainConfigsReqHeaderConfigList []DescribeDomainConfigsReqHeaderConfig

func (list *DescribeDomainConfigsReqHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsReqHeaderConfig)
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

type DescribeDomainConfigsSetVarsConfigList []DescribeDomainConfigsSetVarsConfig

func (list *DescribeDomainConfigsSetVarsConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsSetVarsConfig)
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
