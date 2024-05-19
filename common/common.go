package common

import (
	"time"
)

type OrderStatusType uint

const (
	OrderStatusCreated   OrderStatusType = 100
	OrderStatusAccepted                  = 200
	OrderStatusReady                     = 300
	OrderStatusSent                      = 400
	OrderStatusDelivered                 = 500
	OrderStatusCanceled                  = 900
)

type VersionInfo struct {
	Version    string
	CommitHash string
	BuildTime  string
}

var (
	Version    string
	CommitHash string
	BuildTime  string

	lastRun time.Time
)

func GetVersion() *VersionInfo {
	return &VersionInfo{
		Version:    Version,
		CommitHash: CommitHash,
		BuildTime:  BuildTime,
	}
}

func SetLastRun(t time.Time) {
	lastRun = t
}

func GetLastRun() time.Time {
	return lastRun
}

func (o OrderStatusType) ToInt() uint {
	return uint(o)
}

func (o OrderStatusType) ToString() string {
	switch o {
	case OrderStatusCreated:
		return "created"
	case OrderStatusAccepted:
		return "accepted"
	case OrderStatusReady:
		return "ready"
	case OrderStatusSent:
		return "sent"
	case OrderStatusDelivered:
		return "delivered"
	case OrderStatusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

func OrderStatus(i uint) OrderStatusType {
	return OrderStatusType(i)
}
