package ozon

import (
	"time"
)

const (
	testTimeout = 5 * time.Second
)

type Order string

const (
	Ascending  Order = "ASC"
	Descending Order = "DESC"
)

type GetAnalyticsDataFilterOperation string

const (
	Equal        GetAnalyticsDataFilterOperation = "EQ"
	Greater      GetAnalyticsDataFilterOperation = "GT"
	GreaterEqual GetAnalyticsDataFilterOperation = "GTE"
	Lesser       GetAnalyticsDataFilterOperation = "LT"
	LesserEqual  GetAnalyticsDataFilterOperation = "LTE"
)

type GetAnalyticsDataFilterMetric string

const (
	UnknownMetric       GetAnalyticsDataFilterMetric = "unknown_metric"
	HitsViewSearch      GetAnalyticsDataFilterMetric = "hits_view_search"
	HistViewPDP         GetAnalyticsDataFilterMetric = "hits_view_pdp"
	HitsView            GetAnalyticsDataFilterMetric = "hist_view"
	HitsToCartSearch    GetAnalyticsDataFilterMetric = "hits_tocart_search"
	HitsToCartPDP       GetAnalyticsDataFilterMetric = "hits_tocart_pdp"
	SessionViewSearch   GetAnalyticsDataFilterMetric = "session_view_search"
	SessionViewPDP      GetAnalyticsDataFilterMetric = "session_view_pdp"
	SessionView         GetAnalyticsDataFilterMetric = "session_view"
	ConvToCartSearch    GetAnalyticsDataFilterMetric = "conv_tocart_search"
	ConvToCartPDP       GetAnalyticsDataFilterMetric = "conv_tocart_pdp"
	ConvToCart          GetAnalyticsDataFilterMetric = "conv_tocart"
	Revenue             GetAnalyticsDataFilterMetric = "revenue"
	ReturnsMetric       GetAnalyticsDataFilterMetric = "returns"
	CancellationsMetric GetAnalyticsDataFilterMetric = "cancellations"
	OrderedUnits        GetAnalyticsDataFilterMetric = "ordered_units"
	DeliveredUnits      GetAnalyticsDataFilterMetric = "delivered_units"
	PositionCategory    GetAnalyticsDataFilterMetric = "position_category"
)

type WarehouseType string

const (
	// Ozon warehouses with Fresh delivery
	ExpressDarkStore WarehouseType = "EXPRESS_DARK_STORE"

	// Ozon warehouses without Fresh delivery
	NotExressDarkStore WarehouseType = "NOT_EXPRESS_DARK_STORE"

	// All Ozon warehouses
	ALLWarehouseType WarehouseType = "ALL"
)

type Language string

const (
	Default Language = "DEFAULT"
	Russian Language = "RU"
	English Language = "EN"
	Turkish Language = "TR"
	Chinese Language = "ZH_HANS"
)

type AttributeType string

const (
	All      AttributeType = "ALL"
	Required AttributeType = "REQUIRED"
	Optional AttributeType = "OPTIONAL"
)

type ListDiscountRequestsStatus string

const (
	New            ListDiscountRequestsStatus = "NEW"
	Seen           ListDiscountRequestsStatus = "SEEN"
	Approved       ListDiscountRequestsStatus = "APPROVED"
	PartlyApproved ListDiscountRequestsStatus = "PARTLY_APPROVED"
	Declined       ListDiscountRequestsStatus = "DECLINED"
	AutoDeclined   ListDiscountRequestsStatus = "AUTO_DECLINED"
	DeclinedByUser ListDiscountRequestsStatus = "DECLINED_BY_USER"
	Coupon         ListDiscountRequestsStatus = "COUPON"
	Purchased      ListDiscountRequestsStatus = "PURCHASED"
)

type WorkingDay int

const (
	Mon WorkingDay = 1
	Tue WorkingDay = 2
	Wed WorkingDay = 3
	Thu WorkingDay = 4
	Fri WorkingDay = 5
	Sat WorkingDay = 6
	Sun WorkingDay = 7
)

type GetAnalyticsDataDimension string

const (
	UnknownDimension   GetAnalyticsDataDimension = "unknownDimension"
	SKUDimension       GetAnalyticsDataDimension = "sku"
	SPUDimension       GetAnalyticsDataDimension = "spu"
	DayDimension       GetAnalyticsDataDimension = "day"
	WeekDimension      GetAnalyticsDataDimension = "week"
	MonthDimension     GetAnalyticsDataDimension = "month"
	YearDimension      GetAnalyticsDataDimension = "year"
	Category1Dimension GetAnalyticsDataDimension = "category1"
	Category2Dimension GetAnalyticsDataDimension = "category2"
	Category3Dimension GetAnalyticsDataDimension = "category3"
	Category4Dimension GetAnalyticsDataDimension = "category4"
	BrandDimension     GetAnalyticsDataDimension = "brand"
	ModelIDDimension   GetAnalyticsDataDimension = "modelID"
)

