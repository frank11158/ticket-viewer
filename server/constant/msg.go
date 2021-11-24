package constant

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	INVALID_PARAMS: "Invalid params error",
	ERROR:          "Fail",
	ERROR_AUTH:     "Auth fail",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
