(function () {
    'use strict';
    angular.module('WidgetsApp')
        .config(configApp)
        .directive('loading', loading)
        .run(['SecurityService', '$rootScope', '$q', '$location', function (SecurityService, $rootScope, $q, $location) {

            $rootScope.baseServiceUrl = 'http://localhost:3000/';
            SecurityService.fillAuthData();
            $rootScope.logout = function () {
                SecurityService.logout();
                $location.path('/login');
            }
            $rootScope.confirm = function (titulo) {

                var prom = $q.defer();

                swal({
                    title: titulo,

                    type: 'warning',
                    showCancelButton: true,
                    confirmButtonText: 'Yes',
                    cancelButtonText: 'No'

                }, function (v) {

                    prom.resolve(v);
                });


                return prom.promise;
            };
            $rootScope.alertError = function (mensagem) {
                swal("Oops!", mensagem, "error");
            }
            $rootScope.alertSuccess = function (mensagem) {
                swal("Done!", mensagem, "success");
            }

            moment.locale("pt-BR");

            $rootScope.toLocalTime = function (hora) {
                return moment.utc((hora)).local().toDate();
            }
        }]);

    /*FUNCAO DE CONFIGURACAO*/
    function configApp($stateProvider, $urlRouterProvider, $httpProvider, $locationProvider, $translateProvider) {

        $urlRouterProvider.otherwise("/");

        $stateProvider.state('app', {
            abstract: true,
            url: "",
            templateUrl: 'main.html',
            data: { requiresLogin: true }

        })

            .state('app.home', {
                url: '/',
                templateUrl: "dashboard.html"
            })
            .state('app.user', {
                url: '/user',
                templateUrl: "user.html"
            })
            .state('app.widget', {
                url: '/widget',
                templateUrl: "widget.html"
            })
            .state('login', {

                url: "/login",
                templateUrl: "login.html",
                data: { pageTitle: 'Autenticação' },

            })

        $httpProvider.interceptors.push('authInterceptorService');
      //  $locationProvider.html5Mode(true);
    }


    loading.$inject = ['$rootScope'];
    function loading($rootScope) {


        var directive = {
            link: link,
            restrict: 'A',
            scope: true

        };
        return directive;



        function link(scope, element, model) {
            scope.$watch("isLoading", function () {

                if (scope.isLoading)
                    element.addClass("whirl traditional");
                else
                    element.removeClass("whirl traditional");
            });



        }



    }




})(window.angular);



