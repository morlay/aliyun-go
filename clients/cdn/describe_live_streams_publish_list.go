package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsPublishListRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (r DescribeLiveStreamsPublishListRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsPublishListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsPublishListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamsPublishList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsPublishListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsPublishListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsPublishListResponse struct {
	RequestId   string
	PublishInfo DescribeLiveStreamsPublishListLiveStreamPublishInfoList
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfo struct {
	DomainName   string
	AppName      string
	StreamName   string
	StreamUrl    string
	PublishTime  string
	StopTime     string
	PublishUrl   string
	ClientAddr   string
	EdgeNodeAddr string
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfoList []DescribeLiveStreamsPublishListLiveStreamPublishInfo

func (list *DescribeLiveStreamsPublishListLiveStreamPublishInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsPublishListLiveStreamPublishInfo)
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
