package freshservice

import (
	"fmt"
	"time"
)

// ServiceCatalogDetails holds the details for a specific Freshservice service catalog item
type ServiceCatalogDetails struct {
	ID                     int               `json:"id"`
	CreatedAt              time.Time         `json:"created_at"`
	UpdatedAt              time.Time         `json:"updated_at"`
	Name                   string            `json:"name"`
	DeliveryTime           int               `json:"delivery_time"`
	DisplayID              int               `json:"display_id"`
	CategoryID             int               `json:"category_id"`
	ProductID              int               `json:"product_id"`
	Quantity               int               `json:"quantity"`
	Deleted                bool              `json:"deleted"`
	IconName               string            `json:"icon_name"`
	GroupVisibility        int               `json:"group_visibility"`
	ItemType               int               `json:"item_type"`
	CiTypeID               int               `json:"ci_type_id"`
	CostVisibility         bool              `json:"cost_visibility"`
	DeliveryTimeVisibility bool              `json:"delivery_time_visibility"`
	Configs                map[string]string `json:"configs"`
	Botified               bool              `json:"botified"`
	Visibility             int               `json:"visibility"`
	AllowAttachments       bool              `json:"allow_attachments"`
	AllowQuantity          bool              `json:"allow_quantity"`
	IsBundle               bool              `json:"is_bundle"`
	CreateChild            bool              `json:"create_child"`
	Description            string            `json:"description"`
	ShortDescription       string            `json:"short_description"`
	Cost                   string            `json:"cost"`
	CustomFields           []interface{}     `json:"custom_fields"`
	ChildItems             []interface{}     `json:"child_items"`
}

// ServiceCategories represents service catalog item categories in Freshservice
type ServiceCategories struct {
	List []ServiceCategory `json:"service_categories"`
}

// ServiceCategory represents a category assigned to a service catalog item in Freshservice
type ServiceCategory struct {
	Description string    `json:"description"`
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Position    int       `json:"position"`
}

// ServiceCatalogItemListFilter are the available filter options
// for a service catalog API list request
type ServiceCatalogItemListFilter struct {
	CatalogID int
}

// QueryString allows the available filter items to meet the QueryFilter interface
func (scf *ServiceCatalogItemListFilter) QueryString() string {
	return fmt.Sprintf("category_id=%d", scf.CatalogID)
}

// ServiceRequestOptions options for creating a new Service Request
type ServiceRequestOptions struct {
	Quantity     int          `json:"quantity"`
	Email        string       `json:"email"`
	CustomFields CustomFields `json:"custom_fields"`
	ChildItems   []ChildItem  `json:"child_items"`
}

type ChildItem struct {
	Quantity      int          `json:"quantity"`
	ServiceItemID int          `json:"service_item_id"`
	Email         string       `json:"email"`
	CustomFields  CustomFields `json:"custom_fields"`
}

// ServiceRequestResponse is the response from a new service request API call
type ServiceRequestResponse struct {
	ServiceRequest TicketDetails `json:"service_request"`
}
