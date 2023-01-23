package pv_types

var (
	PV_URL = "https://api.peoplevine.co.uk/customer.asmx/customerLogin"
)

// *********************************************************************
// *** Peoplevine auth request
// *********************************************************************
type TPVRequest struct {
	Auth struct {
		ApiUsername string `json:"api_username"`
		ApiPassword string `json:"api_password"`
		ApiKey      string `json:"api_key"`
		CompanyNo   int    `json:"company_no"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	} `json:"auth"`
	Credentials struct {
		CompanyNo int    `json:"company_no"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	} `json:"credentials"`
}

// *********************************************************************
// *** Peoplevine response summary data structure
// *********************************************************************

type TPVDData struct {
	D string `json:"d"`
}

// *********************************************************************
// *** Peoplevine response with header + detail
// *********************************************************************

type TPVDataWithDetail struct {
	ResponseCode    string `json:"responseCode"`
	Message         string `json:"message"`
	IsError         bool   `json:"isError"`
	MethodFailed    string `json:"methodFailed"`
	ExtendedMessage string `json:"extendedMessage"`
	Reason          string `json:"reason"`
	ReturnObject    struct {
		AppInstalledOn         string        `json:"app_installed_on"`
		AppInstalled           bool          `json:"app_installed"`
		NetworkCompanyNo       int           `json:"network_company_no"`
		NetworkCompanyName     string        `json:"network_company_name"`
		AdminOverride          bool          `json:"adminOverride"`
		MapImage               string        `json:"map_image"`
		LifecycleStage         string        `json:"lifecycle_stage"`
		IpAddress              string        `json:"ip_address"`
		MemberDirectoryEnabled bool          `json:"memberDirectoryEnabled"`
		EnableHouseAccount     bool          `json:"enableHouseAccount"`
		KeepPrivate            bool          `json:"keep_private"`
		CustomerToken          string        `json:"customer_token"`
		MemberDiscount         int           `json:"member_discount"`
		Name                   string        `json:"name"`
		CcEmail                string        `json:"cc_email"`
		IsMember               bool          `json:"isMember"`
		WasMember              bool          `json:"wasMember"`
		HasApplied             bool          `json:"hasApplied"`
		MemberSince            string        `json:"member_since"`
		WhenApplied            string        `json:"when_applied"`
		UpdateByReference      bool          `json:"updateByReference"`
		Addresses              []interface{} `json:"addresses"`
		Email2                 string        `json:"email2"`
		Notes                  []interface{} `json:"notes"`
		Gender                 string        `json:"gender"`
		CompanyTitle           string        `json:"company_title"`
		AnniversaryDate        string        `json:"anniversary_date"`
		HasPassword            bool          `json:"hasPassword"`
		LastLogin              string        `json:"last_login"`
		LastPageView           string        `json:"lastPageView"`
		LastEmailOpened        string        `json:"lastEmailOpened"`
		LastNewsletter         string        `json:"lastNewsletter"`
		LastTouchpoint         string        `json:"lastTouchpoint"`
		LastCommunication      string        `json:"lastCommunication"`
		TotalTouchpoints       int           `json:"totalTouchpoints"`
		User                   struct {
			BusinessLocationNo int    `json:"business_location_no"`
			EmployeeId         string `json:"employee_id"`
			UserNo             int    `json:"user_no"`
			FirstName          string `json:"first_name"`
			LastName           string `json:"last_name"`
			MobileNumber       string `json:"mobile_number"`
			Email              string `json:"email"`
			Username           string `json:"username"`
			Status             string `json:"status"`
		} `json:"user"`
		UserModified struct {
			BusinessLocationNo int    `json:"business_location_no"`
			EmployeeId         string `json:"employee_id"`
			UserNo             int    `json:"user_no"`
			FirstName          string `json:"first_name"`
			LastName           string `json:"last_name"`
			MobileNumber       string `json:"mobile_number"`
			Email              string `json:"email"`
			Username           string `json:"username"`
			Status             string `json:"status"`
		} `json:"user_modified"`
		UserLastViewed struct {
			BusinessLocationNo int    `json:"business_location_no"`
			EmployeeId         string `json:"employee_id"`
			UserNo             int    `json:"user_no"`
			FirstName          string `json:"first_name"`
			LastName           string `json:"last_name"`
			MobileNumber       string `json:"mobile_number"`
			Email              string `json:"email"`
			Username           string `json:"username"`
			Status             string `json:"status"`
		} `json:"user_last_viewed"`
		LastViewed          string        `json:"last_viewed"`
		LastViewedBy        int           `json:"last_viewed_by"`
		ModifiedBy          int           `json:"modified_by"`
		AffiliateName       interface{}   `json:"affiliate_name"`
		AffiliateNo         int           `json:"affiliate_no"`
		SessionId           string        `json:"session_id"`
		PasswordExpires     string        `json:"password_expires"`
		Flag                string        `json:"flag"`
		VipFlag             bool          `json:"vip_flag"`
		LastContacted       string        `json:"last_contacted"`
		KloutScore          int           `json:"klout_score"`
		LinkedinHandle      string        `json:"linkedin_handle"`
		LinkedinId          string        `json:"linkedin_id"`
		LifecycleStageNo    int           `json:"lifecycle_stage_no"`
		CustomerSource      string        `json:"customer_source"`
		Timezone            string        `json:"timezone"`
		Username            string        `json:"username"`
		Settings            string        `json:"settings"`
		GeneratePassword    bool          `json:"generatePassword"`
		ProfilePhoto        string        `json:"profile_photo"`
		FoursquareHandle    string        `json:"foursquare_handle"`
		FoursquareId        string        `json:"foursquare_id"`
		CustomerStatus      string        `json:"customer_status"`
		Attributes          []interface{} `json:"attributes"`
		FacebookHandle      string        `json:"facebook_handle"`
		TwitterHandle       string        `json:"twitter_handle"`
		TiktokHandle        string        `json:"tiktok_handle"`
		InstagramHandle     string        `json:"instagram_handle"`
		Latitude            string        `json:"latitude"`
		Longitude           string        `json:"longitude"`
		CustomerNo          int           `json:"customer_no"`
		FirstName           string        `json:"first_name"`
		LastName            string        `json:"last_name"`
		Address             string        `json:"address"`
		Address2            string        `json:"address2"`
		City                string        `json:"city"`
		State               string        `json:"state"`
		ZipCode             string        `json:"zip_code"`
		Country             string        `json:"country"`
		PhoneNumber         string        `json:"phone_number"`
		MobileNumber        string        `json:"mobile_number"`
		Email               string        `json:"email"`
		Password            string        `json:"password"`
		CreatedOn           string        `json:"created_on"`
		DeviceType          string        `json:"device_type"`
		PassbookEnabled     string        `json:"passbook_enabled"`
		WalletEnabled       string        `json:"wallet_enabled"`
		OptInEmail          string        `json:"opt_in_email"`
		OptInSms            string        `json:"opt_in_sms"`
		OptInUpdates        string        `json:"opt_in_updates"`
		CompanyNo           int           `json:"company_no"`
		UserNo              int           `json:"user_no"`
		CompanyName         string        `json:"company_name"`
		Birthdate           string        `json:"birthdate"`
		ModifiedOn          string        `json:"modified_on"`
		Address3            string        `json:"address3"`
		CustomerReference   string        `json:"customer_reference"`
		FacebookId          string        `json:"facebook_id"`
		TwitterId           string        `json:"twitter_id"`
		TiktokId            string        `json:"tiktok_id"`
		InstagramId         string        `json:"instagram_id"`
		MiddleName          string        `json:"middle_name"`
		Website             string        `json:"website"`
		CustomerType        string        `json:"customer_type"`
		Schema              interface{}   `json:"schema"`
		TotalFavorite       int           `json:"totalFavorite"`
		SkipTriggers        bool          `json:"skipTriggers"`
		IncludeObjects      bool          `json:"includeObjects"`
		IncludeStats        bool          `json:"includeStats"`
		IncludeTotalRecords bool          `json:"includeTotalRecords"`
		TotalRecords        int           `json:"totalRecords"`
		RowsPerPage         int           `json:"rows_per_page"`
		PageNumber          int           `json:"page_number"`
		UpdateFields        string        `json:"updateFields"`
		IncludeCompany      bool          `json:"includeCompany"`
		Lightweight         bool          `json:"lightweight"`
		ItemDetailsFields   interface{}   `json:"itemDetailsFields"`
		IncludeItemDetails  bool          `json:"includeItemDetails"`
		SearchValue         interface{}   `json:"search_value"`
		SortBy              interface{}   `json:"sort_by"`
		ReturnTotal         int           `json:"return_total"`
		FromDate            string        `json:"from_date"`
		ToDate              string        `json:"to_date"`
		IncludeDetails      bool          `json:"includeDetails"`
		TimezoneOffset      int           `json:"timezone_offset"`
		TimezoneId          interface{}   `json:"timezone_id"`
	} `json:"returnObject"`
	ResponseTime    float64     `json:"responseTime"`
	TotalRecords    int         `json:"totalRecords"`
	MinPrice        int         `json:"minPrice"`
	MaxPrice        int         `json:"maxPrice"`
	PermissionLevel interface{} `json:"permission_level"`
	TimezoneOffset  int         `json:"timezone_offset"`
	TimezoneId      interface{} `json:"timezone_id"`
}

