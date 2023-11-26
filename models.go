package gocachebench

type key1 struct {
	userid      int64
	requestid   string
	timestamp   int64
	method      string
	contentType string
	userAgent   string
	path        string
	queryParam  map[string]string
	statusCode  int
	region      string
	apiVersion  string
}

type key2 struct {
	userid       int64
	requestid    string
	method       string
	ip           string
	header       map[string][]string
	isAuthorized bool
	locale       string
	timeZone     string
}

type key3 struct {
	userid      int64
	timestamp   int64
	userAgent   string
	ip          string
	queryParam  map[string]string
	region      string
	deviceType  string
	networkType string
}

type key4 struct {
	requestid   string
	timestamp   int64
	contentType string
	userAgent   string
	path        string
	header      map[string][]string
	statusCode  int
	apiVersion  string
	requestSize int64
}

type key5 struct {
	userid       int64
	method       string
	contentType  string
	path         string
	isAuthorized bool
	deviceType   string
	locale       string
	networkType  string
}

type value1 struct {
	data         []byte
	headers      map[string][]string
	statusCode   int
	contentTypes []string
	payload      struct {
		items []struct {
			id        int64
			name      string
			price     float64
			available bool
		}
	}
	queryParams  map[string][]string
	responseSize int64
}

type value2 struct {
	headers map[string][]string
	cookies []struct {
		name    string
		value   string
		expires int64
		secure  bool
	}
	redirectURL  string
	timestamps   []int64
	userActivity []struct {
		activityid int64
		timestamp  int64
		kind       string
	}
}

type value3 struct {
	data    []byte
	headers map[string][]string
	cookies []struct {
		name     string
		value    string
		path     string
		domain   string
		expires  int64
		secure   bool
		httpOnly bool
	}
	statusCode   int
	contentTypes []string
	timestamps   []int64
	errorCode    string
	payload      struct {
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
	}
	isCached    bool
	cacheTime   int64
	queryParams map[string][]string
	apiInfo     struct {
		Version string
		server  string
	}
	userActivity []struct {
		activityid int64
		timestamp  int64
		kind       string
		details    map[string]string
	}
	responseSize   int64
	additionalData map[string]interface{}
}

type value4 struct {
	headers      map[string][]string
	contentTypes []string
	payload      struct {
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
	}
	queryParams map[string][]string
}

type value5 struct {
	data    []byte
	headers map[string][]string
	cookies []struct {
		name     string
		value    string
		path     string
		domain   string
		expires  int64
		secure   bool
		httpOnly bool
	}
	redirectURL  string
	statusCode   int
	contentTypes []string
	timestamps   []int64
	payload      struct {
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
	}
	isCached    bool
	cacheTime   int64
	queryParams map[string][]string
	apiInfo     struct {
		Version string
		server  string
	}
	userActivity []struct {
		activityid int64
		timestamp  int64
		kind       string
		details    map[string]string
	}
	responseSize   int64
	additionalInfo map[string]interface{}
}
