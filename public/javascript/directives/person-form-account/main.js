angular
.module('app.directives.person-form-account', [])


.directive('personFormAccount', function(){

  return {
  	scope: {
  		account: '=',
  		actions: '=',
  		onremove: '&'
  	},
    templateUrl: 'javascript/directives/person-form-account/template.html'    	
  }

})