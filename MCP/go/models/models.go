package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// File represents the File schema from the OpenAPI specification
type File struct {
	TypeField string `json:"type,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Url string `json:"url"`
	Content_type string `json:"content_type,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
}

// Listing represents the Listing schema from the OpenAPI specification
type Listing struct {
	Third_party_integration bool `json:"third_party_integration,omitempty"`
	Id string `json:"id,omitempty"`
	Blendr_id string `json:"blendr_id,omitempty"`
	Meta_tag_title string `json:"meta_tag_title,omitempty"`
	Segment_id string `json:"segment_id,omitempty"`
	Card_background_color string `json:"card_background_color,omitempty"`
	Description string `json:"description,omitempty"`
	Detail_page_disabled bool `json:"detail_page_disabled,omitempty"`
	Slug string `json:"slug"`
	Automate_id string `json:"automate_id,omitempty"`
	Collections []Collection `json:"collections,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
	Third_party_integration_link string `json:"third_party_integration_link,omitempty"`
	Published bool `json:"published,omitempty"`
	Logo Logo `json:"logo,omitempty"`
	Products []Product `json:"products,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Integromat_id string `json:"integromat_id,omitempty"`
	Tag_line string `json:"tag_line,omitempty"`
	Name string `json:"name"`
	Upcoming bool `json:"upcoming,omitempty"`
	Media []Media `json:"media,omitempty"`
	Cloud_service_id string `json:"cloud_service_id,omitempty"`
	Partner Partner `json:"partner,omitempty"`
	Screenshots []Screenshot `json:"screenshots,omitempty"`
	Native_integration_link string `json:"native_integration_link,omitempty"`
	Piesync_id string `json:"piesync_id,omitempty"`
	Sticky bool `json:"sticky,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Features string `json:"features,omitempty"`
	Pricing string `json:"pricing,omitempty"`
	Native_integration bool `json:"native_integration,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Unify_connector_id string `json:"unify_connector_id,omitempty"`
	Tray_io_id string `json:"tray_io_id,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	Combidesk_id string `json:"combidesk_id,omitempty"`
	Meta_tag_description string `json:"meta_tag_description,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Zapier_id string `json:"zapier_id,omitempty"`
	Meta_tag_keywords string `json:"meta_tag_keywords,omitempty"`
	Card_background_image File `json:"card_background_image,omitempty"`
	Microsoft_flow_id string `json:"microsoft_flow_id,omitempty"`
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Count int `json:"count,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
	Listing_description_text_template string `json:"listing_description_text_template,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name"`
	Id string `json:"id,omitempty"`
	Listing_features_text_template string `json:"listing_features_text_template,omitempty"`
	Listing_pricing_text_template string `json:"listing_pricing_text_template,omitempty"`
	Logo Logo `json:"logo,omitempty"`
	Slug string `json:"slug"`
	Updated_at string `json:"updated_at,omitempty"`
}

// MetaTagSettings represents the MetaTagSettings schema from the OpenAPI specification
type MetaTagSettings struct {
	Description string `json:"description,omitempty"`
	Description_category_page string `json:"description_category_page,omitempty"`
	Description_collection_page string `json:"description_collection_page,omitempty"`
	Description_listing_page string `json:"description_listing_page,omitempty"`
	Keywords string `json:"keywords,omitempty"`
	Title string `json:"title,omitempty"`
	Title_postfix string `json:"title_postfix,omitempty"`
}

// Collection represents the Collection schema from the OpenAPI specification
type Collection struct {
	Show_max_items_homepage int `json:"show_max_items_homepage,omitempty"`
	Slug string `json:"slug"`
	Logo File `json:"logo,omitempty"`
	Card_background_color string `json:"card_background_color,omitempty"`
	Sequence int `json:"sequence,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Translations []Translation `json:"translations,omitempty"`
	Card_style string `json:"card_style,omitempty"`
	Card_background_image File `json:"card_background_image,omitempty"`
	Count int `json:"count,omitempty"`
	Visible bool `json:"visible"`
	Card_columns int `json:"card_columns,omitempty"`
	Description string `json:"description,omitempty"`
	Hidden_from_homepage bool `json:"hidden_from_homepage,omitempty"`
}

