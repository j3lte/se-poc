angular.module('app.services.people', [])

.service('People', function($http) {


  return {

    all: function(){
      return $http.get('/api/people');
    },

    get: function(id) {
      return $http.get('/api/people?person='+id);
    },

    add: function(data) {
      return $http.post('/api/people/add', data)
    },

  }

})