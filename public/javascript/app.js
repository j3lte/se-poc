angular
.module('app', [
  'ngRoute',

  /* Controllers */
  'app.controllers.people',

  /* Services */
  'app.services.people',

  /* Directives */
  'app.directives.person-form-general',
  'app.directives.person-form-account',
  'app.directives.segment'
])


// Routes
.config(function($routeProvider, $locationProvider) {

  $routeProvider
  .when('/people', {
    templateUrl: 'views/people/main-list.html',
    controller: 'PeopleMainListCtrl'
  })
  .when('/people/new', {
    templateUrl: 'views/people/form.html',
    controller: 'PeopleAddCtrl'
  })
  .when('/people/:id/view', {
    templateUrl: 'views/people/form.html',
    controller: 'PeopleViewCtrl'
  })

  .otherwise('/people')
});