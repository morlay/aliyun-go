package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteRouteEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                            `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string                           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	NextHopId            string                           `position:"Query" name:"NextHopId"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	NextHopLists         *DeleteRouteEntryNextHopListList `position:"Query" type:"Repeated" name:"NextHopList"`
	RouteTableId         string                           `position:"Query" name:"RouteTableId"`
}

func (req *DeleteRouteEntryRequest) Invoke(client *sdk.Client) (resp *DeleteRouteEntryResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouteEntry", "vpc", "")
	resp = &DeleteRouteEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
}

type DeleteRouteEntryResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type DeleteRouteEntryNextHopListList []DeleteRouteEntryNextHopList

func (list *DeleteRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteRouteEntryNextHopList)
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
