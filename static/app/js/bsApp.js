'use strict';

// Declare app level module which depends on filters, and services
angular.module('bsApp', ['bsApp.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {
      templateUrl: 'partials/bs.html',
      controller: 'BsCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);

angular.module('bsApp.controllers', []).
  controller('BsCtrl', function($scope, $http) {
    $scope.$watch(
      'size', function(newValue, oldValue) {
        $scope.map = [];
        for (var i = 0; i < newValue; ++i) {
          $scope.map[i] = [];
          for (var j = 0; j < newValue; ++j) {
            $scope.map[i][j] = -1;
          }
        }
      }
    );
    var colorMap = [
      "none",
      "red",
      "blue",
      "yellow",
    ]
    $scope.color = function(colorCode) {
      return colorMap[colorCode];
    };
    $scope.submit = function() {
      $http.post('/j', {
        method: "bs.Solve",
        params: [{"board": $scope.map}],
        id: 0,
      }).
        success(function(data, status, headers, config) {
          $scope.lines = angular.fromJson(data).result.lines;
        }).
        error(function(data, status, headers, config) {
          $scope.lines = ["Load failed."];
        });
    };
  });
