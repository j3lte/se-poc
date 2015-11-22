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

  People.new()
  .success(function(person){
    console.debug(person)
    $scope.input = person
  })

  $scope.formSubmit = function() {
    People.add($scope.input)
    .then(function(){
      $window.location('#/people/')
    })
  }

  $scope.removeAccount = function(account) {
    $scope.input.accounts.splice($scope.input.accounts.indexOf(account), 1)
  }

  $scope.addAccount = function() {
    $scope.input.accounts.push({})
  }

})


.controller('PeopleViewCtrl', function($scope, $routeParams, People) {

  People.get($routeParams.id)
  .success(function(data){
    $scope.input = data
  })

  $scope.removeAccount = function(account) {
    $scope.input.accounts.splice($scope.input.accounts.indexOf(account), 1)
  }

  $scope.addAccount = function() {
    $scope.input.accounts.push({})
  }

  $scope.formSubmit = function() {
    People.save($scope.input)
    .then(function(){
      console.debug('Done saving');
    })
  }

})