// *********************************************************************
// *** Peoplevine response header only
// *********************************************************************

type TPVDataHeader struct {
	ResponseCode    string      `json:"responseCode"`
	Message         string      `json:"message"`
	IsError         bool        `json:"isError"`
	MethodFailed    string      `json:"methodFailed"`
	ExtendedMessage string      `json:"extendedMessage"`
	Reason          string      `json:"reason"`
	ReturnObject    interface{} `json:"returnObject"`
	ResponseTime    float64     `json:"responseTime"`
	TotalRecords    int         `json:"totalRecords"`
	MinPrice        int         `json:"minPrice"`
	MaxPrice        int         `json:"maxPrice"`
	PermissionLevel interface{} `json:"permission_level"`
	TimezoneOffset  int         `json:"timezone_offset"`
	TimezoneId      interface{} `json:"timezone_id"`
}

// *********************************************************************
// *** Peoplevine response with (D) root tag
// *********************************************************************
type TPVSumData struct {
	D struct {
		ResponseCode    string `json:"responseCode"`
		Message         string `json:"message"`
		IsError         bool   `json:"isError"`
		MethodFailed    string `json:"methodFailed"`
		ExtendedMessage string `json:"extendedMessage"`
		Reason          string `json:"reason"`
		ReturnObject    struct {
			AppInstalledOn         string        `json:"app_installed_on"`
			AppInstalled           bool          `json:"app_installed"`
			NetworkCompanyNo       int           `json:"network_company_no"`
			NetworkCompanyName     string        `json:"network_company_name"`
			AdminOverride          bool          `json:"adminOverride"`
			MapImage               string        `json:"map_image"`
			LifecycleStage         string        `json:"lifecycle_stage"`
			IpAddress              string        `json:"ip_address"`
			MemberDirectoryEnabled bool          `json:"memberDirectoryEnabled"`
			EnableHouseAccount     bool          `json:"enableHouseAccount"`
			KeepPrivate            bool          `json:"keep_private"`
			CustomerToken          string        `json:"customer_token"`
			MemberDiscount         int           `json:"member_discount"`
			Name                   string        `json:"name"`
			CcEmail                string        `json:"cc_email"`
			IsMember               bool          `json:"isMember"`
			WasMember              bool          `json:"wasMember"`
			HasApplied             bool          `json:"hasApplied"`
			MemberSince            string        `json:"member_since"`
			WhenApplied            string        `json:"when_applied"`
			UpdateByReference      bool          `json:"updateByReference"`
			Addresses              []interface{} `json:"addresses"`
			Email2                 string        `json:"email2"`
			Notes                  []interface{} `json:"notes"`
			Gender                 string        `json:"gender"`
			CompanyTitle           string        `json:"company_title"`
			AnniversaryDate        string        `json:"anniversary_date"`
			HasPassword            bool          `json:"hasPassword"`
			LastLogin              string        `json:"last_login"`
			LastPageView           string        `json:"lastPageView"`
			LastEmailOpened        string        `json:"lastEmailOpened"`
			LastNewsletter         string        `json:"lastNewsletter"`
			LastTouchpoint         string        `json:"lastTouchpoint"`
			LastCommunication      string        `json:"lastCommunication"`
			TotalTouchpoints       int           `json:"totalTouchpoints"`
			User                   struct {
				BusinessLocationNo int    `json:"business_location_no"`
				EmployeeId         string `json:"employee_id"`
				UserNo             int    `json:"user_no"`
				FirstName          string `json:"first_name"`
				LastName           string `json:"last_name"`
				MobileNumber       string `json:"mobile_number"`
				Email              string `json:"email"`
				Username           string `json:"username"`
				Status             string `json:"status"`
			} `json:"user"`
			UserModified struct {
				BusinessLocationNo int    `json:"business_location_no"`
				EmployeeId         string `json:"employee_id"`
				UserNo             int    `json:"user_no"`
				FirstName          string `json:"first_name"`
				LastName           string `json:"last_name"`
				MobileNumber       string `json:"mobile_number"`
				Email              string `json:"email"`
				Username           string `json:"username"`
				Status             string `json:"status"`
			} `json:"user_modified"`
			UserLastViewed struct {
				BusinessLocationNo int    `json:"business_location_no"`
				EmployeeId         string `json:"employee_id"`
				UserNo             int    `json:"user_no"`
				FirstName          string `json:"first_name"`
				LastName           string `json:"last_name"`
				MobileNumber       string `json:"mobile_number"`
				Email              string `json:"email"`
				Username           string `json:"username"`
				Status             string `json:"status"`
			} `json:"user_last_viewed"`
			LastViewed          string        `json:"last_viewed"`
			LastViewedBy        int           `json:"last_viewed_by"`
			ModifiedBy          int           `json:"modified_by"`
			AffiliateName       interface{}   `json:"affiliate_name"`
			AffiliateNo         int           `json:"affiliate_no"`
			SessionId           string        `json:"session_id"`
			PasswordExpires     string        `json:"password_expires"`
			Flag                string        `json:"flag"`
			VipFlag             bool          `json:"vip_flag"`
			LastContacted       string        `json:"last_contacted"`
			KloutScore          int           `json:"klout_score"`
			LinkedinHandle      string        `json:"linkedin_handle"`
			LinkedinId          string        `json:"linkedin_id"`
			LifecycleStageNo    int           `json:"lifecycle_stage_no"`
			CustomerSource      string        `json:"customer_source"`
			Timezone            string        `json:"timezone"`
			Username            string        `json:"username"`
			Settings            string        `json:"settings"`
			GeneratePassword    bool          `json:"generatePassword"`
			ProfilePhoto        string        `json:"profile_photo"`
			FoursquareHandle    string        `json:"foursquare_handle"`
			FoursquareId        string        `json:"foursquare_id"`
			CustomerStatus      string        `json:"customer_status"`
			Attributes          []interface{} `json:"attributes"`
			FacebookHandle      string        `json:"facebook_handle"`
			TwitterHandle       string        `json:"twitter_handle"`
			TiktokHandle        string        `json:"tiktok_handle"`
			InstagramHandle     string        `json:"instagram_handle"`
			Latitude            string        `json:"latitude"`
			Longitude           string        `json:"longitude"`
			CustomerNo          int           `json:"customer_no"`
			FirstName           string        `json:"first_name"`
			LastName            string        `json:"last_name"`
			Address             string        `json:"address"`
			Address2            string        `json:"address2"`
			City                string        `json:"city"`
			State               string        `json:"state"`
			ZipCode             string        `json:"zip_code"`
			Country             string        `json:"country"`
			PhoneNumber         string        `json:"phone_number"`
			MobileNumber        string        `json:"mobile_number"`
			Email               string        `json:"email"`
			Password            string        `json:"password"`
			CreatedOn           string        `json:"created_on"`
			DeviceType          string        `json:"device_type"`
			PassbookEnabled     string        `json:"passbook_enabled"`
			WalletEnabled       string        `json:"wallet_enabled"`
			OptInEmail          string        `json:"opt_in_email"`
			OptInSms            string        `json:"opt_in_sms"`
			OptInUpdates        string        `json:"opt_in_updates"`
			CompanyNo           int           `json:"company_no"`
			UserNo              int           `json:"user_no"`
			CompanyName         string        `json:"company_name"`
			Birthdate           string        `json:"birthdate"`
			ModifiedOn          string        `json:"modified_on"`
			Address3            string        `json:"address3"`
			CustomerReference   string        `json:"customer_reference"`
			FacebookId          string        `json:"facebook_id"`
			TwitterId           string        `json:"twitter_id"`
			TiktokId            string        `json:"tiktok_id"`
			InstagramId         string        `json:"instagram_id"`
			MiddleName          string        `json:"middle_name"`
			Website             string        `json:"website"`
			CustomerType        string        `json:"customer_type"`
			Schema              interface{}   `json:"schema"`
			TotalFavorite       int           `json:"totalFavorite"`
			SkipTriggers        bool          `json:"skipTriggers"`
			IncludeObjects      bool          `json:"includeObjects"`
			IncludeStats        bool          `json:"includeStats"`
			IncludeTotalRecords bool          `json:"includeTotalRecords"`
			TotalRecords        int           `json:"totalRecords"`
			RowsPerPage         int           `json:"rows_per_page"`
			PageNumber          int           `json:"page_number"`
			UpdateFields        string        `json:"updateFields"`
			IncludeCompany      bool          `json:"includeCompany"`
			Lightweight         bool          `json:"lightweight"`
			ItemDetailsFields   interface{}   `json:"itemDetailsFields"`
			IncludeItemDetails  bool          `json:"includeItemDetails"`
			SearchValue         interface{}   `json:"search_value"`
			SortBy              interface{}   `json:"sort_by"`
			ReturnTotal         int           `json:"return_total"`
			FromDate            string        `json:"from_date"`
			ToDate              string        `json:"to_date"`
			IncludeDetails      bool          `json:"includeDetails"`
			TimezoneOffset      int           `json:"timezone_offset"`
			TimezoneId          interface{}   `json:"timezone_id"`
		} `json:"returnObject"`
		ResponseTime    float64     `json:"responseTime"`
		TotalRecords    int         `json:"totalRecords"`
		MinPrice        int         `json:"minPrice"`
		MaxPrice        int         `json:"maxPrice"`
		PermissionLevel interface{} `json:"permission_level"`
		TimezoneOffset  int         `json:"timezone_offset"`
		TimezoneId      interface{} `json:"timezone_id"`
	} `json:"d"`
}

type TPVRetError struct {
}

// *********************************************************************
// *** Rest response header err and data
// *********************************************************************

type TRetError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// *********************************************************************
// *** Rest response data
// *********************************************************************

type TRetData struct {
	Status TRetError   `json:"status"`
	Data   interface{} `json:"data"`
}

// *********************************************************************
// *** Rest response simple data
// *********************************************************************

type TRetSimpleData struct {
	Status        TRetError `json:"status"`
	Authenticated bool      `json:"authenticated"`
}

// *********************************************************************
// *** Rest response err and data
// *********************************************************************

const (
	RET_OK                = 0
	RET_ERROR             = 8
	RET_PARAM_PARSE_ERROR = 16
	RET_CALL_ERROR        = 32
)