type SupplyRequestState string

const (
	// filling in the data
	DATA_FILLING SupplyRequestState = "DATA_FILLING"

	// ready for shipment
	ReadyToSupply SupplyRequestState = "READY_TO_SUPPLY"

	// accepted at the shipping point
	AcceptedAtSupplyWarehouse SupplyRequestState = "ACCEPTED_AT_SUPPLY_WAREHOUSE"

	// on the way
	InTransit SupplyRequestState = "IN_TRANSIT"

	// acceptance at the warehouse
	AcceptanceAtStorageWarehouse SupplyRequestState = "ACCEPTANCE_AT_STORAGE_WAREHOUSE"

	// acts being approved
	ReportsConfirmationAwaiting SupplyRequestState = "REPORTS_CONFIRMATION_AWAITING"

	// dispute
	ReportRejected SupplyRequestState = "REPORT_REJECTED"

	// completed
	Completed SupplyRequestState = "COMPLETED"

	// refused acceptance
	RejectedAtSupplyWarehouse SupplyRequestState = "REJECTED_AT_SUPPLY_WAREHOUSE"

	// cancelled
	Cancelled SupplyRequestState = "CANCELLED"

	// overdue
	Overdue SupplyRequestState = "OVERDUE"
)

type ShipmentStatus string

const (
	// acceptance is in progress
	AcceptanceInProgress ShipmentStatus = "acceptance_in_progress"

	// arbitration
	Arbitration ShipmentStatus = "arbitration"

	// awaiting confirmation
	AwaitingApprove ShipmentStatus = "awaiting_approve"

	// awaiting shipping
	AwaitingDeliver ShipmentStatus = "awaiting_deliver"

	// awaiting packaging
	AwaitingPackaging ShipmentStatus = "awaiting_packaging"

	// created
	AwaitingVerification ShipmentStatus = "awaiting_verification"

	// cancelled
	CancelledSubstatus ShipmentStatus = "cancelled"

	// delivered
	Delivered ShipmentStatus = "delivered"

	// delivery is in progress
	Delivering ShipmentStatus = "delivering"

	// picked up by driver
	DriverPickup ShipmentStatus = "driver_pickup"

	// not accepted at the sorting center
	NotAccepted ShipmentStatus = "not_accepted"

	// sent by the seller
	SentBySeller ShipmentStatus = "sent_by_seller"
)

type ShipmentSubstatus string

const (
	// acceptance in progress
	PostingAcceptanceInProgress ShipmentStatus = "posting_acceptance_in_progress"

	// arbitrage
	PostingInArbitration ShipmentStatus = "posting_in_arbitration"

	// created
	PostingCreated ShipmentStatus = "posting_created"

	// in the freight
	PostingInCarriage ShipmentStatus = "posting_in_carriage"

	// not added to the freight
	PostingNotInCarriage ShipmentStatus = "posting_not_in_carriage"

	// registered
	PostingRegistered ShipmentStatus = "posting_registered"

	// is handed over to the delivery service
	PostingTransferringToDelivery ShipmentStatus = "posting_transferring_to_delivery"

	// waiting for passport data
	PostingAwaitingPassportData ShipmentStatus = "posting_awaiting_passport_data"

	// created
	PostingCreatedSubstatus ShipmentStatus = "posting_created"

	// awaiting registration
	PostingAwaitingRegistration ShipmentStatus = "posting_awaiting_registration"

	// registration error
	PostingRegistrationError ShipmentStatus = "posting_registration_error"

	// created
	PostingSplitPending ShipmentStatus = "posting_split_pending"

	// canceled
	PostingCancelled ShipmentStatus = "posting_canceled"

	// customer delivery arbitrage
	PostingInClientArbitration ShipmentStatus = "posting_in_client_arbitration"

	// delivered
	PostingDelivered ShipmentStatus = "posting_delivered"

	// recieved
	PostingReceived ShipmentStatus = "posting_received"

	// presumably delivered
	PostingConditionallyDelivered ShipmentStatus = "posting_conditionally_delivered"

	// courier on the way
	PostingInCourierService ShipmentStatus = "posting_in_courier_service"

	// at the pick-up point
	PostingInPickupPoint ShipmentStatus = "posting_in_pickup_point"

	// on the way to the city
	PostingOnWayToCity ShipmentStatus = "posting_on_way_to_city"

	// on the way to the pick-up point
	PostingOnWayToPickupPoint ShipmentStatus = "posting_on_way_to_pickup_point"

	// returned to the warehouse
	PostingReturnedToWarehouse ShipmentStatus = "posting_returned_to_warehouse"

	// is handed over to the courier
	PostingTransferredToCourierService ShipmentStatus = "posting_transferred_to_courier_service"

	// handed over to the driver
	PostingDriverPickup ShipmentStatus = "posting_driver_pick_up"

	// not accepted at the sorting center
	PostingNotInSortCenter ShipmentStatus = "posting_not_in_sort_center"

	// sent by the seller
	SentBySellerSubstatus ShipmentStatus = "sent_by_seller"
)

