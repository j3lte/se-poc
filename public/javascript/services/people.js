angular.module('app.services.people', [])

.service('People', function($http) {

  function mutator(person) {

    if(person.birthdate) person.birthdate = new Date(person.birthdate)
    if(!person.accounts) person.accounts = []
    if(!person.phonenumbers) person.phonenumbers = []
    if(!person.addresses) person.addresses = []

    return person;
  }

  function arrayMutator(data) {
    return data.map(mutator)
  }

  return {

    // Returns all people from the API
    all: function(){
      return $http.get('/api/people')
      .success(arrayMutator)
    },

    // Returns a person based on the id passed in
    get: function(id) {
      return $http.get('/api/people?person='+id)
      .success(mutator)
    },

    // Saves the person
    save: function(data) {
      return $http.post('/api/people/save', data)
    },

    // Adds a new person
    add: function(data) {
      return $http.post('/api/people/add', data)
    },

    // Generates an empty person
    new: function() {
      return $http.get('/api/people/new')
      .success(mutator)
    }

  }

})