package service

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"sort"
	"strconv"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
	wraperrors "github.com/pkg/errors"
)

// GetToken adds or update user token info into database and return token.
func (s *Service) GetToken(ctx context.Context, user *models.User) (string, error) {
	existingUser, err := s.db.GetUserByName(ctx, user.Username)
	if err != nil {
		log().Debugf("User Not Found: %+v", err)
	}

	if existingUser != nil {
		user.ID = existingUser.ID
	}

	if _, err = s.db.AddUser(ctx, user); err != nil {
		return "", err
	}

	jwtToken, err := generateJWT(user)
	if err != nil {
		return "", wraperrors.Wrap(err, "failed to create jwt token")
	}

	return jwtToken, nil
}

// GetUsers retrieve list of users.
func (s *Service) GetUsers(ctx context.Context) ([]string, error) {
	usersList := make([]string, 0)

	users, err := s.db.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		usersList = append(usersList, user.Username)
	}

	return usersList, nil
}

// GetAllCases get all cases reported today.
func (s *Service) GetAllCases() (string, string, error) {
	totalCases := 0
	dataset, err := s.apiClient.GetDataSet()
	if err != nil {
		return "", "", err
	}

	headRow := dataset[0]

	for index, row := range dataset {
		if index == 0 {
			continue
		}
		caseCount, convertErr := strconv.Atoi(row[len(row)-1])
		if convertErr != nil {
			return "", "", wraperrors.Wrap(convertErr, "failed to covert datatype")
		}

		totalCases += caseCount
	}

	return fmt.Sprintf("%s", headRow[len(headRow)-1]), fmt.Sprintf("%d", totalCases), nil
}

// GetCountryCases get all cases reported in country today or from date.
func (s *Service) GetCountryCases(filter map[string]interface{}) ([]interface{}, error) {
	results := make([]interface{}, 0)
	dataset, err := s.apiClient.GetDataSet()
	if err != nil {
		return nil, err
	}

	fmt.Println(filter)
	filterType, ok := filter["type"].(string)
	if !ok {
		fmt.Println(filterType)
		return nil, wraperrors.New("failed to get filter type key")
	}

	switch filterType {
	case "country":
		countryName, filterOk := filter["country_name"].(*string)
		if !filterOk {
			return nil, wraperrors.New("failed to get filter country key")
		}

		return getCountryCases(swag.StringValue(countryName), dataset), nil
	case "date":
		countryName, filterOk := filter["country_name"].(*string)
		if !filterOk {
			return nil, wraperrors.New("failed to get filter country key")
		}

		dateFrom, filterOk := filter["date_from"].(*string)
		if !filterOk {
			return nil, wraperrors.New("failed to get filter date key")
		}

		results, err = getCountryCasesFromDate(swag.StringValue(countryName), swag.StringValue(dateFrom), dataset)
		if err != nil {
			return nil, err
		}
	default:
		results, err = getAllCasesCountryWise(dataset)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// GetTopCasesCountries get all top N cases reported today.
func (s *Service) GetTopCasesCountries(count int64) ([]interface{}, error) {
	results := make([]interface{}, 0)

	dataset, err := s.apiClient.GetDataSet()
	if err != nil {
		return nil, err
	}

	sortedResults, err := getAllCasesCountryWise(dataset)
	if err != nil {
		return nil, err
	}

	for index, countryData := range sortedResults {
		if count == int64(index) {
			break
		}

		results = append(results, countryData)
	}

	return results, nil
}

func getAllCasesCountryWise(dataset [][]string) ([]interface{}, error) {
	structArray := make([]countryCases, 0)
	sortedData := make([]interface{}, 0)

	for index, row := range dataset {
		if index == 0 {
			continue
		}

		caseCount, convertErr := strconv.Atoi(row[len(row)-1])
		if convertErr != nil {
			return nil, wraperrors.Wrap(convertErr, "failed to covert datatype")
		}

		structArray = append(structArray, countryCases{
			Name:       row[1],
			CasesCount: caseCount,
		})
	}

	sort.SliceStable(structArray, func(i, j int) bool {
		return structArray[i].CasesCount > structArray[j].CasesCount
	})

	for _, structVal := range structArray {
		sortedData = append(sortedData,
			map[string]interface{}{
				structVal.Name: structVal.CasesCount,
			})
	}

	return sortedData, nil
}

func getCountryCases(countryName string, dataset [][]string) []interface{} {
	countryData := make([]interface{}, 0)

	for _, row := range dataset {
		if row[1] != countryName {
			continue
		}

		countryData = append(countryData,
			map[string]interface{}{
				row[1]: row[len(row)-1],
			})

	}

	return countryData
}

func getCountryCasesFromDate(countryName, dateFrom string, dataset [][]string) ([]interface{}, error) {
	countryData := make([]interface{}, 0)
	columnIndex := 0
	rowIndex := 0
	totalCases := 0

	for i, row := range dataset {
		if row[1] == countryName {
			rowIndex = i
		}
	}

	for i, val := range dataset[0] {
		if val == dateFrom {
			columnIndex = i
		}
	}

	for i := columnIndex; i < len(dataset[rowIndex]); i++ {
		caseCount, convertErr := strconv.Atoi(dataset[rowIndex][i])
		if convertErr != nil {
			return nil, wraperrors.Wrap(convertErr, "failed to covert datatype")
		}

		totalCases += caseCount
	}

	countryData = append(countryData,
		map[string]interface{}{
			countryName: totalCases,
			"DateFrom":  dateFrom,
		})

	return countryData, nil
}
