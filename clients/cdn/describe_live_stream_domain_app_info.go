package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamDomainAppInfoRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	AppDomain     string `position:"Query" name:"AppDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamDomainAppInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamDomainAppInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamDomainAppInfo", "", "")
	resp = &DescribeLiveStreamDomainAppInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamDomainAppInfoResponse struct {
	responses.BaseResponse
	RequestId     common.String
	DomainAppList DescribeLiveStreamDomainAppInfoDomainAppInfoList
}

type DescribeLiveStreamDomainAppInfoDomainAppInfo struct {
	AppDomain    common.String
	AppId        common.String
	AppKey       common.String
	AppOssBucket common.String
	AppOssHost   common.String
	AppOwnerId   common.String
	AppSecret    common.String
	UpdateTime   common.String
}

type DescribeLiveStreamDomainAppInfoDomainAppInfoList []DescribeLiveStreamDomainAppInfoDomainAppInfo

func (list *DescribeLiveStreamDomainAppInfoDomainAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamDomainAppInfoDomainAppInfo)
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
