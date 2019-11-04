package service_test

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"

	"github.com/milhamhidayat/golang-clean-code-v2/domain"
	"github.com/milhamhidayat/golang-clean-code-v2/domain/mocks"
	"github.com/milhamhidayat/golang-clean-code-v2/employee/service"
	"github.com/milhamhidayat/golang-clean-code-v2/testdata"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	var employee domain.Employee
	testdata.UnmarshallGoldenToJSON(t, "employee-1S9XpJCvJbt1plvU36tAcJWS2ZW", &employee)

	mockDepartmentRepo := new(mocks.DepartmentRepository)
	mockEmployeeRepo := new(mocks.EmployeeRepository)

	tests := map[string]struct {
		employeeRepo map[string]testdata.FuncCall
		expectedErr  error
	}{
		"success": {
			employeeRepo: map[string]testdata.FuncCall{
				"Create": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), &employee},
					Output: []interface{}{nil},
				},
			},
			expectedErr: nil,
		},
		"with error create an employee": {
			employeeRepo: map[string]testdata.FuncCall{
				"Create": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), &employee},
					Output: []interface{}{errors.New("unexpected error")},
				},
			},
			expectedErr: errors.New("unexpected error"),
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			for name, fn := range tc.employeeRepo {
				if fn.Called {
					mockEmployeeRepo.On(name, fn.Input...).Return(fn.Output...).Once()
				}
			}

			employeeService := service.New(mockDepartmentRepo, mockEmployeeRepo)
			err := employeeService.Create(context.Background(), &employee)

			mockEmployeeRepo.AssertExpectations(t)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	var (
		employee   domain.Employee
		department domain.Department
	)
	testdata.UnmarshallGoldenToJSON(t, "employee-1S9XpJCvJbt1plvU36tAcJWS2ZW", &employee)
	testdata.UnmarshallGoldenToJSON(t, "department-0ujsswThIGTUYm2K8FjOOfXtY1K", &department)

	mockDepartmentRepo := new(mocks.DepartmentRepository)
	mockEmployeeRepo := new(mocks.EmployeeRepository)

	newEmployee := employee
	newEmployee.LastName = "Diana"
	newEmployee.Department = department

	tests := map[string]struct {
		employeeRepo   map[string]testdata.FuncCall
		departmentRepo map[string]testdata.FuncCall
		expectedRes    domain.Employee
		expectedErr    error
	}{
		"success": {
			employeeRepo: map[string]testdata.FuncCall{
				"Update": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee},
					Output: []interface{}{newEmployee, nil},
				},
			},
			departmentRepo: map[string]testdata.FuncCall{
				"Get": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee.Department.ID},
					Output: []interface{}{department, nil},
				},
			},
			expectedRes: newEmployee,
			expectedErr: nil,
		},
		"with error update an employee": {
			employeeRepo: map[string]testdata.FuncCall{
				"Update": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee},
					Output: []interface{}{domain.Employee{}, errors.New("unexpected error")},
				},
			},
			departmentRepo: map[string]testdata.FuncCall{
				"Get": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee.Department.ID},
					Output: []interface{}{department, nil},
				},
			},
			expectedRes: domain.Employee{},
			expectedErr: errors.New("unexpected error"),
		},
		"with error get a department": {
			employeeRepo: map[string]testdata.FuncCall{
				"Update": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee},
					Output: []interface{}{newEmployee, nil},
				},
			},
			departmentRepo: map[string]testdata.FuncCall{
				"Get": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), newEmployee.Department.ID},
					Output: []interface{}{domain.Department{}, errors.New("unknown error")},
				},
			},
			expectedRes: domain.Employee{},
			expectedErr: errors.New("unknown error"),
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			for name, fn := range tc.employeeRepo {
				if fn.Called {
					mockEmployeeRepo.On(name, fn.Input...).Return(fn.Output...).Once()
				}
			}

			for name, fn := range tc.departmentRepo {
				if fn.Called {
					mockDepartmentRepo.On(name, fn.Input...).Return(fn.Output...).Once()
				}
			}

			employeeService := service.New(mockDepartmentRepo, mockEmployeeRepo)
			res, err := employeeService.Update(context.Background(), newEmployee)

			mockEmployeeRepo.AssertExpectations(t)
			mockDepartmentRepo.AssertExpectations(t)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, newEmployee, res)
		})
	}
}

func TestDelete(t *testing.T) {
	var employee domain.Employee
	testdata.UnmarshallGoldenToJSON(t, "employee-1S9XpJCvJbt1plvU36tAcJWS2ZW", &employee)

	mockDepartmentRepo := new(mocks.DepartmentRepository)
	mockEmployeeRepo := new(mocks.EmployeeRepository)

	tests := map[string]struct {
		employeeRepo map[string]testdata.FuncCall
		expectedErr  error
	}{
		"success": {
			employeeRepo: map[string]testdata.FuncCall{
				"Delete": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), employee.ID},
					Output: []interface{}{nil},
				},
			},
			expectedErr: nil,
		},
		"with error from employee repo": {
			employeeRepo: map[string]testdata.FuncCall{
				"Delete": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), employee.ID},
					Output: []interface{}{errors.New("unexpected error")},
				},
			},
			expectedErr: errors.New("unexpected error"),
		},
		"not found": {
			employeeRepo: map[string]testdata.FuncCall{
				"Delete": testdata.FuncCall{
					Called: true,
					Input:  []interface{}{context.Background(), employee.ID},
					Output: []interface{}{domain.ErrNotFound},
				},
			},
			expectedErr: domain.ErrNotFound,
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			for name, fn := range tc.employeeRepo {
				if fn.Called {
					mockEmployeeRepo.On(name, fn.Input...).Return(fn.Output...).Once()
				}
			}

			employeeService := service.New(mockDepartmentRepo, mockEmployeeRepo)
			err := employeeService.Delete(context.Background(), employee.ID)

			mockEmployeeRepo.AssertExpectations(t)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			require.NoError(t, err)
		})
	}
}