// CardSettings represents the CardSettings schema from the OpenAPI specification
type CardSettings struct {
	Background_color string `json:"background_color,omitempty"`
	Color string `json:"color,omitempty"`
	Description_lines int `json:"description_lines,omitempty"`
	Icon_shadow_enabled bool `json:"icon_shadow_enabled,omitempty"`
	Shadow_enabled bool `json:"shadow_enabled,omitempty"`
	Show_description bool `json:"show_description,omitempty"`
	Border_radius string `json:"border_radius,omitempty"`
	Show_action bool `json:"show_action,omitempty"`
	Show_category bool `json:"show_category,omitempty"`
	Columns int `json:"columns,omitempty"`
	Show_badges bool `json:"show_badges,omitempty"`
	Border_color string `json:"border_color,omitempty"`
	Border_size string `json:"border_size,omitempty"`
	Icon_size int `json:"icon_size,omitempty"`
	Icon_border_radius string `json:"icon_border_radius,omitempty"`
	Style string `json:"style,omitempty"`
}

// Screenshot represents the Screenshot schema from the OpenAPI specification
type Screenshot struct {
	File File `json:"file,omitempty"`
	Id string `json:"id,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
	Caption string `json:"caption,omitempty"`
}

// GetProductResponse represents the GetProductResponse schema from the OpenAPI specification
type GetProductResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Product `json:"data"`
}

// Media represents the Media schema from the OpenAPI specification
type Media struct {
	Translations []Translation `json:"translations,omitempty"`
	TypeField string `json:"type,omitempty"`
	Url string `json:"url"`
	Caption string `json:"caption,omitempty"`
	Id string `json:"id,omitempty"`
	Sequence int `json:"sequence,omitempty"`
}

// CustomSettings represents the CustomSettings schema from the OpenAPI specification
type CustomSettings struct {
	Css string `json:"css,omitempty"`
	Css_link string `json:"css_link,omitempty"`
	Domain string `json:"domain,omitempty"`
	Html_footer string `json:"html_footer,omitempty"`
	Html_nav string `json:"html_nav,omitempty"`
	Java_script_link string `json:"java_script_link,omitempty"`
}

// CTASettings represents the CTASettings schema from the OpenAPI specification
type CTASettings struct {
	Button_label string `json:"button_label,omitempty"`
	Color string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
	Background_color string `json:"background_color,omitempty"`
	Button_background_color string `json:"button_background_color,omitempty"`
	Button_color string `json:"button_color,omitempty"`
	Button_link string `json:"button_link,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// GetListingResponse represents the GetListingResponse schema from the OpenAPI specification
type GetListingResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Listing `json:"data"`
}

