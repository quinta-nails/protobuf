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
(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Botsbotsproto);
  }
}(this, function(expect, Botsbotsproto) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('UsersGetByTelegramIDRequest', function() {
      beforeEach(function() {
        instance = new Botsbotsproto.UsersGetByTelegramIDRequest();
      });

      it('should create an instance of UsersGetByTelegramIDRequest', function() {
        // TODO: update the code to test UsersGetByTelegramIDRequest
        expect(instance).to.be.a(Botsbotsproto.UsersGetByTelegramIDRequest);
      });

      it('should have the property telegramId (base name: "telegramId")', function() {
        // TODO: update the code to test the property telegramId
        expect(instance).to.have.property('telegramId');
        // expect(instance.telegramId).to.be(expectedValueLiteral);
      });

    });
  });

}));
