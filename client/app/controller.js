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

    DashController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService', 'UserService','WidgetService'];
    UserController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService', 'UserService'];
    WidgetController.$inject = ['$rootScope', '$scope', '$location', '$uibModal', 'SecurityService', 'WidgetService'];
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

    function DashController($rootScope, $scope, $location, $uibModal, SecurityService, UserService, WidgetService) {

        //calendário
        var vm = this;
        vm.users = [];
        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');

            $rootScope.isLoading = true;
            UserService.list().then(function (result) {

                vm.users = result;
                $rootScope.isLoading = false;
            },
                function (e) {

                    $rootScope.isLoading = false;
                }
            )

            WidgetService.list().then(function (result) {

                vm.widgets = result;
                $rootScope.isLoading = false;
            },
                function (e) {

                    $rootScope.isLoading = false;
                }
            )

        }
        activate();
    };

    function UserController($rootScope, $scope, $location, $uibModal, SecurityService, UserService) {


        var vm = this;
        vm.setSelected = setSelected;
        vm.users = [];
        vm.userSelected = null;
        vm.isSelected = false;
        vm.filter = '';
        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');

            UserService.list().then(function (result) {

                vm.users = result;
                $rootScope.isLoading = false;
            },
                function (e) {

                    $rootScope.isLoading = false;
                }
            )

        }
        function setSelected(codigo) {
            if (codigo == null) {

                vm.isSelected = false;
                vm.userSelected = null;
                return;
            }
            vm.isSelected = true;
            $scope.isLoading = true;

            UserService.getById(codigo).then(function (e) {
                vm.userSelected = e;

                $scope.isLoading = false;
            })

        }
        activate();
    };
    function WidgetController($rootScope, $scope, $location, $uibModal, SecurityService, WidgetService) {

        var vm = this;

        vm.widgets = [];
        vm.widgetSelected = null;
        vm.isSelected = false;
        vm.filter = '';
        vm.abrirPopup = abrirPopup;
        vm.load = load;

        function activate() {
            //this code could be in route change event
            if (!SecurityService.authentication.isAuth)
                $location.path('/login');

            vm.load();

        }
        function load() {
            WidgetService.list().then(function (result) {

                vm.widgets = result;
                $rootScope.isLoading = false;
            },
                function (e) {

                    $rootScope.isLoading = false;
                }
            )
        }

        function abrirPopup(item) {
            var modalInstance = $uibModal.open({
                animation: vm.animationsEnabled,
                ariaLabelledBy: 'modal-title',
                ariaDescribedBy: 'modal-body',
                templateUrl: 'newwidget.html',
                controller: ['$uibModalInstance', 'item', function ($uibModalInstance, item) {
                    var vm = this;
                    vm.novo = item == null;
                    vm.codigo = item;
                    activate();
                    vm.names = ["red", "purple", "black", "green", "magenta", "white", "depends on the viewing angle"];
                    function activate() {
                        if (vm.novo) {
                          
                            vm.item = {
                                id: 0,
                                name: "",
                                color: "",
                                price: "",
                                inventory: 0,
                                melts: false

                            }

                        }
                        else {
                            $scope.isLoading = true;
                            WidgetService.getById(item).then(function (e) {
                                vm.item = e;

                                $scope.isLoading = false;
                            })
                        }
                    }
                    vm.fecharPopup = function () {
                        $uibModalInstance.close('cancel');
                    }
                    vm.salvar = function (valido) {
                        if (!valido) {
                            $rootScope.alertError('Todos os campos são obrigatórios');
                        }
                        $scope.isLoading = true;
                        vm.item.inventory = parseInt(vm.item.inventory);
                        WidgetService.save(vm.item, vm.codigo).then(function (s) {
                            $scope.isLoading = false;
                            $rootScope.alertSuccess("Saved!");
                            vm.fecharPopup();

                        },
                            function (s) {
                                $scope.isLoading = false;
                                $rootScope.alertError("An error ocurred");
                            }
                        )


                    }
                }],
                controllerAs: 'vm',
                resolve: {
                    item: function () {
                        return item;
                    }

                }
            }).result.then(function (e) {
                vm.load();
            });
        }

        activate();
    };
})();