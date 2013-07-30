'use strict';

angular.module('varfuri', ['ngResource'])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/', {templateUrl: 'partials/index.html', controller: 'VarfuriCtrl'});
        $routeProvider.when('/harta/:hartaImg', {templateUrl: 'partials/harta.html', controller: 'HartaCtrl'});
        $routeProvider.when('/contact', {templateUrl: 'partials/contact.html', controller: 'ContactCtrl'});
        $routeProvider.otherwise({redirectTo: '/'});
    }])
    .factory('Map', function ($resource) {
        return $resource('/api/maps', {}, {
            query: {method: 'GET', params: {}, isArray: true}
        });
    })
    .controller('VarfuriCtrl', function ($scope, Map) {
        $scope.maps = Map.query();
    })
    .controller('HartaCtrl', function ($scope, $routeParams) {
        $scope.hartaImg = $routeParams.hartaImg;
    })
    .controller('ContactCtrl', function ($scope, $http) {
        $scope.contactResponse = "";

        $scope.sendMessage = function (mes) {
            if (typeof mes === 'undefined') {
                $scope.contactResponse = "Completati campurile mesajului!";
            } else {
                $http({
                    url: "/api/contact",
                    method: 'POST',
                    data: $.param(mes),
                    headers: {'Content-Type': 'application/x-www-form-urlencoded'}})
                    .success(function (data, status, headers, config) {
                        $scope.contactResponse = "Multumesc pentru mesaj!";
                    }).error(function (data, status, headers, config) {
                        $scope.contactResponse = "Eroare: " + status;
                    });
            }
        }
    });
