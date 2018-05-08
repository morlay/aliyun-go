package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListImagesRequest struct {
	requests.RpcRequest
}

func (req *ListImagesRequest) Invoke(client *sdk.Client) (resp *ListImagesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListImages", "ehs", "")
	resp = &ListImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListImagesResponse struct {
	responses.BaseResponse
	RequestId common.String
	OsTags    ListImagesOsInfoList
}

type ListImagesOsInfo struct {
	OsTag        common.String
	Platform     common.String
	Version      common.String
	Architecture common.String
}

type ListImagesOsInfoList []ListImagesOsInfo

func (list *ListImagesOsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListImagesOsInfo)
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