// GetEcosystemResponse represents the GetEcosystemResponse schema from the OpenAPI specification
type GetEcosystemResponse struct {
	Data Ecosystem `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// IntegrationSettings represents the IntegrationSettings schema from the OpenAPI specification
type IntegrationSettings struct {
	Iubenda_cookie_policy_id string `json:"iubenda_cookie_policy_id,omitempty"`
	Segment_enabled bool `json:"segment_enabled,omitempty"`
	Livechat_id string `json:"livechat_id,omitempty"`
	Microsoft_flow_enabled bool `json:"microsoft_flow_enabled,omitempty"`
	Blendr_enabled bool `json:"blendr_enabled,omitempty"`
	Google_analytics_id string `json:"google_analytics_id,omitempty"`
	Iubenda_site_id string `json:"iubenda_site_id,omitempty"`
	Zapier_id string `json:"zapier_id,omitempty"`
	Heap_id string `json:"heap_id,omitempty"`
	Hubspot_portal_id string `json:"hubspot_portal_id,omitempty"`
	Combidesk_enabled bool `json:"combidesk_enabled,omitempty"`
	Crisp_id string `json:"crisp_id,omitempty"`
	Piesync_enabled bool `json:"piesync_enabled,omitempty"`
	Intercom_app_id string `json:"intercom_app_id,omitempty"`
	Journy_io_id string `json:"journy_io_id,omitempty"`
	Drift_id string `json:"drift_id,omitempty"`
	Automate_enabled bool `json:"automate_enabled,omitempty"`
	Google_tag_manager_id string `json:"google_tag_manager_id,omitempty"`
	Integromat_enabled bool `json:"integromat_enabled,omitempty"`
	Segment_id string `json:"segment_id,omitempty"`
	Zapier_enabled bool `json:"zapier_enabled,omitempty"`
	Journy_io_domain string `json:"journy_io_domain,omitempty"`
	Microsoft_flow_id string `json:"microsoft_flow_id,omitempty"`
	Onetrust_id string `json:"onetrust_id,omitempty"`
	Tray_io_enabled bool `json:"tray_io_enabled,omitempty"`
	Zapier_beta_link string `json:"zapier_beta_link,omitempty"`
	Metomic_id string `json:"metomic_id,omitempty"`
	Albacross_id string `json:"albacross_id,omitempty"`
}

// LeadFormSettings represents the LeadFormSettings schema from the OpenAPI specification
type LeadFormSettings struct {
	Work_email_validation bool `json:"work_email_validation,omitempty"`
	Last_name_field_enabled bool `json:"last_name_field_enabled,omitempty"`
	Last_name_field_required bool `json:"last_name_field_required,omitempty"`
	Capture_form_enabled bool `json:"capture_form_enabled,omitempty"`
	Telephone_field_required bool `json:"telephone_field_required,omitempty"`
	First_name_field_required bool `json:"first_name_field_required,omitempty"`
	First_name_field_enabled bool `json:"first_name_field_enabled,omitempty"`
	Integration_enabled bool `json:"integration_enabled,omitempty"`
	Telephone_field_enabled bool `json:"telephone_field_enabled,omitempty"`
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// Translation represents the Translation schema from the OpenAPI specification
type Translation struct {
	Key string `json:"key"`
	Locale string `json:"locale"`
	Value string `json:"value,omitempty"`
}

// ListingSettings represents the ListingSettings schema from the OpenAPI specification
type ListingSettings struct {
	Native_integration_link string `json:"native_integration_link,omitempty"`
	Pricing_text_template string `json:"pricing_text_template,omitempty"`
	Sidebar_position string `json:"sidebar_position,omitempty"`
	Features_title string `json:"features_title,omitempty"`
	Naming string `json:"naming,omitempty"`
	Pricing_disabled bool `json:"pricing_disabled,omitempty"`
	Pricing_title string `json:"pricing_title,omitempty"`
	Description_text_template string `json:"description_text_template,omitempty"`
	Description_title string `json:"description_title,omitempty"`
	Features_text_template string `json:"features_text_template,omitempty"`
	Install_button_label string `json:"install_button_label,omitempty"`
	Name_postfix string `json:"name_postfix,omitempty"`
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Last_name string `json:"last_name,omitempty"`
	Linked_in string `json:"linked_in,omitempty"`
	Name string `json:"name"`
	Role string `json:"role,omitempty"`
	Twitter string `json:"twitter,omitempty"`
	Email string `json:"email,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Id string `json:"id,omitempty"`
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	Id string `json:"id,omitempty"`
	Logo File `json:"logo,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
	Count int `json:"count,omitempty"`
	Name string `json:"name"`
	Visible bool `json:"visible"`
	Sequence int `json:"sequence,omitempty"`
	Description string `json:"description,omitempty"`
	Slug string `json:"slug"`
}

// GetCollectionResponse represents the GetCollectionResponse schema from the OpenAPI specification
type GetCollectionResponse struct {
	Data Collection `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Ecosystem represents the Ecosystem schema from the OpenAPI specification
type Ecosystem struct {
	Primary_color string `json:"primary_color,omitempty"`
	Collections_count_badge bool `json:"collections_count_badge,omitempty"`
	Footer_color string `json:"footer_color,omitempty"`
	Footer_background_color string `json:"footer_background_color,omitempty"`
	Is_published bool `json:"is_published"`
	Terms_link string `json:"terms_link,omitempty"`
	Website string `json:"website,omitempty"`
	Shadow_pages_enabled bool `json:"shadow_pages_enabled,omitempty"`
	Privacy_link string `json:"privacy_link,omitempty"`
	Zaps_page_enabled bool `json:"zaps_page_enabled,omitempty"`
	Home_page_collection_category_cards bool `json:"home_page_collection_category_cards,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	Attribution bool `json:"attribution,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Body_button_color string `json:"body_button_color,omitempty"`
	Google_site_verification_id string `json:"google_site_verification_id,omitempty"`
	Show_requested_listings bool `json:"show_requested_listings,omitempty"`
	Utm_campaign string `json:"utm_campaign,omitempty"`
	Cta_settings CTASettings `json:"cta_settings,omitempty"`
	Detail_pages_enabled bool `json:"detail_pages_enabled,omitempty"`
	Name string `json:"name"`
	Body_button_background_color string `json:"body_button_background_color,omitempty"`
	Shadow_page_description string `json:"shadow_page_description,omitempty"`
	Show_attribution_badge bool `json:"show_attribution_badge,omitempty"`
	Navigation_background_color string `json:"navigation_background_color,omitempty"`
	Categories_count_badge bool `json:"categories_count_badge,omitempty"`
	Integration_settings IntegrationSettings `json:"integration_settings,omitempty"`
	Request_link string `json:"request_link,omitempty"`
	Create_link string `json:"create_link,omitempty"`
	Navigation_sticky bool `json:"navigation_sticky,omitempty"`
	Custom_settings CustomSettings `json:"custom_settings,omitempty"`
	Unify_application_id string `json:"unify_application_id,omitempty"`
	About string `json:"about,omitempty"`
	Body_color string `json:"body_color,omitempty"`
	Categories_show_max_items int `json:"categories_show_max_items,omitempty"`
	Total_published_listings int `json:"total_published_listings,omitempty"`
	Navigation_logo_post_fix string `json:"navigation_logo_post_fix,omitempty"`
	Body_background_color string `json:"body_background_color,omitempty"`
	Custom_domain string `json:"custom_domain,omitempty"`
	Home_page_show_all_listings bool `json:"home_page_show_all_listings,omitempty"`
	Navigation_mobile_menu_type string `json:"navigation_mobile_menu_type,omitempty"`
	Alternatives_background_color string `json:"alternatives_background_color,omitempty"`
	Menu_position string `json:"menu_position,omitempty"`
	Masthead_settings MastheadSettings `json:"masthead_settings,omitempty"`
	Menu_style string `json:"menu_style,omitempty"`
	Card_settings CardSettings `json:"card_settings,omitempty"`
	Slug string `json:"slug"`
	Hide_install_buttons bool `json:"hide_install_buttons,omitempty"`
	Navigation_color string `json:"navigation_color,omitempty"`
	Lead_form_settings LeadFormSettings `json:"lead_form_settings,omitempty"`
	Zaps_menu_title string `json:"zaps_menu_title,omitempty"`
	Installation_request_flow_enabled bool `json:"installation_request_flow_enabled,omitempty"`
	Collections_title string `json:"collections_title,omitempty"`
	Listing_settings ListingSettings `json:"listing_settings,omitempty"`
	Meta_tag_settings MetaTagSettings `json:"meta_tag_settings,omitempty"`
	Body_link_color string `json:"body_link_color,omitempty"`
	Alternatives_color string `json:"alternatives_color,omitempty"`
}

// GetCategoryResponse represents the GetCategoryResponse schema from the OpenAPI specification
type GetCategoryResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Category `json:"data"`
}

