package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainOnlineUserNumberRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainOnlineUserNumberRequest) Invoke(client *sdk.Client) (resp *DescribeDomainOnlineUserNumberResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainOnlineUserNumber", "", "")
	resp = &DescribeDomainOnlineUserNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainOnlineUserNumberResponse struct {
	responses.BaseResponse
	RequestId                    string
	LiveStreamOnlineUserNumInfos DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo struct {
	Time       string
	UserNumber int64
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList []DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo)
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
