package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshCdnDomainConfigsCacheRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Domains              string `position:"Query" name:"Domains"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r RefreshCdnDomainConfigsCacheRequest) Invoke(client *sdk.Client) (response *RefreshCdnDomainConfigsCacheResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RefreshCdnDomainConfigsCacheRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "RefreshCdnDomainConfigsCache", "", "")

	resp := struct {
		*responses.BaseResponse
		RefreshCdnDomainConfigsCacheResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RefreshCdnDomainConfigsCacheResponse

	err = client.DoAction(&req, &resp)
	return
}

type RefreshCdnDomainConfigsCacheResponse struct {
	RequestId     string
	SucessDomains RefreshCdnDomainConfigsCacheSucessDomainList
	FailedDomains RefreshCdnDomainConfigsCacheFailedDomainList
}

type RefreshCdnDomainConfigsCacheSucessDomainList []string

func (list *RefreshCdnDomainConfigsCacheSucessDomainList) UnmarshalJSON(data []byte) error {
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

type RefreshCdnDomainConfigsCacheFailedDomainList []string

func (list *RefreshCdnDomainConfigsCacheFailedDomainList) UnmarshalJSON(data []byte) error {
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
