package gocachebench

var (
	// Test keys data

	testDataKey1 = key1{
		userid:      102030,
		requestid:   "req-101",
		timestamp:   1670051110000,
		method:      "GET",
		contentType: "application/json",
		userAgent:   "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36",
		path:        "/api/items",
		queryParam:  map[string]string{"filter": "new"},
		statusCode:  200,
		region:      "us-west-1",
		apiVersion:  "v1.2",
	}

	testDataKey2 = key2{
		userid:       202040,
		requestid:    "req-202",
		method:       "POST",
		ip:           "203.0.113.0",
		header:       map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}, "Accept": {"application/json"}},
		isAuthorized: true,
		locale:       "en-US",
		timeZone:     "America/Los_Angeles",
	}

	testDataKey2Alt = key2{
		userid:       404040,
		requestid:    "req-303",
		method:       "PUT",
		ip:           "198.51.100.0",
		header:       map[string][]string{"Accept-Language": {"en-GB,en;q=0.5"}, "User-Agent": {"Mozilla/5.0"}},
		isAuthorized: false,
		locale:       "de-DE",
		timeZone:     "Europe/Berlin",
	}

	testDataKey3 = key3{
		userid:      303050,
		timestamp:   1670101110000,
		userAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		ip:          "192.0.2.0",
		queryParam:  map[string]string{"sort": "asc", "page": "1"},
		region:      "eu-central-1",
		deviceType:  "desktop",
		networkType: "ethernet",
	}

	testDataKey3Alt = key3{
		userid:      606060,
		timestamp:   1670202220000,
		userAgent:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
		ip:          "198.51.100.100",
		queryParam:  map[string]string{"category": "books", "author": "John Doe"},
		region:      "asia-southeast-1",
		deviceType:  "mobile",
		networkType: "5G",
	}

	testDataKey4 = key4{
		requestid:   "req-404",
		timestamp:   1670151110000,
		contentType: "text/xml",
		userAgent:   "curl/7.68.0",
		path:        "/api/update",
		header:      map[string][]string{"Authorization": {"Bearer abc123"}},
		statusCode:  204,
		apiVersion:  "v2.0",
		requestSize: 1024,
	}

	testDataKey5 = key5{
		userid:       505060,
		method:       "DELETE",
		contentType:  "text/plain",
		path:         "/api/user/delete",
		isAuthorized: false,
		deviceType:   "tablet",
		locale:       "fr-FR",
		networkType:  "wifi",
	}

	testDataKey5Alt = key5{
		userid:       707070,
		method:       "PATCH",
		contentType:  "application/xml",
		path:         "/api/settings/update",
		isAuthorized: true,
		deviceType:   "smartwatch",
		locale:       "es-ES",
		networkType:  "bluetooth",
	}

	// Test values data

	testDataValue1 = value1{
		data:         []byte("Sample data content"),
		headers:      map[string][]string{"Content-Type": {"application/json"}, "Cache-Control": {"max-age=3600"}},
		statusCode:   200,
		contentTypes: []string{"application/json", "application/xml"},
		payload: struct {
			items []struct {
				id        int64
				name      string
				price     float64
				available bool
			}
		}{
			items: []struct {
				id        int64
				name      string
				price     float64
				available bool
			}{
				{id: 1, name: "Item 1", price: 19.99, available: true},
				{id: 2, name: "Item 2", price: 29.99, available: false},
			},
		},
		queryParams:  map[string][]string{"page": {"1"}, "limit": {"10"}},
		responseSize: 2048,
	}

	testDataValue2 = value2{
		headers: map[string][]string{"Server": {"nginx/1.18.0"}, "Content-Length": {"1234"}},
		cookies: []struct {
			name    string
			value   string
			expires int64
			secure  bool
		}{
			{name: "session_cookie", value: "abc123", expires: 1680000000000, secure: true},
		},
		redirectURL: "https://example.com/redirect",
		timestamps:  []int64{1670001110000, 1670001120000},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
		}{
			{activityid: 101, timestamp: 1670001110000, kind: "login"},
			{activityid: 102, timestamp: 1670001120000, kind: "purchase"},
		},
	}

	testDataValue2Alt = value2{
		headers: map[string][]string{
			"Server":          {"Apache/2.4.41 (Ubuntu)"},
			"X-Powered-By":    {"PHP/7.4.3"},
			"X-Frame-Options": {"SAMEORIGIN"},
		},
		cookies: []struct {
			name    string
			value   string
			expires int64
			secure  bool
		}{
			{"loginToken", "login123456", 1695000000000, true},
			{"preferences", "darkmode=true&fontsize=medium", 1698000000000, false},
		},
		redirectURL: "https://example.org/welcome",
		timestamps:  []int64{1671001110000, 1671001120000, 1671001130000},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
		}{
			{activityid: 2001, timestamp: 1671001110000, kind: "LoginAttempt"},
			{activityid: 2002, timestamp: 1671001120000, kind: "SettingsChange"},
			{activityid: 2003, timestamp: 1671001130000, kind: "Logout"},
			{activityid: 2004, timestamp: 1671001140000, kind: "PasswordChange"},
			{activityid: 2005, timestamp: 1671001150000, kind: "ProfileUpdate"},
		},
	}

	testDataValue3 = value3{
		data: []byte("Enhanced response data for value3"),
		headers: map[string][]string{
			"Content-Encoding": {"br"},
			"Vary":             {"Accept-Encoding"},
			"X-Request-ID":     {"12345-xyz"},
		},
		cookies: []struct {
			name     string
			value    string
			path     string
			domain   string
			expires  int64
			secure   bool
			httpOnly bool
		}{
			{"trackingId", "track456789", "/tracking", "example.org", 1685000000000, false, false},
			{"preferences", "lang=en&currency=USD", "/settings", "example.org", 1697000000000, true, true},
		},
		statusCode:   403,
		contentTypes: []string{"application/xml", "application/json"},
		timestamps:   []int64{1673001110000, 1673001120000, 1673001130000},
		errorCode:    "Forbidden",
		payload: struct {
			items []struct {
				id        int64
				name      string
				category  string
				quantity  int
				price     float64
				available bool
				details   struct {
					description string
					imageURL    string
					rating      float32
					reviews     []struct {
						userID    int64
						comment   string
						rating    int
						timestamp int64
					}
				}
			}
			totalPrice float64
			user       struct {
				id        int64
				username  string
				email     string
				phone     string
				addresses []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}
				preferences map[string]string
			}
		}{
			items: []struct {
				id        int64
				name      string
				category  string
				quantity  int
				price     float64
				available bool
				details   struct {
					description string
					imageURL    string
					rating      float32
					reviews     []struct {
						userID    int64
						comment   string
						rating    int
						timestamp int64
					}
				}
			}{
				{
					id: 201, name: "4K Monitor", category: "Electronics", quantity: 5, price: 350.99, available: true,
					details: struct {
						description string
						imageURL    string
						rating      float32
						reviews     []struct {
							userID    int64
							comment   string
							rating    int
							timestamp int64
						}
					}{
						description: "High resolution monitor, ideal for gaming and professional use.",
						imageURL:    "https://example.org/images/4k-monitor.jpg",
						rating:      4.8,
						reviews: []struct {
							userID    int64
							comment   string
							rating    int
							timestamp int64
						}{
							{userID: 3001, comment: "Excellent image quality!", rating: 5, timestamp: 1673001115000},
						},
					},
				},
			},
			totalPrice: 1754.95,
			user: struct {
				id        int64
				username  string
				email     string
				phone     string
				addresses []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}
				preferences map[string]string
			}{
				id: 4001, username: "displayenthusiast", email: "displaylover@example.org", phone: "+12345678911",
				addresses: []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}{
					{line1: "789 Display Drive", line2: "", city: "Techtown", state: "TX", zipCode: "75001", country: "USA", isPrimary: true},
				},
				preferences: map[string]string{"displayType": "4K", "refreshRate": "144Hz"},
			},
		},
		isCached:    false,
		cacheTime:   1673002220000,
		queryParams: map[string][]string{"product": {"monitor"}, "resolution": {"4k"}},
		apiInfo: struct {
			Version string
			server  string
		}{
			Version: "2.5",
			server:  "api.techproducts.example.org",
		},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
			details    map[string]string
		}{
			{activityid: 6001, timestamp: 1673001110000, kind: "ViewProduct", details: map[string]string{"productId": "201"}},
			{activityid: 6002, timestamp: 1673001120000, kind: "AddToWishlist", details: map[string]string{"productId": "202"}},
		},
		responseSize:   6144,
		additionalData: map[string]interface{}{"saleEvent": "Summer Sale", "memberDiscount": "15%"},
	}

	testDataValue3Alt = value3{
		data: []byte("Extended data content for value3"),
		headers: map[string][]string{
			"Content-Type": {"text/html"},
			"Accept":       {"image/webp,*/*"},
			"Connection":   {"keep-alive"},
		},
		cookies: []struct {
			name     string
			value    string
			path     string
			domain   string
			expires  int64
			secure   bool
			httpOnly bool
		}{
			{"userSession", "xyz789", "/account", "secure.example.com", 1695000000000, true, true},
			{"cart", "items=5", "/shop", "secure.example.com", 1696000000000, false, false},
		},
		statusCode:   202,
		contentTypes: []string{"text/html", "application/xml"},
		timestamps:   []int64{1672001110000, 1672001120000, 1672001130000},
		errorCode:    "Accepted",
		payload: struct {
			items []struct {
				id        int64
				name      string
				category  string
				quantity  int
				price     float64
				available bool
				details   struct {
					description string
					imageURL    string
					rating      float32
					reviews     []struct {
						userID    int64
						comment   string
						rating    int
						timestamp int64
					}
				}
			}
			totalPrice float64
			user       struct {
				id        int64
				username  string
				email     string
				phone     string
				addresses []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}
				preferences map[string]string
			}
		}{
			items: []struct {
				id        int64
				name      string
				category  string
				quantity  int
				price     float64
				available bool
				details   struct {
					description string
					imageURL    string
					rating      float32
					reviews     []struct {
						userID    int64
						comment   string
						rating    int
						timestamp int64
					}
				}
			}{
				{
					id: 101, name: "Smartphone", category: "Electronics", quantity: 20, price: 299.99, available: true,
					details: struct {
						description string
						imageURL    string
						rating      float32
						reviews     []struct {
							userID    int64
							comment   string
							rating    int
							timestamp int64
						}
					}{
						description: "Latest model with advanced features",
						imageURL:    "https://example.com/images/smartphone.jpg",
						rating:      4.7,
						reviews: []struct {
							userID    int64
							comment   string
							rating    int
							timestamp int64
						}{
							{userID: 2001, comment: "Loving the new features!", rating: 5, timestamp: 1672001115000},
						},
					},
				},
			},
			totalPrice: 5999.80,
			user: struct {
				id        int64
				username  string
				email     string
				phone     string
				addresses []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}
				preferences map[string]string
			}{
				id: 3001, username: "techlover", email: "techlover@example.com", phone: "+12345678910",
				addresses: []struct {
					line1     string
					line2     string
					city      string
					state     string
					zipCode   string
					country   string
					isPrimary bool
				}{
					{line1: "456 Tech Ave", line2: "Suite 100", city: "Innovate City", state: "CA", zipCode: "94043", country: "USA", isPrimary: true},
				},
				preferences: map[string]string{"notifications": "enabled", "theme": "dark"},
			},
		},
		isCached:    true,
		cacheTime:   1672002220000,
		queryParams: map[string][]string{"sort": {"price"}, "filter": {"new"}},
		apiInfo: struct {
			Version string
			server  string
		}{
			Version: "3.1",
			server:  "api.newtech.example.com",
		},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
			details    map[string]string
		}{
			{activityid: 5001, timestamp: 1672001110000, kind: "Browse", details: map[string]string{"section": "Electronics"}},
			{activityid: 5002, timestamp: 1672001120000, kind: "Compare", details: map[string]string{"products": "101, 102"}},
		},
		responseSize:   5120,
		additionalData: map[string]interface{}{"promoCode": "NEWYEAR2023", "memberTier": "Gold"},
	}

	testDataValue4 = value4{
		headers:      map[string][]string{"X-Frame-Options": {"DENY"}, "X-Content-Type-Options": {"nosniff"}},
		contentTypes: []string{"text/html", "image/jpeg"},
		payload: struct {
			user struct {
				id        int64
				email     string
				phone     string
				addresses []struct {
					line1     string
					city      string
					country   string
					isPrimary bool
				}
			}
		}{
			user: struct {
				id        int64
				email     string
				phone     string
				addresses []struct {
					line1     string
					city      string
					country   string
					isPrimary bool
				}
			}{
				id:    12345,
				email: "user@example.com",
				phone: "+1234567890",
				addresses: []struct {
					line1     string
					city      string
					country   string
					isPrimary bool
				}{
					{line1: "123 Example St", city: "Anytown", country: "USA", isPrimary: true},
				},
			},
		},
		queryParams: map[string][]string{"filter": {"active"}},
	}

	testDataValue5 = value5{
		data: []byte("Response data for value5"),
		headers: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"text/html, application/xhtml+xml"},
		},
		cookies: []struct {
			name     string
			value    string
			path     string
			domain   string
			expires  int64
			secure   bool
			httpOnly bool
		}{
			{"sessionToken", "token123", "/", "example.com", 1690000000000, true, true},
		},
		redirectURL:  "https://example.com/redirect",
		statusCode:   200,
		contentTypes: []string{"application/json", "text/html"},
		timestamps:   []int64{1670001110000, 1670001120000},
		payload: struct {
			orders []struct {
				orderID int64
				items   []struct {
					id       int64
					name     string
					quantity int
					price    float64
				}
				orderDate   int64
				totalAmount float64
				status      string
			}
			userProfile struct {
				id       int64
				username string
				email    string
				phone    string
				history  []struct {
					timestamp int64
					action    string
					details   map[string]string
				}
			}
		}{
			orders: []struct {
				orderID int64
				items   []struct {
					id       int64
					name     string
					quantity int
					price    float64
				}
				orderDate   int64
				totalAmount float64
				status      string
			}{
				{
					orderID: 12345,
					items: []struct {
						id       int64
						name     string
						quantity int
						price    float64
					}{
						{id: 1, name: "Widget", quantity: 3, price: 19.99},
						{id: 2, name: "Gadget", quantity: 2, price: 29.99},
					},
					orderDate:   1670001100000,
					totalAmount: 129.95,
					status:      "Completed",
				},
			},
			userProfile: struct {
				id       int64
				username string
				email    string
				phone    string
				history  []struct {
					timestamp int64
					action    string
					details   map[string]string
				}
			}{
				id:       1001,
				username: "user1001",
				email:    "user1001@example.com",
				phone:    "+123456789",
				history: []struct {
					timestamp int64
					action    string
					details   map[string]string
				}{
					{timestamp: 1670001110000, action: "Login", details: map[string]string{"ip": "192.168.1.1"}},
					{timestamp: 1670001120000, action: "Purchase", details: map[string]string{"orderID": "12345"}},
				},
			},
		},
		isCached:    true,
		cacheTime:   1670002220000,
		queryParams: map[string][]string{"category": {"electronics"}, "brand": {"brandX"}},
		apiInfo: struct {
			Version string
			server  string
		}{
			Version: "1.0",
			server:  "api.example.com",
		},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
			details    map[string]string
		}{
			{activityid: 3001, timestamp: 1670001110000, kind: "View", details: map[string]string{"page": "Homepage"}},
			{activityid: 3002, timestamp: 1670001120000, kind: "Update", details: map[string]string{"section": "Profile"}},
		},
		responseSize:   4096,
		additionalInfo: map[string]interface{}{"note": "Sample additional info"},
	}

	testDataValue5Alt = value5{
		data: []byte("Alternative sample data for value5"),
		headers: map[string][]string{
			"Content-Type": {"text/xml"},
			"Accept":       {"application/json"},
		},
		cookies: []struct {
			name     string
			value    string
			path     string
			domain   string
			expires  int64
			secure   bool
			httpOnly bool
		}{
			{"authToken", "xyz987", "/", "another-example.com", 1700000000000, false, true},
		},
		redirectURL:  "https://another-example.com/redirect",
		statusCode:   201,
		contentTypes: []string{"text/xml", "application/json"},
		timestamps:   []int64{1670002210000, 1670002220000},
		payload: struct {
			orders []struct {
				orderID int64
				items   []struct {
					id       int64
					name     string
					quantity int
					price    float64
				}
				orderDate   int64
				totalAmount float64
				status      string
			}
			userProfile struct {
				id       int64
				username string
				email    string
				phone    string
				history  []struct {
					timestamp int64
					action    string
					details   map[string]string
				}
			}
		}{
			orders: []struct {
				orderID int64
				items   []struct {
					id       int64
					name     string
					quantity int
					price    float64
				}
				orderDate   int64
				totalAmount float64
				status      string
			}{
				{
					orderID: 54321,
					items: []struct {
						id       int64
						name     string
						quantity int
						price    float64
					}{
						{id: 3, name: "Book", quantity: 1, price: 15.00},
						{id: 4, name: "Pen", quantity: 10, price: 1.50},
					},
					orderDate:   1670002200000,
					totalAmount: 30.00,
					status:      "Pending",
				},
			},
			userProfile: struct {
				id       int64
				username string
				email    string
				phone    string
				history  []struct {
					timestamp int64
					action    string
					details   map[string]string
				}
			}{
				id:       2002,
				username: "user2002",
				email:    "user2002@another-example.com",
				phone:    "+1234567891",
				history: []struct {
					timestamp int64
					action    string
					details   map[string]string
				}{
					{timestamp: 1670002210000, action: "Signup", details: map[string]string{"referral": "user1001"}},
					{timestamp: 1670002220000, action: "AddToCart", details: map[string]string{"itemID": "3", "quantity": "1"}},
				},
			},
		},
		isCached:    false,
		cacheTime:   1670003330000,
		queryParams: map[string][]string{"category": {"stationery"}, "inStock": {"true"}},
		apiInfo: struct {
			Version string
			server  string
		}{
			Version: "2.1",
			server:  "api.another-example.com",
		},
		userActivity: []struct {
			activityid int64
			timestamp  int64
			kind       string
			details    map[string]string
		}{
			{activityid: 4001, timestamp: 1670002210000, kind: "Search", details: map[string]string{"query": "books"}},
			{activityid: 4002, timestamp: 1670002220000, kind: "ViewProduct", details: map[string]string{"productID": "3"}},
		},
		responseSize:   5120,
		additionalInfo: map[string]interface{}{"note": "Alternative response"},
	}
)
