(function () {
    'use strict'

    angular
        .module('WidgetsApp')
        .factory('SecurityService', SecurityService)
        .factory('authInterceptorService', authInterceptorService)
        .factory('UserService', UserService)
        .factory('WidgetService', WidgetService);


    WidgetService.$inject = ['$resource', '$rootScope'];
    function WidgetService($resource, $rootScope) {
        var vm = $resource($rootScope.baseServiceUrl + 'widgets/:id', null, {
            update: {
                method: 'PUT'
            },
            get: {

                method: 'GET',
                isArray: true
            },
            getById: {

                method: 'GET',
                url: $rootScope.baseServiceUrl + 'widgets/:id'
            },
            //update: {
            //    method: 'PUT',
            //    url: $rootScope.baseServiceUrl + 'widgets/:id'
            //},
            save: {
                method: 'POST'
            }
        });

        this.save = function (obj, codigo) {
            if (codigo == null)
                return vm.save(obj).$promise;
            else {
                return vm.update({ id: codigo }, obj).$promise;
            }
        }
        this.getById = function (id) {
            return vm.getById({ id: id }).$promise;
        }
        this.list = function () {
            return vm.get().$promise;
        }

        return this;

    };
    UserService.$inject = ['$resource', '$rootScope'];
    function UserService($resource, $rootScope) {
        var vm = $resource($rootScope.baseServiceUrl + 'users/:id', null, {
            update: {
                method: 'PUT'
            },
            get: {

                method: 'GET',
                isArray: true
            },
            getById: {

                method: 'GET',
                url: $rootScope.baseServiceUrl + 'users/:id'
            },
        });


        this.getById = function (id) {
            return vm.getById({ id: id }).$promise;
        }
        this.list = function () {
            return vm.get().$promise;
        }

        return this;

    };

    SecurityService.$inject = ['$resource', '$rootScope', 'localStorageService', '$q'];
    function SecurityService($resource, $rootScope, localStorageService, $q) {
        this.poolData = {
            UserPoolId: 'us-east-1_xWpxgTXO5',
            ClientId: '3jrkv7kje6m064midc8131fpdj'
        };
        this.authentication = {
            isAuth: false,
            userName: ""
        };




        this.login = function (email, password) {
            var authenticationData = {
                Username: email,
                Password: password,
            };
            var authenticationDetails = new AmazonCognitoIdentity.AuthenticationDetails(authenticationData);
            var userPool = new AmazonCognitoIdentity.CognitoUserPool(this.poolData);
            var userData = {
                Username: email,
                Pool: userPool
            };
            var cognitoUser = new AmazonCognitoIdentity.CognitoUser(userData);

            var deferred = $q.defer();
            var _this = this;
            cognitoUser.authenticateUser(authenticationDetails, {
                onSuccess: function (result) {

                    var accessToken = result.getIdToken().jwtToken;
                    localStorageService.set('authorizationData', { token: accessToken, userName: result.getIdToken().payload.name });
                    _this.authentication.isAuth = true;
                    _this.authentication.userName = result.getIdToken().payload.name;
                    deferred.resolve(_this.authentication);

                },

                onFailure: function (err) {
                    _this.authentication.isAuth = false;
                    deferred.reject(err);
                }

            });

            return deferred.promise;
        }





        this.logout = function () {
            localStorageService.remove('authorizationData');

            this.authentication.isAuth = false;
            this.authentication.userName = "";
        }

        this.fillAuthData = function () {
            var authData = localStorageService.get('authorizationData');
            if (authData) {
                this.authentication.isAuth = true;
                this.authentication.userName = authData.userName;
            }
        }

        return this;

    };

    authInterceptorService.$inject = ['$q', '$injector', '$location', 'localStorageService'];
    function authInterceptorService($q, $injector, $location, localStorageService) {

        var authInterceptorServiceFactory = {};

        var _request = function (config) {

            config.headers = config.headers || {};

            var authData = localStorageService.get('authorizationData');
            if (authData) {
                config.headers.Authorization = 'bearer ' + authData.token;
            }

            return config;
        }

        var _responseError = function (rejection) {
            if (rejection.status === 401) {
                var authService = $injector.get('SecurityService');
                var authData = localStorageService.get('authorizationData');

                if (authData) {
                    if (authData.useRefreshTokens) {
                        $location.path('/refresh');
                        return $q.reject(rejection);
                    }
                }
                authService.logout();
                $location.path('/login');
            }
            return $q.reject(rejection);
        }

        authInterceptorServiceFactory.request = _request;
        authInterceptorServiceFactory.responseError = _responseError;

        return authInterceptorServiceFactory;
    }




})();