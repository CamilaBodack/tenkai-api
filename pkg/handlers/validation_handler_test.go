package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	mockRepo "github.com/softplan/tenkai-api/pkg/dbms/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidateVariables_NotEmpty(t *testing.T) {
	appContext := AppContext{}
	req, err := http.NewRequest("POST", "/validateVariables",
		createPayloadWithScopeAndID(999, "global"))
	assert.NoError(t, err)

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironmentAndScope", 999, "global").Return(variables, nil)

	var varRules []model.VariableRule
	varRules = append(varRules, getVarRule("dbUsername", "NotEmpty", ""))
	mockVariableRuleDAO := &mockRepo.VariableRuleDAOInterface{}
	mockVariableRuleDAO.On("ListVariableRules").Return(varRules, nil)

	appContext.Repositories.VariableDAO = mockVariableDAO
	appContext.Repositories.VariableRuleDAO = mockVariableRuleDAO

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.validateVariables)
	handler.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironmentAndScope", 1)
	mockVariableRuleDAO.AssertNumberOfCalls(t, "ListVariableRules", 1)
	assert.Equal(t, http.StatusOK, rr.Code, "Response should be ok.")

	r := string(rr.Body.Bytes())
	assert.Equal(t, r, `{"InvalidVariables":[{"scope":"global","name":"dbUsername","value":"","variableRule":"dbUsername","ruleType":"NotEmpty","valueRule":""}]}`)
}

func TestValidateVariables_VariabeRule_UsingRegex_NotEmpty(t *testing.T) {
	appContext := AppContext{}
	req, err := http.NewRequest("POST", "/validateVariables",
		createPayloadWithScopeAndID(999, "global"))
	assert.NoError(t, err)

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironmentAndScope", 999, "global").Return(variables, nil)

	var varRules []model.VariableRule
	varRules = append(varRules, getVarRule("dbUser.+", "NotEmpty", "")) // Using RegEx here ====> dbUser.+
	mockVariableRuleDAO := &mockRepo.VariableRuleDAOInterface{}
	mockVariableRuleDAO.On("ListVariableRules").Return(varRules, nil)

	appContext.Repositories.VariableDAO = mockVariableDAO
	appContext.Repositories.VariableRuleDAO = mockVariableRuleDAO

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.validateVariables)
	handler.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironmentAndScope", 1)
	mockVariableRuleDAO.AssertNumberOfCalls(t, "ListVariableRules", 1)
	assert.Equal(t, http.StatusOK, rr.Code, "Response should be ok.")

	r := string(rr.Body.Bytes())
	assert.Equal(t, r, `{"InvalidVariables":[{"scope":"global","name":"dbUsername","value":"","variableRule":"dbUser.+","ruleType":"NotEmpty","valueRule":""}]}`)
}

func TestValidateVariables_VariabeRule_UsingRegex_NotMatch(t *testing.T) {
	appContext := AppContext{}
	req, err := http.NewRequest("POST", "/validateVariables",
		createPayloadWithScopeAndID(999, "global"))
	assert.NoError(t, err)

	invalidVar := getVar("dbPassword", "")
	validVar := getVar("dbPassword", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironmentAndScope", 999, "global").Return(variables, nil)

	var varRules []model.VariableRule
	varRules = append(varRules, getVarRule("dbUser.+", "NotEmpty", "")) // Using RegEx here ====> dbUser.+
	mockVariableRuleDAO := &mockRepo.VariableRuleDAOInterface{}
	mockVariableRuleDAO.On("ListVariableRules").Return(varRules, nil)

	appContext.Repositories.VariableDAO = mockVariableDAO
	appContext.Repositories.VariableRuleDAO = mockVariableRuleDAO

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.validateVariables)
	handler.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironmentAndScope", 1)
	mockVariableRuleDAO.AssertNumberOfCalls(t, "ListVariableRules", 1)
	assert.Equal(t, http.StatusOK, rr.Code, "Response should be ok.")

	r := string(rr.Body.Bytes())
	assert.Equal(t, r, `{"InvalidVariables":[]}`)
}

func TestValidateNotEmpty(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	var vars []model.Variable
	vars = append(vars, invalidVar, validVar)

	var vrs []model.VariableRule
	vrs = append(vrs, getVarRule("dbUsername", "NotEmpty", ""))

	ivr, err := appContext.validate(vars, vrs)
	assert.Nil(t, err)
	assert.NotNil(t, ivr)
	assert.NotEmpty(t, ivr.InvalidVariables)
	assert.Equal(t, 1, len(ivr.InvalidVariables))
}

func TestValidateStartsWith(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("urlapiFoo", "my-server.com/foo")
	validVar := getVar("urlapiFoo", "http://my-server.com/foo")

	var vars []model.Variable
	vars = append(vars, invalidVar, validVar)

	var vrs []model.VariableRule
	vrs = append(vrs, getVarRule("urlapiFoo", "StartsWith", "http://"))

	ivr, err := appContext.validate(vars, vrs)
	assert.Nil(t, err)
	assert.NotNil(t, ivr)
	assert.NotEmpty(t, ivr.InvalidVariables)
	assert.Equal(t, 1, len(ivr.InvalidVariables))
}