type TPLIntegrationType string

const (
	// delivery by the Ozon logistics
	OzonTPLType TPLIntegrationType = "ozon"

	// delivery by a third-party service, Ozon registers the order
	AggregatorTPLType TPLIntegrationType = "aggregator"

	// delivery by a third-party service, the seller registers the order
	TrackingTPLType TPLIntegrationType = "3pl_tracking"

	// delivery by the seller
	NonIntegratedTPLType TPLIntegrationType = "non_integrated"

	// Russian Post delivery scheme
	HybrydTPLType TPLIntegrationType = "hybryd"
)

type DetailsDeliveryItemName string

const (
	DirectFlowLogisticSumDetailsDeliveryItemName DetailsDeliveryItemName = "MarketplaceServiceItemDirectFlowLogisticSum"
	DropoffDetailsDeliveryItemName               DetailsDeliveryItemName = "MarketplaceServiceItemDropoff"
	DelivToCustomerDetailsDeliveryItemName       DetailsDeliveryItemName = "MarketplaceServiceItemDelivToCustomer"
)

type DetailsReturnServiceName string

const (
	ReturnAfterDelivToCustomerDetailsReturnServiceName DetailsReturnServiceName = "MarketplaceServiceItemReturnAfterDelivToCustomer"
	ReturnPartGoodsCustomerDetailsReturnServiceName    DetailsReturnServiceName = "MarketplaceServiceItemReturnPartGoodsCustomer"
	ReturnNotDelivToCustomerDetailsReturnServiceName   DetailsReturnServiceName = "MarketplaceServiceItemReturnNotDelivToCustomer"
	ReturnFlowLogisticDetailsReturnServiceName         DetailsReturnServiceName = "MarketplaceServiceItemReturnFlowLogistic"
)

type DetailsServiceItemName string

const (
	OtherMarketAndTech                          DetailsServiceItemName = "MarketplaceServiceItemOtherMarketAndTechService"
	ReturnStorageServiceAtThePickupPointFbsItem DetailsServiceItemName = "MarketplaceReturnStorageServiceAtThePickupPointFbsItem"
	SaleReviewsItem                             DetailsServiceItemName = "MarketplaceSaleReviewsItem"
	ServicePremiumCashbackIndividualPoints      DetailsServiceItemName = "MarketplaceServicePremiumCashbackIndividualPoints"
	ServiceStorageItem                          DetailsServiceItemName = "MarketplaceServiceStorageItem"
	ServiceStockDisposal                        DetailsServiceItemName = "MarketplaceServiceStockDisposal"
	ReturnDisposalServiceFbsItem                DetailsServiceItemName = "MarketplaceReturnDisposalServiceFbsItem"
	ServiceItemFlexiblePaymentSchedule          DetailsServiceItemName = "MarketplaceServiceItemFlexiblePaymentSchedule"
	ServiceProcessingSpoilage                   DetailsServiceItemName = "MarketplaceServiceProcessingSpoilage"
	ServiceProcessingIdentifiedSurplus          DetailsServiceItemName = "MarketplaceServiceProcessingIdentifiedSurplus"
	ServiceProcessingIdentifiedDiscrepancies    DetailsServiceItemName = "MarketplaceServiceProcessingIdentifiedDiscrepancies"
	ServiceItemInternetSiteAdvertising          DetailsServiceItemName = "MarketplaceServiceItemInternetSiteAdvertising"
	ServiceItemPremiumSubscribtion              DetailsServiceItemName = "MarketplaceServiceItemSubscribtionPremium"
	AgencyFeeAggregator3PLGlobalItem            DetailsServiceItemName = "MarketplaceAgencyFeeAggregator3PLGlobalItem"
)

type DetailsOtherItemName string

