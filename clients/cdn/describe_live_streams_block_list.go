package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsBlockListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamsBlockListRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsBlockListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsBlockListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamsBlockList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsBlockListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsBlockListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsBlockListResponse struct {
	RequestId  string
	DomainName string
	StreamUrls DescribeLiveStreamsBlockListStreamUrlList
}

type DescribeLiveStreamsBlockListStreamUrlList []string

func (list *DescribeLiveStreamsBlockListStreamUrlList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
