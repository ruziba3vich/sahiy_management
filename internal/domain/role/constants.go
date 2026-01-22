package role

// Role represents user role with name and value
type Role struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Role constants
const (
	SuperAdminValue = 99
	ManagerValue    = 50
	EmployeeValue   = 10
)

const (
	SuperAdminName = "Super Admin"
	ManagerName    = "Manager"
	EmployeeName   = "Employee"
)

// GetAllRoles returns all available roles
func GetAllRoles() []Role {
	return []Role{
		{Name: SuperAdminName, Value: SuperAdminValue},
		{Name: ManagerName, Value: ManagerValue},
		{Name: EmployeeName, Value: EmployeeValue},
	}
}

// GetRoleByValue returns role name by value
func GetRoleByValue(value int) string {
	switch value {
	case SuperAdminValue:
		return SuperAdminName
	case ManagerValue:
		return ManagerName
	case EmployeeValue:
		return EmployeeName
	default:
		return "Unknown"
	}
}

// IsValidRole checks if role value is valid
func IsValidRole(value int) bool {
	return value == SuperAdminValue || value == ManagerValue || value == EmployeeValue
}
