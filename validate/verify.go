package validate

// auth module
var (
	LoginVerify          = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Identity": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	QueryPageListVerify  = Rules{"PageNum": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AddRoleVerify        = Rules{"RoleCode": {NotEmpty()}, "RoleName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	RoleAndMenuVerify    = Rules{"RoleId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	UpdateRoleVerify     = Rules{"RoleId": {NotEmpty()}, "RoleCode": {NotEmpty()}, "RoleName": {NotEmpty()}, "ParentId": {NotEmpty(), Ge("0")}}
	DeleteRoleVerify     = Rules{"RoleId": {NotEmpty(), Ge("0")}}
	AddMenuVerify        = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Sort": {NotEmpty()}, "Title": {NotEmpty()}, "Icon": {NotEmpty()}, "ParentId": {NotEmpty()}, "FontType": {NotEmpty()}, "FontSize": {NotEmpty()}, "HasBtn": {NotEmpty()}}
	DeleteMenuVerify     = Rules{"MenuIds": {NotEmpty()}}
)

// bank module

var (
	AddSubjectVerify    = Rules{"Code": {NotEmpty()}, "Name": {NotEmpty()}}
	UpdateSubjectVerify = Rules{"ID": {NotEmpty()}, "Code": {NotEmpty()}, "Name": {NotEmpty()}}
)