const (
	RedistributionOfAcquiringOperation                   DetailsOtherItemName = "MarketplaceRedistributionOfAcquiringOperation"
	CompensationLossOfGoodsOperation                     DetailsOtherItemName = "MarketplaceSellerCompensationLossOfGoodsOperation"
	CorrectionOperation                                  DetailsOtherItemName = "MarketplaceSellerCorrectionOperation"
	OperationCorrectionSeller                            DetailsOtherItemName = "OperationCorrectionSeller"
	OperationMarketplaceWithHoldingForUndeliverableGoods DetailsOtherItemName = "OperationMarketplaceWithHoldingForUndeliverableGoods"
	OperationClaim                                       DetailsOtherItemName = "OperationClaim"
)

type StrategyType string

const (
	MinExtPrice StrategyType = "MIN_EXT_PRICE"
	CompPrice   StrategyType = "COMP_PRICE"
)

type StrategyUpdateType string

const (
	StrategyEnabled          StrategyUpdateType = "strategyEnabled"
	StrategyDisabled         StrategyUpdateType = "strategyDisabled"
	StrategyChanged          StrategyUpdateType = "strategyChanged"
	StrategyCreated          StrategyUpdateType = "strategyCreated"
	StrategyItemsListChanged StrategyUpdateType = "strategyItemsListChanged"
)

type ShipmentCertificateFilterStatus string

const (
	// new
	ShitmentCertificateFilterNew ShipmentCertificateFilterStatus = "new"

	// retry creation
	ShitmentCertificateFilterAwaitingRetry ShipmentCertificateFilterStatus = "awaiting-retry"

	// is being packaged
	ShitmentCertificateFilterInProcess ShipmentCertificateFilterStatus = "in_process"

	// created
	ShitmentCertificateFilterSuccess ShipmentCertificateFilterStatus = "success"

	// creation error
	ShitmentCertificateFilterError ShipmentCertificateFilterStatus = "error"

	// sent
	ShitmentCertificateFilterSend ShipmentCertificateFilterStatus = "sent"

	// received
	ShitmentCertificateFilterReceived ShipmentCertificateFilterStatus = "received"

	// packaged
	ShitmentCertificateFilterFormed ShipmentCertificateFilterStatus = "formed"

	// canceled
	ShitmentCertificateFilterCancelled ShipmentCertificateFilterStatus = "cancelled"

	// in the queue for packaging
	ShitmentCertificateFilterPending ShipmentCertificateFilterStatus = "pending"

	// in the queue for completion
	ShitmentCertificateFilterCompletionEnqueued ShipmentCertificateFilterStatus = "completion_enqueued"

	// in the process of completion
	ShitmentCertificateFilterCompletionProcessing ShipmentCertificateFilterStatus = "completion_processing"

	// completion error
	ShitmentCertificateFilterCompletionFailed ShipmentCertificateFilterStatus = "completion_failed"

	// in the queue for cancellation
	ShitmentCertificateFilterCancelationEnqueued ShipmentCertificateFilterStatus = "cancelation_enqueued"

	// in the process of cancellation
	ShitmentCertificateFilterCancelationProcessing ShipmentCertificateFilterStatus = "cancelation_processing"

	// cancellation error
	ShitmentCertificateFilterCancelationFailed ShipmentCertificateFilterStatus = "cancelation_failed"

	// completed
	ShitmentCertificateFilterCompleted ShipmentCertificateFilterStatus = "completed"

	// closed
	ShitmentCertificateFilterClosed ShipmentCertificateFilterStatus = "closed"
)

type PRROptionStatus string

const (
	// carrying the bulky product using the elevator
	PRROptionLift PRROptionStatus = "lift"

	// carrying the bulky product upstairs
	PRROptionStairs PRROptionStatus = "stairs"

	// the customer canceled the service,
	// you don't need to lift the shipment
	PRROptionNone PRROptionStatus = "none"

	// delivery is included in the price.
	// According to the offer you need to
	// deliver products to the floor
	PRROptionDeliveryDefault PRROptionStatus = "delivery_default"
)

type GetFBSReturnsFilterStatus string

const (
	Moving                    GetFBSReturnsFilterStatus = "moving"
	ReturnedToSeller          GetFBSReturnsFilterStatus = "returned_to_seller"
	WaitingForSeller          GetFBSReturnsFilterStatus = "waiting_for_seller"
	AcceptedFromCustomer      GetFBSReturnsFilterStatus = "accepted_from_customer"
	CancelledWithCompensation GetFBSReturnsFilterStatus = "cancelled_with_compensation"
	ReadyForShipment          GetFBSReturnsFilterStatus = "ready_for_shipment"
	Disposing                 GetFBSReturnsFilterStatus = "disposing"
	Disposed                  GetFBSReturnsFilterStatus = "disposed"
	ArrivedForResale          GetFBSReturnsFilterStatus = "arrived_for_resale"
	MovingToResale            GetFBSReturnsFilterStatus = "moving_to_resale"
)

