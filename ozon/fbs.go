package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type FBS struct {
	client *core.Client
}

type ListUnprocessedShipmentsParams struct {
	Direction string                         `json:"dir"`
	Filter    ListUnprocessedShipmentsFilter `json:"filter"`
	Limit     int64                          `json:"limit"`
	Offset    int64                          `json:"offset"`
	With      ListUnprocessedShipmentsWith   `json:"with"`
}

type ListUnprocessedShipmentsFilter struct {
	CutoffFrom         time.Time `json:"cutoff_from"`
	CutoffTo           time.Time `json:"cutoff_to"`
	DeliveringDateFrom time.Time `json:"delivering_date_from"`
	DeliveringDateTo   time.Time `json:"delivering_date_to"`
	DeliveryMethodId   []int64   `json:"deliveryMethodId"`
	ProviderId         []int64   `json:"provider_id"`
	Status             string    `json:"status"`
	WarehouseId        []int64   `json:"warehouse_id"`
}

type ListUnprocessedShipmentsWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes      bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	Translit      bool `json:"translit"`
}

type ListUnprocessedShipmentsResponse struct {
	core.CommonResponse

	Result ListUnprocessedShipmentsResult `json:"result"`
}

type ListUnprocessedShipmentsResult struct {
	Count    int64        `json:"count"`
	Postings []FBSPosting `json:"postings"`
}

type FBSPosting struct {
	Addressee struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"addressee"`

	AnalyticsData struct {
		City                 string    `json:"city"`
		DeliveryDateBegin    time.Time `json:"delivery_date_begin"`
		DeliveryDateEnd      time.Time `json:"delivery_date_end"`
		DeliveryType         string    `json:"delivery_type"`
		IsLegal              bool      `json:"is_legal"`
		IsPremium            bool      `json:"is_premium"`
		PaymentTypeGroupName string    `json:"payment_type_group_name"`
		Region               string    `json:"region"`
		TPLProvider          string    `json:"tpl_provider"`
		TPLProviderId        int64     `json:"tpl_provider_id"`
		Warehouse            string    `json:"warehouse"`
		WarehouseId          int64     `json:"warehouse_id"`
	} `json:"analytics_data"`

	Barcodes FBSBarcode `json:"barcodes"`

	Cancellation FBSCancellation `json:"cancellation"`

	Customer FBSCustomer `json:"customer"`

	DeliveringDate time.Time `json:"delivering_date"`

	DeliveryMethod FBSDeliveryMethod `json:"delivery_method"`

	FinancialData FBSFinancialData `json:"financial_data"`

	InProccessAt        time.Time `json:"in_process_at"`
	IsExpress           bool      `json:"is_express"`
	IsMultibox          bool      `json:"is_multibox"`
	MultiBoxQuantity    int32     `json:"multi_box_qty"`
	OrderId             int64     `json:"order_id"`
	OrderNumber         string    `json:"order_number"`
	ParentPostingNumber string    `json:"parent_posting_number"`
	PostingNumber       string    `json:"posting_number"`

	Products []PostingProduct `json:"products"`

	Requirements FBSRequirements `json:"requirements"`

	ShipmentDate       time.Time `json:"shipment_date"`
	Status             string    `json:"status"`
	TPLIntegrationType string    `json:"tpl_integration_type"`
	TrackingNumber     string    `json:"tracking_number"`
}

type FBSBarcode struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type FBSCancellation struct {
	AffectCancellationRating bool   `json:"affect_cancellation_rating"`
	CancelReason             string `json:"cancel_reason"`
	CancelReasonId           int64  `json:"cancel_reason_id"`
	CancellationInitiator    string `json:"cancellation_initiator"`
	CancellationType         string `json:"cancellation_type"`
	CancelledAfterShip       bool   `json:"cancelled_after_ship"`
}

type FBSDeliveryMethod struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	TPLProvider   string `json:"tpl_provider"`
	TPLProviderId int64  `json:"tpl_provider_id"`
	Warehouse     string `json:"warehouse"`
	WarehouseId   int64  `json:"warehouse_id"`
}

type FBSFinancialData struct {
	ClusterFrom     string                 `json:"cluster_from"`
	ClusterTo       string                 `json:"cluster_to"`
	PostingServices MarketplaceServices    `json:"posting_services"`
	Products        []FinancialDataProduct `json:"products"`
}

type FBSRequirements struct {
	ProductsRequiringGTD           []string `json:"products_requiring_gtd"`
	ProductsRequiringCountry       []string `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRNPT          []string `json:"products_requiring_rnpt"`
}

