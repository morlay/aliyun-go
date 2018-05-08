package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamHlsOnlineUserNumByDomainRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	HlsSwitch     string `position:"Query" name:"HlsSwitch"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (req *DescribeLiveStreamHlsOnlineUserNumByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamHlsOnlineUserNumByDomain", "", "")
	resp = &DescribeLiveStreamHlsOnlineUserNumByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamHlsOnlineUserNumByDomainResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TotalUserNumber common.Long
	Count           common.Long
	PageNumber      common.Long
	PageSize        common.Long
	OnlineUserInfo  DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  common.String
	UserNumber common.Long
	Time       common.String
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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
