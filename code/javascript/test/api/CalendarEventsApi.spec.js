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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.SpokeHandler);
  }
}(this, function(expect, SpokeHandler) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new SpokeHandler.CalendarEventsApi();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('CalendarEventsApi', function() {
    describe('calendarCalendarEventsGetEvent', function() {
      it('should call calendarCalendarEventsGetEvent successfully', function(done) {
        //uncomment below and update the code to test calendarCalendarEventsGetEvent
        //instance.calendarCalendarEventsGetEvent(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('calendarCalendarEventsListEvents', function() {
      it('should call calendarCalendarEventsListEvents successfully', function(done) {
        //uncomment below and update the code to test calendarCalendarEventsListEvents
        //instance.calendarCalendarEventsListEvents(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('calendarCalendarEventsScheduleRecordEvent', function() {
      it('should call calendarCalendarEventsScheduleRecordEvent successfully', function(done) {
        //uncomment below and update the code to test calendarCalendarEventsScheduleRecordEvent
        //instance.calendarCalendarEventsScheduleRecordEvent(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('calendarCalendarEventsUnscheduleRecordEvent', function() {
      it('should call calendarCalendarEventsUnscheduleRecordEvent successfully', function(done) {
        //uncomment below and update the code to test calendarCalendarEventsUnscheduleRecordEvent
        //instance.calendarCalendarEventsUnscheduleRecordEvent(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
