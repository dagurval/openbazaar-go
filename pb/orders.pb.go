// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderState int32

const (
	// Order has been funded and sent to the vendor but vendor has not yet responded
	OrderState_PENDING OrderState = 0
	// Waiting for the buyer to fund the payment address
	OrderState_AWAITING_PAYMENT OrderState = 1
	// Waiting for the customer to pick up the order (customer pickup option only)
	OrderState_AWAITING_PICKUP OrderState = 2
	// Order has been fully funded and we're waiting for the vendor to fulfill
	OrderState_AWAITING_FULFILLMENT OrderState = 3
	// Vendor has fulfilled part of the order
	OrderState_PARTIALLY_FULFILLED OrderState = 4
	// Vendor has fulfilled the order
	OrderState_FULFILLED OrderState = 5
	// Buyer has completed the order and left a review
	OrderState_COMPLETED OrderState = 6
	// Buyer canceled the order (offline order only)
	OrderState_CANCELED OrderState = 7
	// Vendor declined to confirm the order (offline order only)
	OrderState_DECLINED OrderState = 8
	// Vendor refunded the order
	OrderState_REFUNDED OrderState = 9
	// Contract is under active dispute
	OrderState_DISPUTED OrderState = 10
	// The moderator has resolved the dispute and we are waiting for the winning party to
	// accept the payout.
	OrderState_DECIDED OrderState = 11
	// The winning party has accepted the dispute and it is now complete. After the buyer
	// leaves a review the state should be set to COMPLETE.
	OrderState_RESOLVED OrderState = 12
	// Escrow has been released after waiting the timeout period. After the buyer
	// leaves a review the state should be set to COMPLETE.
	OrderState_PAYMENT_FINALIZED OrderState = 13
	// We screwed up and produced a order which didn't validate. This state is only used for offline orders. If a processing
	// error occured with an open connection between buyer and vendor the vendor just rejects the order on the spot neither party
	// commits the order to the database.
	OrderState_PROCESSING_ERROR OrderState = 14
)

var OrderState_name = map[int32]string{
	0:  "PENDING",
	1:  "AWAITING_PAYMENT",
	2:  "AWAITING_PICKUP",
	3:  "AWAITING_FULFILLMENT",
	4:  "PARTIALLY_FULFILLED",
	5:  "FULFILLED",
	6:  "COMPLETED",
	7:  "CANCELED",
	8:  "DECLINED",
	9:  "REFUNDED",
	10: "DISPUTED",
	11: "DECIDED",
	12: "RESOLVED",
	13: "PAYMENT_FINALIZED",
	14: "PROCESSING_ERROR",
}
var OrderState_value = map[string]int32{
	"PENDING":              0,
	"AWAITING_PAYMENT":     1,
	"AWAITING_PICKUP":      2,
	"AWAITING_FULFILLMENT": 3,
	"PARTIALLY_FULFILLED":  4,
	"FULFILLED":            5,
	"COMPLETED":            6,
	"CANCELED":             7,
	"DECLINED":             8,
	"REFUNDED":             9,
	"DISPUTED":             10,
	"DECIDED":              11,
	"RESOLVED":             12,
	"PAYMENT_FINALIZED":    13,
	"PROCESSING_ERROR":     14,
}

func (x OrderState) String() string {
	return proto.EnumName(OrderState_name, int32(x))
}
func (OrderState) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func init() {
	proto.RegisterEnum("OrderState", OrderState_name, OrderState_value)
}

func init() { proto.RegisterFile("orders.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0xa1, 0x94, 0xfe, 0x4c, 0x53, 0x18, 0xdc, 0x22, 0x38, 0x03, 0x0b, 0x36, 0x9c, 0xc0,
	0x78, 0x26, 0xd5, 0x08, 0xd7, 0xb1, 0x9c, 0x04, 0xd4, 0x6e, 0x2a, 0x2a, 0xba, 0x4e, 0x15, 0x72,
	0x15, 0xee, 0x8b, 0x26, 0xfc, 0x74, 0xf9, 0xbd, 0xef, 0x59, 0xf2, 0x1b, 0xc8, 0x9a, 0xf6, 0xe3,
	0xd0, 0x7e, 0x3e, 0x1e, 0xdb, 0xa6, 0x6b, 0x1e, 0xbe, 0x06, 0x00, 0x85, 0x06, 0x65, 0xf7, 0xde,
	0x1d, 0xcc, 0x0c, 0xc6, 0x91, 0x03, 0x49, 0x58, 0xe1, 0x99, 0x59, 0x02, 0xda, 0x37, 0x2b, 0x95,
	0x84, 0xd5, 0x2e, 0xda, 0xcd, 0x9a, 0x43, 0x85, 0xe7, 0x66, 0x01, 0xd7, 0xa7, 0x54, 0xdc, 0x4b,
	0x1d, 0x71, 0x60, 0xee, 0x61, 0xf9, 0x1f, 0xe6, 0xb5, 0xcf, 0xc5, 0xfb, 0xbe, 0x7e, 0x61, 0xee,
	0x60, 0x11, 0x6d, 0xaa, 0xc4, 0x7a, 0xbf, 0xf9, 0x53, 0x4c, 0x38, 0x34, 0x73, 0x98, 0x9e, 0xf0,
	0x52, 0xd1, 0x15, 0xeb, 0xe8, 0xb9, 0x62, 0xc2, 0x91, 0xc9, 0x60, 0xe2, 0x6c, 0x70, 0xac, 0x72,
	0xac, 0x44, 0xec, 0xbc, 0x04, 0x26, 0x9c, 0x28, 0x25, 0xce, 0xeb, 0x40, 0x4c, 0x38, 0xed, 0x9d,
	0x94, 0xb1, 0xd6, 0x77, 0xa0, 0x03, 0x88, 0x9d, 0xa8, 0x9a, 0xfd, 0x14, 0xcb, 0xc2, 0xbf, 0x32,
	0x61, 0x66, 0x6e, 0xe1, 0xe6, 0x77, 0xc5, 0x2e, 0x97, 0x60, 0xbd, 0x6c, 0x99, 0x70, 0xae, 0x2b,
	0x63, 0x2a, 0x1c, 0x97, 0xa5, 0x7e, 0x9e, 0x53, 0x2a, 0x12, 0x5e, 0x3d, 0x0f, 0xb7, 0x83, 0xe3,
	0x7e, 0x3f, 0xea, 0x8f, 0xf4, 0xf4, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xa4, 0x45, 0x88, 0x34,
	0x01, 0x00, 0x00,
}
