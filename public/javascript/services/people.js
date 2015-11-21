angular.module('app.services.people', [])

.service('People', function() {
  var people = [
    {
      firstname: 'Vincent',
      lastname: 'Swarte',
      birthdate: new Date(1995, 11, 28)
    },
    {
      firstname: 'Gaia',
      lastname: 'van Basten',
      birthdate: new Date(1993, 8, 3)
    }
  ];

  return {
    all: function(){
      return people;
    }

  }

})