type GetFBOReturnsFilterStatus string

const (
	GetFBOReturnsFilterStatusCreated                   GetFBOReturnsFilterStatus = "Created"
	GetFBOReturnsFilterStatusReturnedToOzon            GetFBOReturnsFilterStatus = "ReturnedToOzon"
	GetFBOReturnsFilterStatusCancelled                 GetFBOReturnsFilterStatus = "Cancelled"
	GetFBOReturnsFilterStatusCancelledWithCompensation GetFBOReturnsFilterStatus = "CancelledWithCompensation"
)

type GetFBOReturnsReturnStatus string

const (
	GetFBOReturnsReturnStatusCancelled            GetFBOReturnsReturnStatus = "Возврат отменен"
	GetFBOReturnsReturnStatusAcceptedFromCustomer GetFBOReturnsReturnStatus = "Принят от покупателя"
	GetFBOReturnsReturnStatusReceivedAtOzon       GetFBOReturnsReturnStatus = "Получен в Ozon"
)

type DigitalActType string

const (
	// acceptance certificate
	DigitalActTypeOfAcceptance DigitalActType = "act_of_acceptance"

	// discrepancy certificate
	DigitalActTypeOfMismatch DigitalActType = "act_of_mismatch"

	// surplus certificate
	DigitalActTypeOfExcess DigitalActType = "act_of_excess"
)

type PriceStrategy string

const (
	// enable
	PriceStrategyEnabled PriceStrategy = "ENABLED"

	// disable
	PriceStrategyDisabled PriceStrategy = "DISABLED"

	// don't change anything. Default value
	PriceStrategyUnknown PriceStrategy = "UNKNOWN"
)

type FBPFilter string

const (
	// all shipments matching other filters will be returned in the response
	FBPFilterAll FBPFilter = "all"

	// only FBP shipments will be returned
	FBPFilterOnly FBPFilter = "only"

	// all shipments except FBP will be returned
	FBPFilterWithout FBPFilter = "without"
)

type InvoiceCurrency string

const (
	// dollar
	InvoiceCurrencyUSD InvoiceCurrency = "USD"

	// euro
	InvoiceCurrencyEUR InvoiceCurrency = "EUR"

	// Turkish lira
	InvoiceCurrencyTRY InvoiceCurrency = "TRY"

	// yuan
	InvoiceCurrencyCNY InvoiceCurrency = "CNY"

	// ruble
	InvoiceCurrencyRUB InvoiceCurrency = "RUB"

	// pound sterling
	InvoiceCurrencyGBP InvoiceCurrency = "GBP"
)

type ReportType string

const (
	// products report
	ReportTypeSellerProducts ReportType = "SELLER_PRODUCTS"

	// transactions report
	ReportTypeSellerTransactions ReportType = "SELLER_TRANSACTIONS"

	// product prices report
	ReportTypeSellerProductPrices ReportType = "SELLER_PRODUCT_PRICES"

	// stocks report
	ReportTypeSellerStock ReportType = "SELLER_STOCK"

	// products movement report
	ReportTypeSellerProductMovement ReportType = "SELLER_PRODUCT_MOVEMENT"

	// returns report
	ReportTypeSellerReturns ReportType = "SELLER_RETURNS"

	// shipments report
	ReportTypeSellerPostings ReportType = "SELLER_POSTINGS"

	// financial report
	ReportTypeSellerFinance ReportType = "SELLER_FINANCE"

	// report on sales to legal entities
	ReportTypeDocB2BSales ReportType = "DOCUMENT_B2B_SALES"

	// settlement report
	ReportTypeMutualSettlement ReportType = "MUTUAL_SETTLEMENT"
)

type ReportInfoStatus string

const (
	ReportInfoWaiting    ReportInfoStatus = "waiting"
	ReportInfoProcessing ReportInfoStatus = "processing"
	ReportInfoSuccess    ReportInfoStatus = "success"
	ReportInfoFailed     ReportInfoStatus = "failed"
)

type SKUAvailability string

const (
	SKUAvailabilityHidden    = "HIDDEN"
	SKUAvailabilityAvailable = "AVAILABLE"

	// SKU is deleted
	SKUAvailabilityUnavailable = "UNAVAILABLE"
)

type RFBSReturnsGroupState string

