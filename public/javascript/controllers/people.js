angular
.module('app.controllers.people',[])

.controller('PeopleMainListCtrl', function($scope, People) {
  $scope.people = People.all()

})

.controller('PeopleAddCtrl', function($scope, People) {

})