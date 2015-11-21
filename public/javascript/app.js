angular
.module('app', [
  'ngRoute',

  /* Controllers */
  'app.controllers.people',

  /* Services */
  'app.services.people'
])


// Routes
.config(function($routeProvider, $locationProvider) {

  $routeProvider
  .when('/people', {
    templateUrl: 'views/people/main-list.html',
    controller: 'PeopleMainListCtrl'
  })
  .when('/people/new', {
    templateUrl: 'views/people/add.html',
    controller: 'PeopleAddCtrl'
  })
  .when('/people/:id/view', {
    templateUrl: 'views/people/view.html',
    // controller: 'PeopleViewCtrl'
  })

  .otherwise('/people')
});