const (
	// All requests
	RFBSReturnsGroupStateAll RFBSReturnsGroupState = "All"

	// New
	RFBSReturnsGroupStateNew RFBSReturnsGroupState = "New"

	// Returned product is on the way for check
	RFBSReturnsGroupStateDelivering RFBSReturnsGroupState = "Delivering"

	// Returned product is being checked
	RFBSReturnsGroupStateCheckout RFBSReturnsGroupState = "Checkout"

	// Disputed
	RFBSReturnsGroupStateArbitration RFBSReturnsGroupState = "Arbitration"

	// Approved
	RFBSReturnsGroupStateApproved RFBSReturnsGroupState = "Approved"

	// Rejected
	RFBSReturnsGroupStateRejected RFBSReturnsGroupState = "Rejected"
)

type GetRFBSReturnsCurrency string

const (
	// Russian ruble
	GetRFBSReturnsCurrencyRUB GetRFBSReturnsCurrency = "RUB"

	// Belarusian ruble
	GetRFBSReturnsCurrencyBYN GetRFBSReturnsCurrency = "BYN"

	// Tenge
	GetRFBSReturnsCurrencyKZT GetRFBSReturnsCurrency = "KZT"

	// Euro
	GetRFBSReturnsCurrencyEUR GetRFBSReturnsCurrency = "EUR"

	// US dollar
	GetRFBSReturnsCurrencyUSD GetRFBSReturnsCurrency = "USD"

	// Yuan
	GetRFBSReturnsCurrencyCNY GetRFBSReturnsCurrency = "CNY"
)

type GiveoutStatus string

const (
	// Undefined, contact support team
	GiveoutStatusUnspecified GiveoutStatus = "GIVEOUT_STATUS_UNSPECIFIED"

	// Created
	GiveoutStatusCreated GiveoutStatus = "GIVEOUT_STATUS_CREATED"

	// Approved
	GiveoutStatusApproved GiveoutStatus = "GIVEOUT_STATUS_APPROVED"

	// Completed
	GiveoutStatusCompleted GiveoutStatus = "GIVEOUT_STATUS_COMPLETED"

	// Cancelled
	GiveoutStatusCancelled GiveoutStatus = "GIVEOUT_STATUS_CANCELLED"
)

type GiveoutDeliverySchema string

const (
	// Undefined, contact support team
	GiveoutDeliverySchemaUnspecified GiveoutDeliverySchema = "GIVEOUT_DELIVERY_SCHEMA_UNSPECIFIED"

	// FBO
	GiveoutDeliverySchemaFBO GiveoutDeliverySchema = "GIVEOUT_DELIVERY_SCHEMA_FBO"

	// FBS
	GiveoutDeliverySchemaFBS GiveoutDeliverySchema = "GIVEOUT_DELIVERY_SCHEMA_FBS"
)

type MandatoryMarkStatus string

const (
	// Labeling is processed
	MandatoryMarkStatusProcessing MandatoryMarkStatus = "processing"

	// Check is passed
	MandatoryMarkStatusPassed MandatoryMarkStatus = "passed"

	// Check is failed
	MandatoryMarkStatusFailed MandatoryMarkStatus = "failed"
)

type GetCarriageStatus string

const (
	// acceptance in progress
	GetCarriageStatusReceived GetCarriageStatus = "received"

	// closed after acceptance
	GetCarriageStatusClosed GetCarriageStatus = "closed"

	GetCarriageStatusSended GetCarriageStatus = "sended"

	GetCarriageStatusCancelled GetCarriageStatus = "cancelled"
)

type TransactionOperationService string

