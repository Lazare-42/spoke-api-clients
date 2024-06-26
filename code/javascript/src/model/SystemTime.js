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

import ApiClient from '../ApiClient';

/**
 * The SystemTime model module.
 * @module model/SystemTime
 * @version 0.2.0
 */
class SystemTime {
    /**
     * Constructs a new <code>SystemTime</code>.
     * @alias module:model/SystemTime
     * @param secsSinceEpoch {Number} 
     * @param nanosSinceEpoch {Number} 
     */
    constructor(secsSinceEpoch, nanosSinceEpoch) { 
        
        SystemTime.initialize(this, secsSinceEpoch, nanosSinceEpoch);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, secsSinceEpoch, nanosSinceEpoch) { 
        obj['secs_since_epoch'] = secsSinceEpoch;
        obj['nanos_since_epoch'] = nanosSinceEpoch;
    }

    /**
     * Constructs a <code>SystemTime</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SystemTime} obj Optional instance to populate.
     * @return {module:model/SystemTime} The populated <code>SystemTime</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SystemTime();

            if (data.hasOwnProperty('secs_since_epoch')) {
                obj['secs_since_epoch'] = ApiClient.convertToType(data['secs_since_epoch'], 'Number');
            }
            if (data.hasOwnProperty('nanos_since_epoch')) {
                obj['nanos_since_epoch'] = ApiClient.convertToType(data['nanos_since_epoch'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>SystemTime</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>SystemTime</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of SystemTime.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }

        return true;
    }


}

SystemTime.RequiredProperties = ["secs_since_epoch", "nanos_since_epoch"];

/**
 * @member {Number} secs_since_epoch
 */
SystemTime.prototype['secs_since_epoch'] = undefined;

/**
 * @member {Number} nanos_since_epoch
 */
SystemTime.prototype['nanos_since_epoch'] = undefined;






export default SystemTime;

