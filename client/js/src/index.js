/*
 * bots/bots.proto
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.58
 *
 * Do not edit the class manually.
 *
 */
import ApiClient from './ApiClient';
import BotsAddBotRequest from './model/BotsAddBotRequest';
import BotsAddBotResponse from './model/BotsAddBotResponse';
import BotsBot from './model/BotsBot';
import BotsSetupBotRequest from './model/BotsSetupBotRequest';
import BotsSetupBotResponse from './model/BotsSetupBotResponse';
import CompaniesReserveRequest from './model/CompaniesReserveRequest';
import CompaniesReserveResponse from './model/CompaniesReserveResponse';
import ProtobufAny from './model/ProtobufAny';
import RpcStatus from './model/RpcStatus';
import UsersGetByTelegramIDRequest from './model/UsersGetByTelegramIDRequest';
import UsersGetByTelegramIDResponse from './model/UsersGetByTelegramIDResponse';
import UsersUser from './model/UsersUser';
import CompaniesServiceApi from './api/CompaniesServiceApi';
import TelegramBotsServiceApi from './api/TelegramBotsServiceApi';
import UsersServiceApi from './api/UsersServiceApi';

/**
* Object.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var Botsbotsproto = require('index'); // See note below*.
* var xxxSvc = new Botsbotsproto.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new Botsbotsproto.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new Botsbotsproto.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new Botsbotsproto.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version version not set
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The BotsAddBotRequest model constructor.
     * @property {module:model/BotsAddBotRequest}
     */
    BotsAddBotRequest,

    /**
     * The BotsAddBotResponse model constructor.
     * @property {module:model/BotsAddBotResponse}
     */
    BotsAddBotResponse,

    /**
     * The BotsBot model constructor.
     * @property {module:model/BotsBot}
     */
    BotsBot,

    /**
     * The BotsSetupBotRequest model constructor.
     * @property {module:model/BotsSetupBotRequest}
     */
    BotsSetupBotRequest,

    /**
     * The BotsSetupBotResponse model constructor.
     * @property {module:model/BotsSetupBotResponse}
     */
    BotsSetupBotResponse,

    /**
     * The CompaniesReserveRequest model constructor.
     * @property {module:model/CompaniesReserveRequest}
     */
    CompaniesReserveRequest,

    /**
     * The CompaniesReserveResponse model constructor.
     * @property {module:model/CompaniesReserveResponse}
     */
    CompaniesReserveResponse,

    /**
     * The ProtobufAny model constructor.
     * @property {module:model/ProtobufAny}
     */
    ProtobufAny,

    /**
     * The RpcStatus model constructor.
     * @property {module:model/RpcStatus}
     */
    RpcStatus,

    /**
     * The UsersGetByTelegramIDRequest model constructor.
     * @property {module:model/UsersGetByTelegramIDRequest}
     */
    UsersGetByTelegramIDRequest,

    /**
     * The UsersGetByTelegramIDResponse model constructor.
     * @property {module:model/UsersGetByTelegramIDResponse}
     */
    UsersGetByTelegramIDResponse,

    /**
     * The UsersUser model constructor.
     * @property {module:model/UsersUser}
     */
    UsersUser,

    /**
    * The CompaniesServiceApi service constructor.
    * @property {module:api/CompaniesServiceApi}
    */
    CompaniesServiceApi,

    /**
    * The TelegramBotsServiceApi service constructor.
    * @property {module:api/TelegramBotsServiceApi}
    */
    TelegramBotsServiceApi,

    /**
    * The UsersServiceApi service constructor.
    * @property {module:api/UsersServiceApi}
    */
    UsersServiceApi
};
