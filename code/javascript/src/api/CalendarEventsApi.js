/**
 * spoke_handler
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.2.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import Event from '../model/Event';

/**
* CalendarEvents service.
* @module api/CalendarEventsApi
* @version 0.2.0
*/
export default class CalendarEventsApi {

    /**
    * Constructs a new CalendarEventsApi. 
    * @alias module:api/CalendarEventsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the calendarCalendarEventsGetEvent operation.
     * @callback module:api/CalendarEventsApi~calendarCalendarEventsGetEventCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Event} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * get event with uuid
     * @param {String} uuid 
     * @param {module:api/CalendarEventsApi~calendarCalendarEventsGetEventCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Event}
     */
    calendarCalendarEventsGetEvent(uuid, callback) {
      let postBody = null;
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling calendarCalendarEventsGetEvent");
      }

      let pathParams = {
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Event;
      return this.apiClient.callApi(
        '/calendar_events/{uuid}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the calendarCalendarEventsListEvents operation.
     * @callback module:api/CalendarEventsApi~calendarCalendarEventsListEventsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Event>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * list all events of a calendar
     * @param {Number} calendarId 
     * @param {Number} offset 
     * @param {Number} limit 
     * @param {Object} opts Optional parameters
     * @param {Number} [cursor] 
     * @param {module:api/CalendarEventsApi~calendarCalendarEventsListEventsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Event>}
     */
    calendarCalendarEventsListEvents(calendarId, offset, limit, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'calendarId' is set
      if (calendarId === undefined || calendarId === null) {
        throw new Error("Missing the required parameter 'calendarId' when calling calendarCalendarEventsListEvents");
      }
      // verify the required parameter 'offset' is set
      if (offset === undefined || offset === null) {
        throw new Error("Missing the required parameter 'offset' when calling calendarCalendarEventsListEvents");
      }
      // verify the required parameter 'limit' is set
      if (limit === undefined || limit === null) {
        throw new Error("Missing the required parameter 'limit' when calling calendarCalendarEventsListEvents");
      }

      let pathParams = {
      };
      let queryParams = {
        'calendar_id': calendarId,
        'cursor': opts['cursor'],
        'offset': offset,
        'limit': limit
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Event];
      return this.apiClient.callApi(
        '/calendar_events', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the calendarCalendarEventsScheduleRecordEvent operation.
     * @callback module:api/CalendarEventsApi~calendarCalendarEventsScheduleRecordEventCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Event} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} uuid 
     * @param {module:api/CalendarEventsApi~calendarCalendarEventsScheduleRecordEventCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Event}
     */
    calendarCalendarEventsScheduleRecordEvent(uuid, callback) {
      let postBody = null;
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling calendarCalendarEventsScheduleRecordEvent");
      }

      let pathParams = {
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Event;
      return this.apiClient.callApi(
        '/calendar_events/{uuid}/bot', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the calendarCalendarEventsUnscheduleRecordEvent operation.
     * @callback module:api/CalendarEventsApi~calendarCalendarEventsUnscheduleRecordEventCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Event} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} uuid 
     * @param {module:api/CalendarEventsApi~calendarCalendarEventsUnscheduleRecordEventCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Event}
     */
    calendarCalendarEventsUnscheduleRecordEvent(uuid, callback) {
      let postBody = null;
      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling calendarCalendarEventsUnscheduleRecordEvent");
      }

      let pathParams = {
        'uuid': uuid
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Event;
      return this.apiClient.callApi(
        '/calendar_events/{uuid}/bot', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
