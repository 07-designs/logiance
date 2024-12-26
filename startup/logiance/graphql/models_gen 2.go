// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Shridhar2104/logilo/graphql/models"
)

type AccountInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AccountShipmentsInput struct {
	AccountID string `json:"accountId"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}

type AccountShipmentsResponse struct {
	Success   bool            `json:"success"`
	Shipments []*ShipmentInfo `json:"shipments,omitempty"`
	Error     *string         `json:"error,omitempty"`
}

type Accounts struct {
	Orders []*models.Order `json:"orders"`
}

type AddressInput struct {
	Name         string  `json:"name"`
	CompanyName  *string `json:"company_name,omitempty"`
	Phone        string  `json:"phone"`
	Email        *string `json:"email,omitempty"`
	AddressLine1 string  `json:"address_line1"`
	AddressLine2 *string `json:"address_line2,omitempty"`
	Landmark     *string `json:"landmark,omitempty"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      *string `json:"country,omitempty"`
	Pincode      string  `json:"pincode"`
	Gstin        *string `json:"gstin,omitempty"`
}

type AvailabilityInput struct {
	OriginPincode      int          `json:"origin_pincode"`
	DestinationPincode int          `json:"destination_pincode"`
	Weight             float64      `json:"weight"`
	PaymentMode        *PaymentMode `json:"payment_mode,omitempty"`
}