const (
	// return of unclaimed products from the customer to the warehouse
	TransactionNotDelivered TransactionOperationService = "MarketplaceNotDeliveredCostItem"

	// return from the customer to the warehouse after delivery
	TransactionReturnAfterDelivery TransactionOperationService = "TransactionOperationServiceNotDelivered"

	// product delivery to the customer
	TransactionDelivery TransactionOperationService = "MarketplaceDeliveryCostItem"

	// purchasing reviews on the platform
	TransactionSaleReviews TransactionOperationService = "MarketplaceSaleReviewsItem"

	// products delivery to the Ozon warehouse (cross docking)
	TransactionItemAdForSupplierLogistic TransactionOperationService = "ItemAdvertisementForSupplierLogistic"

	// product placement service
	TransactionServiceStorageItem TransactionOperationService = "OperationMarketplaceServiceStorage"

	// products promotion
	TransactionMarketingActionCost TransactionOperationService = "MarketplaceMarketingActionCostItem"

	// promotion and selling on an instalment plan
	TransactionServiceItemInstallment TransactionOperationService = "MarketplaceServiceItemInstallment"

	// mandatory products labeling
	TransactionServiceMarkingItems TransactionOperationService = "MarketplaceServiceItemMarkingItems"

	// flexible payment schedule
	TransactionServiceFlexiblePaymentSchedule TransactionOperationService = "MarketplaceServiceItemFlexiblePaymentSchedule"

	// picking up products for removal by the seller
	TransactionServiceReturnFromStock TransactionOperationService = "MarketplaceServiceItemReturnFromStock"

	TransactionServiceStarsMembership TransactionOperationService = "ItemAgentServiceStarsMembership"

	// forwarding trade
	TransactionItemAdForSupplierLogisticSeller TransactionOperationService = "ItemAdvertisementForSupplierLogisticSeller"

	// last mile
	TransactionServiceDeliveryToCustomer TransactionOperationService = "MarketplaceServiceItemDelivToCustomer"

	// pipeline
	TransactionServiceDirectFlowTrans TransactionOperationService = "MarketplaceServiceItemDirectFlowTrans"

	// shipment processing
	TransactionServiceDropoffFF TransactionOperationService = "MarketplaceServiceItemDropoffFF"

	// shipment processing
	TransactionServiceDropoffPVZ TransactionOperationService = "MarketplaceServiceItemDropoffPVZ"

	// shipment processing
	TransactionServiceDropoffSC TransactionOperationService = "MarketplaceServiceItemDropoffSC"

	// order packaging
	TransactionServiceFulfillment TransactionOperationService = "MarketplaceServiceItemFulfillment"

	// picking products up by car from the seller's address (Pick-up)
	TransactionServicePickup TransactionOperationService = "MarketplaceServiceItemPickup"

	// return processing
	TransactionServiceReturnAfterDelivToCustomer TransactionOperationService = "MarketplaceServiceItemReturnAfterDelivToCustomer"

	// reverse pipeline
	TransactionServiceReturnFlowTrans TransactionOperationService = "MarketplaceServiceItemReturnFlowTrans"

	// cancellation processing
	TransactionServiceReturnNotDelivToCustomer TransactionOperationService = "MarketplaceServiceItemReturnNotDelivToCustomer"

	// unredeemed order processing
	TransactionServiceReturnPartGoodsCustomer TransactionOperationService = "MarketplaceServiceItemReturnPartGoodsCustomer"

	// acquiring payment
	TransactionRedistributionOfAcquiringOperation TransactionOperationService = "MarketplaceRedistributionOfAcquiringOperation"

	// FBS return short-term placement
	TransactionServiceAtPickupPointFBS TransactionOperationService = "MarketplaceReturnStorageServiceAtThePickupPointFbsItem"

	// FBS return long-term placement
	TransactionServiceInWarehouseFBS TransactionOperationService = "MarketplaceReturnStorageServiceInTheWarehouseFbsItem"

	// bulky products delivery
	TransactionServiceDeliveryKGT TransactionOperationService = "MarketplaceServiceItemDeliveryKGT"

	// logistics
	TransactionServiceDirectFlowLogistic TransactionOperationService = "MarketplaceServiceItemDirectFlowLogistic"

	// reverse logistics
	TransactionServiceReturnFlowLogistic TransactionOperationService = "MarketplaceServiceItemReturnFlowLogistic"

	// "Seller's Bonus" promotion service
	TransactionServicePremiumCashbackIndPoints TransactionOperationService = "MarketplaceServicePremiumCashbackIndividualPoints"

	// Premium promotion service, fixed commission
	TransactionServicePremiumPromotion TransactionOperationService = "MarketplaceServicePremiumPromotion"

	// withholding for product shortage
	TransactionServiceWithHoldingForUndeliverableGoods TransactionOperationService = "OperationMarketplaceWithHoldingForUndeliverableGoods"

	// drop-off service at the pick-up point
	TransactionServiceDropoffPPZ TransactionOperationService = "MarketplaceServiceItemDropoffPPZ"

	// reissue of returns at the pick-up point
	TransactionServiceRedistributionReturnsPVZ TransactionOperationService = "MarketplaceServiceItemRedistributionReturnsPVZ"

	// Agregator 3PL Globalagency service tariffication
	TransactionServiceAgencyFeeAggregator3PLGlobal TransactionOperationService = "OperationMarketplaceAgencyFeeAggregator3PLGlobal"
)

type PaymentTypeGroupName string

