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
* Enum class Provider.
* @enum {}
* @readonly
*/
export default class Provider {
    
        /**
         * value: "Google"
         * @const
         */
        "Google" = "Google";

    
        /**
         * value: "Microsoft"
         * @const
         */
        "Microsoft" = "Microsoft";

    

    /**
    * Returns a <code>Provider</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/Provider} The enum <code>Provider</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

