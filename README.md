# Documentation for spoke_handler

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *CalendarEventsApi* | [**calendarCalendarEventsGetEvent**](Apis/CalendarEventsApi.md#calendarcalendareventsgetevent) | **GET** /calendar_events/{uuid} | get event with uuid |
*CalendarEventsApi* | [**calendarCalendarEventsListEvents**](Apis/CalendarEventsApi.md#calendarcalendareventslistevents) | **GET** /calendar_events | list all events of a calendar |
*CalendarEventsApi* | [**calendarCalendarEventsScheduleRecordEvent**](Apis/CalendarEventsApi.md#calendarcalendareventsschedulerecordevent) | **POST** /calendar_events/{uuid}/bot |  |
*CalendarEventsApi* | [**calendarCalendarEventsUnscheduleRecordEvent**](Apis/CalendarEventsApi.md#calendarcalendareventsunschedulerecordevent) | **DELETE** /calendar_events/{uuid}/bot |  |
| *CalendarsApi* | [**calendarCalendarsCreateCalendar**](Apis/CalendarsApi.md#calendarcalendarscreatecalendar) | **POST** /calendars | synchronize all events of a new calendar |
*CalendarsApi* | [**calendarCalendarsDeleteCalendar**](Apis/CalendarsApi.md#calendarcalendarsdeletecalendar) | **DELETE** /calendars/{uuid} |  |
*CalendarsApi* | [**calendarCalendarsGetCalendar**](Apis/CalendarsApi.md#calendarcalendarsgetcalendar) | **GET** /calendars/{uuid} |  |
*CalendarsApi* | [**calendarCalendarsListCalendars**](Apis/CalendarsApi.md#calendarcalendarslistcalendars) | **GET** /calendars |  |
*CalendarsApi* | [**calendarCalendarsUpdateCalendar**](Apis/CalendarsApi.md#calendarcalendarsupdatecalendar) | **PATCH** /calendars/{uuid} |  |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [CreateCalendarParams](./Models/CreateCalendarParams.md)
 - [Event](./Models/Event.md)
 - [EventStatus](./Models/EventStatus.md)
 - [Provider](./Models/Provider.md)
 - [SystemTime](./Models/SystemTime.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
