angular.module('app.services.people', [])

.service('People', function($http) {

  function mutator(person) {
    if(person.birthdate) person.birthdate = new Date(person.birthdate)
    if(!person.accounts) person.accounts = []
    if(!person.phonenumbers) person.phonenumbers = []
    if(!person.addresses) person.addresses = []
    if(!person.relations) person.relations = []

    return person;
  }

  function arrayMutator(data) {
    return data.map(mutator)
  }

  function formatForSearch(person) {
    return {
      title: person.firstname + ' ' + person.lastname
    };
  }

  function arrayFormatForSearch(data) {
    return data.map(formatForSearch)
  }

  return {

    // Returns all people from the API
    all: function(){
      return $http.get('/api/people')
      .success(arrayMutator)
    },

    // Returns the all() output, only mapped for search and typeaheads
    allSearch: function() {
      return $http.get('/api/people')
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
    },

    // Generates a network map
    generateNetworkMap: function(id) {
      return $http.get('/api/people/map?person='+id)
    }

  }

})