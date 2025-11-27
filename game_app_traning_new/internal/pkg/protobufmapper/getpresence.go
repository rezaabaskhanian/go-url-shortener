package protobuf

import (
	"game_app/internal/contract/golang/presence"
	"game_app/internal/param"
)

func MapGetPresenceResponseToProtobuf(g param.GetPresenceResponse) *presence.GetPresenceResponse {
	r := &presence.GetPresenceResponse{}
	for _, item := range g.Items {
		r.Items = append(r.Items, &presence.GetPresenceItem{
			UserId:    uint64(item.UserID),
			Timestamp: item.TimeStamp,
		})
	}
	return r
}

func MapGetPresenceResponseFromProtobuf(g *presence.GetPresenceResponse) param.GetPresenceResponse {
	r := param.GetPresenceResponse{}
	for _, item := range g.Items {
		r.Items = append(r.Items, param.GetPresenceItem{
			UserID:    uint(item.UserId),
			TimeStamp: item.Timestamp,
		})
	}
	return r
}
