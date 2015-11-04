(function() {

  var CustomersController = function($scope, customersService, appSettings) {
    $scope.sortBy = 'name';
    $scope.reverse = false;
    $scope.customers = [];

    $scope.appSettings = null;

    function init() {
      $scope.customers = customersService.getCustomers();
      $scope.appSettings = appSettings;
    }
    init();

    $scope.doSort = function(propName) {
      $scope.sortBy = propName;
      $scope.reverse = !$scope.reverse;
    };

  };

  CustomersController.$inject = ['$scope', 'customersService', 'appSettings'];

  angular.module('customerApp')
    .controller('CustomersController', CustomersController);

}());
