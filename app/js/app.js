'use strict';

angular.module('varfuri', ['ngResource'])
.config(function($interpolateProvider){
        $interpolateProvider.startSymbol('{[{').endSymbol('}]}');
})
.factory('Map', function($resource){
	return $resource('/maps', {}, {
		query: {method:'GET', params:{}, isArray:true}
	});
})
.controller('VarfuriCtrl', function($scope, $http, Map) {
  	$scope.maps = Map.query();
})
.controller('ContactCtrl', function($scope, $http, Map) {
  	$scope.contactResponse = "";

  	$scope.sendMessage = function(mes){
  		if (typeof mes === 'undefined') {
  			$scope.contactResponse = "Completati campurile mesajului!";
		} else {
  			$http({
  				url: "/contact",
  				method: 'POST',
  				data: $.param(mes),
  				headers: {'Content-Type': 'application/x-www-form-urlencoded'}})
			.success(function(data, status, headers, config) {
			    $scope.contactResponse = "Multumesc pentru mesaj!";
			}).error(function(data, status, headers, config) {
			    $scope.contactResponse = "Eroare: " + status;
			});
  		}
    }
});
