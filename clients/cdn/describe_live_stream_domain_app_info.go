package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamDomainAppInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	AppDomain     string `position:"Query" name:"AppDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamDomainAppInfoRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamDomainAppInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamDomainAppInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamDomainAppInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamDomainAppInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamDomainAppInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamDomainAppInfoResponse struct {
	RequestId     string
	DomainAppList DescribeLiveStreamDomainAppInfoDomainAppInfoList
}

type DescribeLiveStreamDomainAppInfoDomainAppInfo struct {
	AppDomain    string
	AppId        string
	AppKey       string
	AppOssBucket string
	AppOssHost   string
	AppOwnerId   string
	AppSecret    string
	UpdateTime   string
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