type PostingProduct struct {
	MandatoryMark []string `json:"mandatory_mark"`
	Name          string   `json:"name"`
	OfferId       string   `json:"offer_id"`
	CurrencyCode  string   `json:"currency_code"`
	Price         string   `json:"price"`
	Quantity      int32    `json:"quantity"`
	SKU           int64    `json:"sku"`
}

type FBSCustomer struct {
	Address struct {
		AddressTail     string  `json:"address_tail"`
		City            string  `json:"city"`
		Comment         string  `json:"comment"`
		Country         string  `json:"country"`
		District        string  `json:"district"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		ProviderPVZCode string  `json:"provider_pvz_code"`
		PVZCode         int64   `json:"pvz_code"`
		Region          string  `json:"region"`
		ZIPCode         string  `json:"zip_code"`
	} `json:"customer"`

	CustomerEmail string `json:"customer_email"`
	CustomerId    int64  `json:"customer_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
}

type MarketplaceServices struct {
	DeliveryToCustomer            float64 `json:"marketplace_service_item_deliv_to_customer"`
	DirectFlowTrans               float64 `json:"marketplace_service_item_direct_flow_trans"`
	DropoffFF                     float64 `json:"marketplace_service_item_item_dropoff_ff"`
	DropoffPVZ                    float64 `json:"marketplace_service_item_dropoff_pvz"`
	DropoffSC                     float64 `json:"marketplace_service_item_dropoff_sc"`
	Fulfillment                   float64 `json:"marketplace_service_item_fulfillment"`
	Pickup                        float64 `json:"marketplace_service_item_pickup"`
	ReturnAfterDeliveryToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`
	ReturnFlowTrans               float64 `json:"marketplace_service_item_return_flow_trans"`
	ReturnNotDeliveryToCustomer   float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`
	ReturnPartGoodsCustomer       float64 `json:"marketplace_service_item_return_part_goods_customer"`
}

type FinancialDataProduct struct {
	Actions                 []string            `json:"actions"`
	ClientPrice             string              `json:"client_price"`
	CommissionAmount        float64             `json:"commission_amount"`
	CommissionPercent       int64               `json:"commission_percent"`
	CommissionsCurrencyCode string              `json:"commissions_currency_code"`
	ItemServices            MarketplaceServices `json:"item_services"`
	CurrencyCode            string              `json:"currency_code"`
	OldPrice                float64             `json:"old_price"`
	Payout                  float64             `json:"payout"`
	Picking                 struct {
		Amount float64   `json:"amount"`
		Moment time.Time `json:"moment"`
		Tag    string    `json:"tag"`
	} `json:"picking"`
	Price                float64 `json:"price"`
	ProductId            int64   `json:"product_id"`
	Quantity             int64   `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue   float64 `json:"total_discount_value"`
}

func (c FBS) ListUnprocessedShipments(params *ListUnprocessedShipmentsParams) (*ListUnprocessedShipmentsResponse, error) {
	url := "/v3/posting/fbs/unfulfilled/list"

	resp := &ListUnprocessedShipmentsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSShipmentsListParams struct {
	// Sorting direction
	Direction string `json:"direction"`

	//Filter
	Filter GetFBSShipmentsListFilter `json:"filter"`

	// Number of shipments in the response:
	//   - maximum is 50,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Additional fields that should be added to the response
	With GetFBSShipmentsListWith `json:"with"`
}

type GetFBSShipmentsListFilter struct {
	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_id"`

	// Order identifier
	OrderId int64 `json:"order_id"`

	// Delivery service identifier
	ProviderId []int64 `json:"provider_id"`

	// Start date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z
	Since time.Time `json:"since"`

	// End date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z.
	To time.Time `json:"to"`

	// Shipment status
	Status string `json:"status"`

	// Warehouse identifier
	WarehouseId []int64 `json:"warehouse_id"`
}

type GetFBSShipmentsListWith struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type GetFBSShipmentsListResponse struct {
	core.CommonResponse

	// Array of shipments
	Result struct {
		// Indicates that the response returned not the entire array of shipments:
		//
		//   - true — it is necessary to make a new request with a different offset value to get information on the remaining shipments;
		//   - false — the entire array of shipments for the filter specified in the request was returned in the response
		HasNext bool `json:"has_next"`

		// Shipment details
		Postings []FBSPosting `json:"postings"`
	} `json:"result"`
}

// Returns a list of shipments for the specified time period: it shouldn't be longer than one year.
//
// You can filter shipments by their status. The list of available statuses is specified in the description of the filter.status parameter.
//
// The true value of the has_next parameter in the response means there is not the entire array of shipments in the response. To get information on the remaining shipments, make a new request with a different offset value.
func (c FBS) GetFBSShipmentsList(params *GetFBSShipmentsListParams) (*GetFBSShipmentsListResponse, error) {
	url := "/v3/posting/fbs/list"

	resp := &GetFBSShipmentsListResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PackOrderParams struct {
	// List of packages. Each package contains a list of shipments that the order was divided into
	Packages []PackOrderPackage `json:"packages"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Additional information
	With PackOrderWith `json:"with"`
}

type PackOrderPackage struct {
	Products []PackOrderPackageProduct `json:"products"`
}

type PackOrderPackageProduct struct {
	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product items quantity
	Quantity int32 `json:"quantity"`
}

type PackOrderWith struct {
	// Pass true to get additional information
	AdditionalData bool `json:"additional_data"`
}

type PackOrderResponse struct {
	core.CommonResponse

	// Additional information about shipments
	AdditionalData []struct {
		// Shipment number
		PostingNumber string `json:"posting_number"`

		// List of products in the shipment
		Products []PostingProduct `json:"products"`
	} `json:"additional_data"`

	// Order packaging result
	Result []string `json:"result"`
}

// Divides the order into shipments and changes its status to awaiting_deliver.
//
// Each element of the packages may contain several instances of the products. One instance of the products is one shipment. Each element of the products is a product included into the shipment.
//
// It is necessary to split the order if:
//
// the products do not fit in one package,
// the products cannot be put in one package.
// Differs from /v2/posting/fbs/ship by the presence of the field exemplar_info in the request.
//
// If necessary, specify the number of the cargo customs declaration in the gtd parameter. If it is missing, pass the value is_gtd_absent = true
func (c FBS) PackOrder(params *PackOrderParams) (*PackOrderResponse, error) {
	url := "/v4/posting/fbs/ship"

	resp := &PackOrderResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ValidateLabelingCodesParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Products list
	Products []ValidateLabelingCodesProduct `json:"products"`
}

type ValidateLabelingCodesProduct struct {
	// Product items data
	Exemplars []ValidateLabelingCodesExemplar `json:"exemplars"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type ValidateLabelingCodesExemplar struct {
	// Сustoms cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// Mandatory “Chestny ZNAK” labeling
	MandatoryMark string `json:"mandatory_mark"`

	// Product batch registration number
	RNPT string `json:"rnpt"`
}

type ValidateLabelingCodesResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Products list
		Products []struct {
			// Error code
			Error string `json:"error"`

			// Product items data
			Exemplars []struct {
				// Product item validation errors
				Errors []string `json:"errors"`

				// Сustoms cargo declaration (CCD) number
				GTD string `json:"gtd"`

				// Mandatory “Chestny ZNAK” labeling
				MandatoryMark string `json:"mandatory_mark"`

				// Check result. true if the labeling code of product item meets the requirements
				Valid bool `json:"valid"`

				// Product batch registration number
				RNPT string `json:"rnpt"`
			} `json:"exemplars"`

			// Product identifier
			ProductId int64 `json:"product_id"`

			// Check result. true if the labeling codes of all product items meet the requirements
			Valid bool `json:"valid"`
		} `json:"products"`
	} `json:"result"`
}

// Method for checking whether labeling codes meet the "Chestny ZNAK" system requirements on length and symbols.
//
// If you don't have the customs cargo declaration (CCD) number, you don't have to specify it
func (c FBS) ValidateLabelingCodes(params *ValidateLabelingCodesParams) (*ValidateLabelingCodesResponse, error) {
	url := "/v4/fbs/posting/product/exemplar/validate"

	resp := &ValidateLabelingCodesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentDataByBarcodeParams struct {
	// Shipment barcode
	Barcode string `json:"barcode"`
}

type GetShipmentDataByBarcodeResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Analytical data
		AnalyticsData struct {
			// Delivery city
			City string `json:"city"`

			// Delivery method
			DeliveryType string `json:"delivery_type"`

			// Indication that the recipient is a legal entity:
			//   - true — a legal entity
			//   - false — a natural person
			IsLegal bool `json:"is_legal"`

			// Premium subscription availability
			IsPremium bool `json:"is_premium"`

			// Payment method
			PaymentTypeGroupName string `json:"payment_type_group_name"`

			// Delivery region
			Region string `json:"region"`
		} `json:"analytics_data"`

		// Shipment barcodes
		Barcodes FBSBarcode `json:"barcodes"`

		// Cancellation reason identifier
		CancelReasonId int64 `json:"cancel_reason_id"`

		// Date and time when the shipment was created
		CreatedAt time.Time `json:"created_at"`

		// Financial data
		FinancialData struct {
			// Identifier of the cluster, where the shipment is sent from
			ClusterFrom string `json:"cluster_from"`

			// Identifier of the cluster, where the shipment is delivered to
			ClusterTo string `json:"cluster_to"`

			// Services
			PostingServices []MarketplaceServices `json:"posting_services"`

			// Products list
			Products []FinancialDataProduct `json:"products"`
		} `json:"financial_data"`

		// Start date and time of shipment processing
		InProcessAt time.Time `json:"in_process_at"`

		// Order identifier to which the shipment belongs
		OrderId int64 `json:"order_id"`

		// Order number to which the shipment belongs
		OrderNumber string `json:"order_number"`

		// Shipment number
		PostingNumber string `json:"posting_number"`

		// List of products in the shipment
		Products []PostingProduct `json:"products"`

		// Date and time before which the shipment must be packaged.
		// If the shipment is not packaged by this date, it will be canceled automatically
		ShipmentDate time.Time `json:"shipment_date"`

		// Shipment status
		Status string `json:"status"`
	} `json:"result"`
}

