package incomsms

type StatusCode string

// each sms type has different phone number and cost
type FeeType int
type MsgContentType int

const (
	CodeSuccess               StatusCode = "1"
	CodeNonRegisteredTemplate StatusCode = "537"
	// The message template has to have [QC] or (QC) prefix
	CodeTemplateRequireQCPrefix StatusCode = "536"
	// The phone number no longer has 11 digits but 10 digits
	CodePhoneNumberNoLonger11Digits   StatusCode = "535"
	CodeKeywordBlacklist              StatusCode = "530"
	CodeMessageLengthTooLong          StatusCode = "515"
	CodeInvalidUsernameOrPasswordOrIP StatusCode = "267"
	CodeSendMTNotAllowed              StatusCode = "510"
	CodeSessionPrefixNotDeclared      StatusCode = "511"
	CodeSendDuplicatedMessage         StatusCode = "304"
	CodeCreateInvalidConcentrator     StatusCode = "253"
	CodeEmptyServiceCode              StatusCode = "356"
	CodeServiceNotFoundOrActivated    StatusCode = "357"
	CodePhoneNumberBlocked            StatusCode = "360"
	CodeSessionNotFoundOrActivated    StatusCode = "359"
	CodeInvalidPhoneNumber            StatusCode = "392"
	CodeInvalidUsernameOrPassword     StatusCode = "393"
	CodeUserNotFound                  StatusCode = "394"
	CodeNonRegisteredIP               StatusCode = "395"
	CodeServiceSessionNotFound        StatusCode = "396"
	CodeProviderNotFound              StatusCode = "397"
	CodePartnerNotFound               StatusCode = "398"
	CodeDuplicatedPartnerMT           StatusCode = "399"
	CodeNonRegisteredBrandName        StatusCode = "509"

	NoFee  FeeType = 0
	HasFee FeeType = 1

	MsgContentAscii   MsgContentType = 0
	MsgContentUnicode MsgContentType = 12
)

type Credential struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