// GetCollectionsResponse represents the GetCollectionsResponse schema from the OpenAPI specification
type GetCollectionsResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Collection `json:"data"`
}

// MastheadSettings represents the MastheadSettings schema from the OpenAPI specification
type MastheadSettings struct {
	Description string `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
	Background string `json:"background,omitempty"`
	Background_color string `json:"background_color,omitempty"`
	Color string `json:"color,omitempty"`
	Columns int `json:"columns,omitempty"`
}

// Logo represents the Logo schema from the OpenAPI specification
type Logo struct {
	Id string `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Url string `json:"url"`
	Content_type string `json:"content_type,omitempty"`
}

// Partner represents the Partner schema from the OpenAPI specification
type Partner struct {
	Twitter string `json:"twitter,omitempty"`
	Website string `json:"website,omitempty"`
	Id string `json:"id,omitempty"`
	Listed string `json:"listed,omitempty"`
	Company string `json:"company"`
	Contacts []Contact `json:"contacts,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Icon File `json:"icon,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

// GetCategoriesResponse represents the GetCategoriesResponse schema from the OpenAPI specification
type GetCategoriesResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Category `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
}

// GetListingsResponse represents the GetListingsResponse schema from the OpenAPI specification
type GetListingsResponse struct {
	Data []Listing `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetProductsResponse represents the GetProductsResponse schema from the OpenAPI specification
type GetProductsResponse struct {
	Data []Product `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}
