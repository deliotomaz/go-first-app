(function () {
    'use strict'

    angular
        .module('WidgetsApp')
        .controller('DashController', DashController)
        .controller('UserController', UserController)
        .controller('WidgetController', WidgetController)
        .controller('LoginController', LoginController)
        .controller('ProfileController', ProfileController)
        ;

    DashController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService'];
    UserController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService'];
    WidgetController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService'];
    LoginController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService'];
    ProfileController.$inject = ['$rootScope'];

    function ProfileController($rootScope) {
        var vm = this;
    }
    function LoginController($rootScope, $scope, $location, $uibModal, SecurityService) {

        var vm = this;
        vm.login = login;

        function activate() {

            if (SecurityService.authentication.isAuth)
                $location.path('/home');


        }
        function login(valid) {

            if (!valid)
                return;

            $rootScope.isLoading = true;
            SecurityService.login(vm.Email, vm.Pwd).then(function (e) {
                $rootScope.isLoading = false;
                $location.path('/home');
            }
                , function (error) {

                    $rootScope.alertError(error.message);
                    $rootScope.isLoading = false;
                })
        }
        activate();

    }

    function DashController($rootScope, $scope, $location, $uibModal, SecurityService) {

        //calendário
        var vm = this;
        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');



        }
        activate();
    };

    function UserController($rootScope, $scope, $location, $uibModal, SecurityService) {

        //calendário
        var vm = this;
        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');



        }
        activate();
    };
    function WidgetController($rootScope, $scope, $location, $uibModal, SecurityService) {

        //calendário
        var vm = this;
        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');



        }
        activate();
    };
})();