type BankAccount struct {
	UserID          string `json:"userId"`
	AccountNumber   string `json:"accountNumber"`
	BeneficiaryName string `json:"beneficiaryName"`
	IfscCode        string `json:"ifscCode"`
	BankName        string `json:"bankName"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type BankAccountInput struct {
	AccountNumber   string `json:"accountNumber"`
	BeneficiaryName string `json:"beneficiaryName"`
	IfscCode        string `json:"ifscCode"`
	BankName        string `json:"bankName"`
}

type CreateShipmentInput struct {
	AccountID         string            `json:"accountId"`
	CourierCode       string            `json:"courier_code"`
	OrderNumber       string            `json:"order_number"`
	PaymentType       PaymentType       `json:"payment_type"`
	PackageWeight     float64           `json:"package_weight"`
	PackageLength     float64           `json:"package_length"`
	PackageBreadth    float64           `json:"package_breadth"`
	PackageHeight     float64           `json:"package_height"`
	OrderAmount       float64           `json:"order_amount"`
	CollectableAmount float64           `json:"collectable_amount"`
	Consignee         *AddressInput     `json:"consignee"`
	Pickup            *AddressInput     `json:"pickup"`
	Items             []*OrderItemInput `json:"items"`
	AutoPickup        *bool             `json:"auto_pickup,omitempty"`
	ReturnInfo        *ReturnInfoInput  `json:"return_info,omitempty"`
}

type Customer struct {
	ID        *string `json:"id,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

type DeductBalanceInput struct {
	AccountID string  `json:"accountId"`
	Amount    float64 `json:"amount"`
	OrderID   string  `json:"orderId"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetWalletDetailsInput struct {
	AccountID string `json:"accountId"`
}

type Mutation struct {
}

type OrderConnection struct {
	Edges      []*OrderEdge `json:"edges"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type OrderEdge struct {
	Node   *models.Order `json:"node"`
	Cursor string        `json:"cursor"`
}

type OrderFilter struct {
	CreatedAtStart    *string  `json:"createdAtStart,omitempty"`
	CreatedAtEnd      *string  `json:"createdAtEnd,omitempty"`
	FinancialStatus   *string  `json:"financialStatus,omitempty"`
	FulfillmentStatus *string  `json:"fulfillmentStatus,omitempty"`
	MinTotalPrice     *float64 `json:"minTotalPrice,omitempty"`
	MaxTotalPrice     *float64 `json:"maxTotalPrice,omitempty"`
	SearchTerm        *string  `json:"searchTerm,omitempty"`
}

type OrderInput struct {
	AccountID string                `json:"accountId"`
	LineItems []*OrderLineItemInput `json:"lineItems"`
}

type OrderItemInput struct {
	Sku          string   `json:"sku"`
	Name         string   `json:"name"`
	Quantity     int      `json:"quantity"`
	Price        float64  `json:"price"`
	HsnCode      *string  `json:"hsn_code,omitempty"`
	Category     *string  `json:"category,omitempty"`
	ActualWeight *float64 `json:"actual_weight,omitempty"`
}

type OrderLineItemInput struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type OrderPaginationInput struct {
	Page     *int         `json:"page,omitempty"`
	PageSize *int         `json:"pageSize,omitempty"`
	Filter   *OrderFilter `json:"filter,omitempty"`
	Sort     *OrderSort   `json:"sort,omitempty"`
}

type OrderSort struct {
	Field     OrderSortField `json:"field"`
	Direction SortDirection  `json:"direction"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
	TotalPages      int     `json:"totalPages"`
	CurrentPage     int     `json:"currentPage"`
}

type PaginationInput struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

type Query struct {
}

type RechargeWalletInput struct {
	AccountID string  `json:"accountId"`
	Amount    float64 `json:"amount"`
}

type ReturnInfoInput struct {
	Address       *AddressInput `json:"address"`
	AwbNumber     *string       `json:"awb_number,omitempty"`
	ReturnReason  *string       `json:"return_reason,omitempty"`
	ReturnComment *string       `json:"return_comment,omitempty"`
}

type ShipmentInfo struct {
	OrderNumber string  `json:"orderNumber"`
	TrackingID  string  `json:"trackingId"`
	CourierAwb  string  `json:"courierAwb"`
	Status      string  `json:"status"`
	Label       *string `json:"label,omitempty"`
	CreatedAt   string  `json:"createdAt"`
}

type ShipmentResponse struct {
	Success    bool    `json:"success"`
	TrackingID *string `json:"tracking_id,omitempty"`
	CourierAwb *string `json:"courier_awb,omitempty"`
	Label      *string `json:"label,omitempty"`
	Error      *string `json:"error,omitempty"`
}

type ShippingRateInput struct {
	OriginPincode      int         `json:"origin_pincode"`
	DestinationPincode int         `json:"destination_pincode"`
	Weight             float64     `json:"weight"`
	Length             *float64    `json:"length,omitempty"`
	Width              *float64    `json:"width,omitempty"`
	Height             *float64    `json:"height,omitempty"`
	PaymentMode        PaymentMode `json:"payment_mode"`
	CollectableAmount  float64     `json:"collectable_amount"`
	CourierCodes       []string    `json:"courier_codes,omitempty"`
}

type TrackingEvent struct {
	Status      string `json:"status"`
	Location    string `json:"location"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
}

type TrackingInput struct {
	CourierCode string `json:"courier_code"`
	TrackingID  string `json:"tracking_id"`
}

type TrackingResponse struct {
	Success bool             `json:"success"`
	Events  []*TrackingEvent `json:"events,omitempty"`
	Error   *string          `json:"error,omitempty"`
}

type WalletDetails struct {
	AccountID   string        `json:"accountId"`
	Balance     *float64      `json:"balance,omitempty"`
	Currency    *string       `json:"currency,omitempty"`
	Status      *WalletStatus `json:"status,omitempty"`
	LastUpdated *string       `json:"lastUpdated,omitempty"`
}

type WalletDetailsResponse struct {
	WalletDetails *WalletDetails `json:"walletDetails,omitempty"`
	Errors        []*Error       `json:"errors,omitempty"`
}

type WalletOperationResponse struct {
	NewBalance float64  `json:"newBalance"`
	Errors     []*Error `json:"errors,omitempty"`
}

type OrderSortField string

const (
	OrderSortFieldCreatedAt   OrderSortField = "CREATED_AT"
	OrderSortFieldUpdatedAt   OrderSortField = "UPDATED_AT"
	OrderSortFieldOrderNumber OrderSortField = "ORDER_NUMBER"
	OrderSortFieldTotalPrice  OrderSortField = "TOTAL_PRICE"
)

var AllOrderSortField = []OrderSortField{
	OrderSortFieldCreatedAt,
	OrderSortFieldUpdatedAt,
	OrderSortFieldOrderNumber,
	OrderSortFieldTotalPrice,
}

func (e OrderSortField) IsValid() bool {
	switch e {
	case OrderSortFieldCreatedAt, OrderSortFieldUpdatedAt, OrderSortFieldOrderNumber, OrderSortFieldTotalPrice:
		return true
	}
	return false
}

func (e OrderSortField) String() string {
	return string(e)
}

func (e *OrderSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderSortField", str)
	}
	return nil
}

func (e OrderSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentMode string

const (
	PaymentModeCod     PaymentMode = "COD"
	PaymentModePrepaid PaymentMode = "PREPAID"
)

var AllPaymentMode = []PaymentMode{
	PaymentModeCod,
	PaymentModePrepaid,
}

func (e PaymentMode) IsValid() bool {
	switch e {
	case PaymentModeCod, PaymentModePrepaid:
		return true
	}
	return false
}

func (e PaymentMode) String() string {
	return string(e)
}

func (e *PaymentMode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentMode", str)
	}
	return nil
}

func (e PaymentMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentType string

const (
	PaymentTypeCod     PaymentType = "COD"
	PaymentTypePrepaid PaymentType = "PREPAID"
)

var AllPaymentType = []PaymentType{
	PaymentTypeCod,
	PaymentTypePrepaid,
}

func (e PaymentType) IsValid() bool {
	switch e {
	case PaymentTypeCod, PaymentTypePrepaid:
		return true
	}
	return false
}

func (e PaymentType) String() string {
	return string(e)
}

func (e *PaymentType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentType", str)
	}
	return nil
}

func (e PaymentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WalletStatus string

const (
	WalletStatusActive    WalletStatus = "ACTIVE"
	WalletStatusInactive  WalletStatus = "INACTIVE"
	WalletStatusSuspended WalletStatus = "SUSPENDED"
	WalletStatusClosed    WalletStatus = "CLOSED"
)

var AllWalletStatus = []WalletStatus{
	WalletStatusActive,
	WalletStatusInactive,
	WalletStatusSuspended,
	WalletStatusClosed,
}

func (e WalletStatus) IsValid() bool {
	switch e {
	case WalletStatusActive, WalletStatusInactive, WalletStatusSuspended, WalletStatusClosed:
		return true
	}
	return false
}

func (e WalletStatus) String() string {
	return string(e)
}

func (e *WalletStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WalletStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WalletStatus", str)
	}
	return nil
}

func (e WalletStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}