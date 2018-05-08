package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
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
	ConfigId     common.String
	CacheType    common.String
	CacheContent common.String
	TTL          common.String
	Weight       common.String
	Status       common.String
}

type DescribeDomainConfigsHttpErrorPageConfig struct {
	ConfigId  common.String
	ErrorCode common.String
	PageUrl   common.String
	Status    common.String
}

type DescribeDomainConfigsHttpHeaderConfig struct {
	ConfigId    common.String
	HeaderKey   common.String
	HeaderValue common.String
	Status      common.String
}

type DescribeDomainConfigsDynamicConfig struct {
	ConfigId            common.String
	DynamicOrigin       common.String
	StaticType          common.String
	StaticUri           common.String
	StaticPath          common.String
	DynamicCacheControl common.String
	Status              common.String
}

type DescribeDomainConfigsReqHeaderConfig struct {
	ConfigId common.String
	Key      common.String
	Value    common.String
	Status   common.String
}

type DescribeDomainConfigsSetVarsConfig struct {
	ConfigId common.String
	VarName  common.String
	VarValue common.String
	Status   common.String
}

type DescribeDomainConfigsCcConfig struct {
	ConfigId common.String
	Enable   common.String
	AllowIps common.String
	BlockIps common.String
	Status   common.String
}

type DescribeDomainConfigsErrorPageConfig struct {
	ConfigId      common.String
	ErrorCode     common.String
	PageType      common.String
	CustomPageUrl common.String
	Status        common.String
}

type DescribeDomainConfigsOptimizeConfig struct {
	ConfigId common.String
	Enable   common.String
	Status   common.String
}

type DescribeDomainConfigsPageCompressConfig struct {
	ConfigId common.String
	Enable   common.String
	Status   common.String
}

type DescribeDomainConfigsIgnoreQueryStringConfig struct {
	ConfigId    common.String
	HashKeyArgs common.String
	Enable      common.String
	Status      common.String
}

type DescribeDomainConfigsRangeConfig struct {
	ConfigId common.String
	Enable   common.String
	Status   common.String
}

type DescribeDomainConfigsRefererConfig struct {
	ConfigId   common.String
	ReferType  common.String
	ReferList  common.String
	AllowEmpty common.String
	DisableAst common.String
	Status     common.String
}

type DescribeDomainConfigsReqAuthConfig struct {
	ConfigId         common.String
	AuthType         common.String
	Key1             common.String
	Key2             common.String
	Status           common.String
	AliAuthWhiteList common.String
	AuthM3u8         common.String
	AuthAddr         common.String
	AuthRemoteDesc   common.String
	TimeOut          common.String
}

type DescribeDomainConfigsSrcHostConfig struct {
	ConfigId   common.String
	DomainName common.String
	Status     common.String
}

type DescribeDomainConfigsVideoSeekConfig struct {
	ConfigId common.String
	Enable   common.String
	Status   common.String
}

type DescribeDomainConfigsWafConfig struct {
	ConfigId common.String
	Enable   common.String
	Status   common.String
}

type DescribeDomainConfigsNotifyUrlConfig struct {
	Enable    common.String
	NotifyUrl common.String
}

type DescribeDomainConfigsRedirectTypeConfig struct {
	RedirectType common.String
}

type DescribeDomainConfigsForwardSchemeConfig struct {
	ConfigId         common.String
	Enable           common.String
	SchemeOrigin     common.String
	SchemeOriginPort common.String
	Status           common.String
}

type DescribeDomainConfigsRemoveQueryStringConfig struct {
	AliRemoveArgs common.String
	ConfigId      common.String
	Status        common.String
}

type DescribeDomainConfigsL2OssKeyConfig struct {
	PrivateOssAuth common.String
	ConfigId       common.String
	Status         common.String
}

type DescribeDomainConfigsMacServiceConfig struct {
	AppList       common.String
	Enabled       common.String
	ProcessResult common.String
	ConfigId      common.String
	Status        common.String
}

type DescribeDomainConfigsGreenManagerConfig struct {
	Enabled  common.String
	ConfigId common.String
	Status   common.String
}

type DescribeDomainConfigsHttpsOptionConfig struct {
	Http2    common.String
	ConfigId common.String
	Status   common.String
}

type DescribeDomainConfigsAliBusinessConfig struct {
	AliBusinessTable common.String
	AliBusinessType  common.String
	ConfigId         common.String
	Status           common.String
}

type DescribeDomainConfigsIpAllowListConfig struct {
	ConfigId  common.String
	IpList    common.String
	IpAclXfwd common.String
	Status    common.String
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
