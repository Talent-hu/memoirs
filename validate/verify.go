package validate

var (
	LoginVerify          = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Identity": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	GetUserListVerify    = Rules{"PageNum": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
