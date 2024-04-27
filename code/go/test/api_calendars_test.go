/*
spoke_handler

Testing CalendarsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_CalendarsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalendarsAPIService CalendarCalendarsCreateCalendar", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CalendarsAPI.CalendarCalendarsCreateCalendar(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalendarsAPIService CalendarCalendarsDeleteCalendar", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.CalendarsAPI.CalendarCalendarsDeleteCalendar(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalendarsAPIService CalendarCalendarsGetCalendar", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.CalendarsAPI.CalendarCalendarsGetCalendar(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalendarsAPIService CalendarCalendarsListCalendars", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CalendarsAPI.CalendarCalendarsListCalendars(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalendarsAPIService CalendarCalendarsUpdateCalendar", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.CalendarsAPI.CalendarCalendarsUpdateCalendar(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}