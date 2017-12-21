package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsOnlineListRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamsOnlineListRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsOnlineListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsOnlineListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsOnlineList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsOnlineListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsOnlineListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsOnlineListResponse struct {
	RequestId  string
	OnlineInfo DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfo struct {
	DomainName  string
	AppName     string
	StreamName  string
	PublishTime string
	PublishUrl  string
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList []DescribeLiveStreamsOnlineListLiveStreamOnlineInfo

func (list *DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsOnlineListLiveStreamOnlineInfo)
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