// Method for getting shipments data by barcode
func (c FBS) GetShipmentDataByBarcode(params *GetShipmentDataByBarcodeParams) (*GetShipmentDataByBarcodeResponse, error) {
	url := "/v2/posting/fbs/get-by-barcode"

	resp := &GetShipmentDataByBarcodeResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentDataByIdentifierParams struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`

	// Additional fields that should be added to the response
	With GetShipmentDataByIdentifierWith `json:"with"`
}

type GetShipmentDataByIdentifierWith struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Add data on products and their instances to the response
	ProductExemplars bool `json:"product_exemplars"`

	// Add related shipment numbers to the response.
	// Related shipments are ones into which the parent shipment was split during packaging
	RelatedPostings bool `json:"related_postings"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type GetShipmentDataByIdentifierResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Additional Data Key-Value
		AdditionalData []struct {
			// Key
			Key string `json:"key"`

			// value
			Value string `json:"value"`
		} `json:"additional_data"`

		// Recipient details
		Addressee struct {
			// Recipient name
			Name string `json:"name"`

			// Recipient phone number
			Phone string `json:"phone"`
		} `json:"addressee"`

		// Analytics data
		AnalyticsData struct {
			// Delivery city
			City string `json:"city"`

			// Delivery start date and time
			DeliveryDateBegin time.Time `json:"delivery_date_begin"`

			// Delivery end date and time
			DeliveryDateEnd time.Time `json:"delivery_date_end"`

			// Delivery method
			DeliveryType string `json:"delivery_type"`

			// Indication that the recipient is a legal entity:
			//   - true — a legal entity,
			//   - false — a natural person
			IsLegal bool `json:"is_legal"`

			// Premium subscription availability
			IsPremium bool `json:"is_premium"`

			// Payment method
			PaymentTypeGroupName string `json:"payment_type_group_name"`

			// Delivery region
			Region string `json:"region"`

			// Delivery service
			TPLProvider string `json:"tpl_provider"`

			// Delivery service identifier
			TPLProviderId int64 `json:"tpl_provider_id"`

			// Order shipping warehouse name
			Warehouse string `json:"warehouse"`

			// Warehouse identifier
			WarehouseId int64 `json:"warehouse_id"`
		} `json:"analytics_data"`

		// Shipment barcodes
		Barcodes FBSBarcode `json:"barcodes"`

		// Cancellation details
		Cancellation FBSCancellation `json:"calcellation"`

		// Courier information
		Courier struct {
			// Car model
			CarModel string `json:"car_model"`

			// Car number
			CarNumber string `json:"car_number"`

			// Courier's full name
			Name string `json:"name"`

			// Courier's phone number
			Phone string `json:"phone"`
		} `json:"courier"`

		// Customer details
		Customer FBSCustomer `json:"customer"`

		// Date when the shipment was transferred for delivery
		DeliveringDate time.Time `json:"delivering_date"`

		// Delivery method
		DeliveryMethod FBSDeliveryMethod `json:"delivery_method"`

		// Delivery cost
		DeliveryPrice string `json:"delivery_type"`

		// Data on the product cost, discount amount, payout and commission
		FinancialData FBSFinancialData `json:"financial_date"`

		// Start date and time of shipment processing
		InProcessAt time.Time `json:"in_process_at"`

		// If Ozon Express fast delivery was used—true
		IsExpress bool `json:"is_express"`

		// Indication that there is a multi-box product in the shipment and you need to pass the number of boxes for it:
		//   - true — before packaging pass the number of boxes using the /v3/posting/multiboxqty/set method.
		//   - false — you packed the shipment specifying the number of boxes in the multi_box_qty parameter, or there is no multi-box product in the shipment
		IsMultibox bool `json:"is_multibox"`

		// Number of boxes in which the product is packed
		MultiBoxQuantity int32 `json:"multi_box_qty"`

		// Order identifier to which the shipment belongs
		OrderId int64 `json:"order_id"`

		// Order number to which the shipment belongs
		OrderNumber string `json:"order_number"`

		// Number of the parent shipment which split resulted in the current shipment
		ParentPostingNumber string `json:"parent_posting_number"`

		// Shipment number
		PostingNumber string `json:"posting_number"`

		// Information on products and their instances.
		//
		// The response contains the field product_exemplars, if the attribute with.product_exemplars = true is passed in the request
		ProductExemplars struct {
			// Products
			Products []struct {
				// Product identifier in the Ozon system, SKU
				SKU int64 `json:"sku"`

				// Array of exemplars
				Exemplars []struct {
					// Mandatory “Chestny ZNAK” labeling
					MandatoryMark string `json:"mandatory_mark"`

					// Сustoms cargo declaration (CCD) number
					GTD string `json:"gtd"`

					// Indication that a сustoms cargo declaration (CCD) number hasn't been specified
					IsGTDAbsest bool `json:"is_gtd_absent"`

					// Product batch registration number
					RNPT string `json:"rnpt"`

					// Indication that a product batch registration number hasn't been specified
					IsRNPTAbsent bool `json:"is_rnpt_absent"`
				} `json:"exemplars"`
			} `json:"products"`
		} `json:"product_exemplars"`

		// Array of products in the shipment
		Products []struct {
			PostingProduct

			// Product dimensions
			Dimensions struct {
				// Package height
				Height string `json:"height"`

				// Product length
				Length string `json:"length"`

				// Weight of product in the package
				Weight string `json:"weight"`

				// Package width
				Width string `json:"width"`
			} `json:"dimensions"`
		} `json:"products"`

		// Delivery service status
		ProviderStatus string `json:"provider_status"`

		// Related shipments
		RelatedPostings struct {
			RelatedPostingNumbers []string `json:"related_posting_numbers"`
		} `json:"related_postings"`

		// Array of Ozon Product IDs (SKU) for which you need to pass the customs cargo declaration (CCD) number, the manufacturing country,
		// product batch registration number, or "Chestny ZNAK" labeling to change the shipment status to the next one
		Requirements FBSRequirements `json:"requirements"`

		// Date and time before which the shipment must be packaged.
		// If the shipment is not packaged by this date, it will be canceled automatically
		ShipmentDate time.Time `json:"shipment_date"`

		// Shipment status
		Status string `json:"status"`

		// Type of integration with the delivery service:
		//   - ozon — delivery by the Ozon logistics.
		//   - aggregator — delivery by a third-party service, Ozon registers the order.
		//   - 3pl_tracking — delivery by a third-party service, the seller registers the order.
		//   - non_integrated — delivery by the seller
		TPLIntegrationType string `json:"tpl_integration_type"`

		// Shipment tracking number
		TrackingNumber string `json:"tracking_number"`
	} `json:"result"`
}

// Method for getting shipment details by identifier
func (c FBS) GetShipmentDataByIdentifier(params *GetShipmentDataByIdentifierParams) (*GetShipmentDataByIdentifierResponse, error) {
	url := "/v3/posting/fbs/get"

	resp := &GetShipmentDataByIdentifierResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