const (
	PaymentTypeGroupByCardOnline              PaymentTypeGroupName = "by card online"
	PaymentTypeGroupOzonCard                  PaymentTypeGroupName = "Ozon Card"
	PaymentTypeGroupOzonCardAtCheckout        PaymentTypeGroupName = "Ozon Card at checkout"
	PaymentTypeGroupBySavedBankCardUponPickup PaymentTypeGroupName = "by saved bank card upon pick-up"
	PaymentTypeGroupFasterPaymentSystem       PaymentTypeGroupName = "Faster payment system"
	PaymentTypeGroupOzonInstallment           PaymentTypeGroupName = "Ozon Installment"
	PaymentTypeGroupPaymentToCurrentAccount   PaymentTypeGroupName = "payment to current account"
	PaymentTypeGroupSberpay                   PaymentTypeGroupName = "Sberpay"
)

type VisualStatus string

const (
	// dispute with the customer has been opened
	VisualStatusDisputeOpened VisualStatus = "DisputeOpened"

	// pending with the seller
	VisualStatusOnSellerApproval VisualStatus = "OnSellerApproval"

	// at the pick-up point
	VisualStatusArrivedAtReturnPlace VisualStatus = "ArrivedAtReturnPlace"

	// pending clarification by the seller
	VisualStatusOnSellerClarification VisualStatus = "OnSellerClarification"

	// pending clarification by the seller after partial compensation
	VisualStatusOnSellerClarificationPartial VisualStatus = "OnSellerClarificationAfterPartialCompensation"

	// partial compensation offered
	VisualStatusOfferedPartial VisualStatus = "OfferedPartialCompensation"

	// refund approved
	VisualStatusReturnMoneyApproved VisualStatus = "ReturnMoneyApproved"

	// partial compensation provided
	VisualStatusPartialReturned VisualStatus = "PartialCompensationReturned"

	// refund rejected, dispute isn't opened
	VisualStatusCancelledDisputeNotOpen VisualStatus = "CancelledDisputeNotOpen"

	// request rejected
	VisualStatusRejected VisualStatus = "Rejected"

	// request rejected by Ozon
	VisualStatusCrmRejected VisualStatus = "CrmRejected"

	// request canceled
	VisualStatusCancelled VisualStatus = "Cancelled"

	// request approved by the seller
	VisualStatusApproved VisualStatus = "Approved"

	// request approved by Ozon
	VisualStatusApprovedByOzon VisualStatus = "ApprovedByOzon"

	// seller received the return
	VisualStatusReceivedBySeller VisualStatus = "ReceivedBySeller"

	// return is on its way to the seller
	VisualStatusMovingToSeller VisualStatus = "MovingToSeller"

	// seller received the refund
	VisualStatusReturnCompensated VisualStatus = "ReturnCompensated"

	// courier is taking the return to the seller
	VisualStatusReturningByCourier VisualStatus = "ReturningByCourier"

	// on disposal
	VisualStatusUtilizing VisualStatus = "Utilizing"

	// disposed of
	VisualStatusUtilized VisualStatus = "Utilized"

	// customer received full refund
	VisualStatusMoneyReturned VisualStatus = "MoneyReturned"

	// partial refund has been approved
	VisualStatusPartialInProcess VisualStatus = "PartialCompensationInProcess"

	// seller opened a dispute
	VisualStatusDisputeYouOpened VisualStatus = "DisputeYouOpened"

	// compensation rejected
	VisualStatusCompensationRejected VisualStatus = "CompensationRejected"

	// support request sent
	VisualStatusDisputeOpening VisualStatus = "DisputeOpening"

	// awaiting your decision on compensation
	VisualStatusCompensationOffered VisualStatus = "CompensationOffered"

	// awaiting compensation
	VisualStatusWaitingCompensation VisualStatus = "WaitingCompensation"

	// an error occurred when sending the support request
	VisualStatusSendingError VisualStatus = "SendingError"

	// decision period has expired
	VisualStatusCompensationRejectedBySla VisualStatus = "CompensationRejectedBySla"

	// seller has refused compensation
	VisualStatusCompensationRejectedBySeller VisualStatus = "CompensationRejectedBySeller"

	// on the way to the Ozon warehouse
	VisualStatusMovingToOzon VisualStatus = "MovingToOzon"

	// arrived at the Ozon warehouse
	VisualStatusReturnedToOzon VisualStatus = "ReturnedToOzon"

	// quick refund
	VisualStatusMoneyReturnedBySystem VisualStatus = "MoneyReturnedBySystem"

	// awaiting shipping
	VisualStatusWaitingShipment VisualStatus = "WaitingShipment"
)

type VAT string

const (
	VAT0   VAT = "0"
	VAT005 VAT = "0.05"
	VAT007 VAT = "0.07"
	VAT01  VAT = "0.1"
	VAT02  VAT = "0.2"
)
