angular
.module('app.controllers.people',[])

.controller('PeopleMainListCtrl', function($scope, People, $window) {

  People.all()
  .success(function(data){
    $scope.people = data
  })

  $scope.OpenPerson = function(id) {
    $window.location = '#/people/' + id + '/view'
  }

})

.controller('PeopleAddCtrl', function($scope, People, $window) {

  $scope.input = {}

  $scope.formSubmit = function() {
    People.add($scope.input)
    .then(function(){
      $window.location('#/people/')
    })
  }

})


.controller('PeopleViewCtrl', function($scope, $routeParams, People) {

  People.get($routeParams.id)
  .success(function(data){
    $scope.input = data
  })

  $scope.formSubmit = function() {
    People.save($scope.input)
    .then(function(){
      console.debug('Done saving');
    })
  }

})