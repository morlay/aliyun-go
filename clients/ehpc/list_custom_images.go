package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListCustomImagesRequest struct {
	requests.RpcRequest
	BaseOsTag       string `position:"Query" name:"BaseOsTag"`
	ImageOwnerAlias string `position:"Query" name:"ImageOwnerAlias"`
}

func (req *ListCustomImagesRequest) Invoke(client *sdk.Client) (resp *ListCustomImagesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "ListCustomImages", "ehs", "")
	resp = &ListCustomImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCustomImagesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Images    ListCustomImagesImageInfoList
}

type ListCustomImagesImageInfo struct {
	Uid               common.String
	ImageId           common.String
	ImageName         common.String
	ImageOwnerAlias   common.String
	Description       common.String
	Status            common.String
	ProductCode       common.String
	SkuCode           common.String
	PricingCycle      common.String
	PostInstallScript common.String
	BaseOsTag         ListCustomImagesBaseOsTag
}

type ListCustomImagesBaseOsTag struct {
	OsTag        common.String
	Platform     common.String
	Version      common.String
	Architecture common.String
}

type ListCustomImagesImageInfoList []ListCustomImagesImageInfo

func (list *ListCustomImagesImageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCustomImagesImageInfo)
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