func TestValidateEndsWith(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("urlApolloServer", "http://apollo-server.com/api-docs/graphql")
	validVar := getVar("urlApolloServer", "http://apollo-server.com/api/graphql")

	var vars []model.Variable
	vars = append(vars, invalidVar, validVar)

	var vrs []model.VariableRule
	vrs = append(vrs, getVarRule("urlApolloServer", "EndsWith", "/api/graphql"))

	ivr, err := appContext.validate(vars, vrs)
	assert.Nil(t, err)
	assert.NotNil(t, ivr)
	assert.NotEmpty(t, ivr.InvalidVariables)
	assert.Equal(t, 1, len(ivr.InvalidVariables))
}

func TestValidateRegEx(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("authTypes", "Other")
	validVar := getVar("authTypes", "OpenID")

	var vars []model.Variable
	vars = append(vars, invalidVar, validVar)

	var vrs []model.VariableRule
	vrs = append(vrs, getVarRule("authTypes", "RegEx", "OpenID|Internals"))

	ivr, err := appContext.validate(vars, vrs)
	assert.Nil(t, err)
	assert.NotNil(t, ivr)
	assert.NotEmpty(t, ivr.InvalidVariables)
	assert.Equal(t, 1, len(ivr.InvalidVariables))
}

func TestValidateEnvironmentVariables(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironment", 999).Return(variables, nil)

	var varRules []model.VariableRule
	varRules = append(varRules, getVarRule("dbUsername", "NotEmpty", ""))
	mockVariableRuleDAO := &mockRepo.VariableRuleDAOInterface{}
	mockVariableRuleDAO.On("ListVariableRules").Return(varRules, nil)

	appContext.Repositories.VariableDAO = mockVariableDAO
	appContext.Repositories.VariableRuleDAO = mockVariableRuleDAO

	req, err := http.NewRequest("GET", "/validateEnvVars/999", bytes.NewBuffer(nil))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/validateEnvVars/{envId}", appContext.validateEnvironmentVariables).Methods("GET")
	r.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironment", 1)
	mockVariableRuleDAO.AssertNumberOfCalls(t, "ListVariableRules", 1)
	assert.Equal(t, http.StatusOK, rr.Code, "Response should be ok.")

	res := string(rr.Body.Bytes())
	assert.Equal(t, res, `{"InvalidVariables":[{"scope":"global","name":"dbUsername","value":"","variableRule":"dbUsername","ruleType":"NotEmpty","valueRule":""}]}`)
}

func TestValidateEnvironmentVariables_Error1(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironment", 999).Return(variables, errors.New("some error"))

	appContext.Repositories.VariableDAO = mockVariableDAO

	req, err := http.NewRequest("GET", "/validateEnvVars/999", bytes.NewBuffer(nil))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/validateEnvVars/{envId}", appContext.validateEnvironmentVariables).Methods("GET")
	r.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironment", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500.")
}

func TestValidateEnvironmentVariables_Error2(t *testing.T) {
	appContext := AppContext{}

	invalidVar := getVar("dbUsername", "")
	validVar := getVar("dbUsername", "user")

	mockVariableDAO := &mockRepo.VariableDAOInterface{}
	var variables []model.Variable
	variables = append(variables, invalidVar, validVar)
	mockVariableDAO.On("GetAllVariablesByEnvironment", 999).Return(variables, nil)

	var varRules []model.VariableRule
	varRules = append(varRules, getVarRule("dbUsername", "NotEmpty", ""))
	mockVariableRuleDAO := &mockRepo.VariableRuleDAOInterface{}
	mockVariableRuleDAO.On("ListVariableRules").Return(varRules, errors.New("some error"))

	appContext.Repositories.VariableDAO = mockVariableDAO
	appContext.Repositories.VariableRuleDAO = mockVariableRuleDAO

	req, err := http.NewRequest("GET", "/validateEnvVars/999", bytes.NewBuffer(nil))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	mockPrincipal(req)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/validateEnvVars/{envId}", appContext.validateEnvironmentVariables).Methods("GET")
	r.ServeHTTP(rr, req)

	mockVariableDAO.AssertNumberOfCalls(t, "GetAllVariablesByEnvironment", 1)
	mockVariableRuleDAO.AssertNumberOfCalls(t, "ListVariableRules", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500.")
}

func getVar(name string, value string) model.Variable {
	var v model.Variable
	v.ID = 999
	v.Scope = "global"
	v.Name = name
	v.Value = value
	v.Secret = false
	v.Description = "Mock variable."
	v.EnvironmentID = 999
	return v
}

func getValueRule(rType string, rValue string) model.ValueRule {
	var vlr model.ValueRule
	vlr.ID = 888
	vlr.Type = rType
	vlr.Value = rValue
	vlr.VariableRuleID = 999
	return vlr
}

func getVarRule(rName string, rType string, rValue string) model.VariableRule {
	var vrr model.VariableRule
	vrr.ID = 999
	vrr.Name = rName
	vlr := getValueRule(rType, rValue)
	vrr.ValueRules = append(vrr.ValueRules, &vlr)
	return vrr
}
