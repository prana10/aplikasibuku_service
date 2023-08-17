package infra

type Meta struct {
	Message string
	Status  string
	Code    int
}

type Response struct {
	Meta Meta
	Data interface{}
}

type ResponseList struct {
	Meta  Meta
	Data  interface{}
	Total int
}

// response API Format
// meta : message, status, code
// data : data
func NewResponseAPI(message string, status string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

// response API List Format
// meta : message, status, code
// data : data
// total : total data
func NewResponseAPIList(message string, status string, code int, data interface{}, total int) ResponseList {
	meta := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}

	response := ResponseList{
		Meta:  meta,
		Data:  data,
		Total: total,
	}

	return response
}
