angular
.module('app.directives.person-form-address', [])


.directive('personFormAddress', function(){

  return {
  	scope: {
  		address: '=',
  		actions: '=',
  		onremove: '&'
  	},
    templateUrl: 'javascript/directives/person-form-address/template.html'    	
  }

})