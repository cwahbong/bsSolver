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
    $scope.$watch('size', function(newValue, oldValue) {
      $scope.map = [];
      for (var i = 0; i < newValue; ++i) {
        $scope.map[i] = [];
        for (var j = 0; j < newValue; ++j) {
          $scope.map[i][j] = -1;
        }
      }
    });
    $scope.$watch('selected_code', function(newValue, oldValue) {
      $scope.colors[oldValue-1].classes.selected = false;
      $scope.colors[newValue-1].classes.selected = true;
    });
    var colorMap = [
      "none",
      "red",
      "blue",
      "yellow,"
    ];
    $scope.color = function(colorCode) {
      return colorMap[colorCode];
    };
    $scope.colors = [];
    for (var i = 1; i <= 3; ++i) {
      var classes = {selected: false};
      classes[$scope.color(i)] = true;
      var color = {code: i};
      $scope.colors.push({code: i, classes: classes});
    }
    $scope.selected_code = 1;
    $scope.select = function(colorCode) {
      $scope.selected_code = colorCode;
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
