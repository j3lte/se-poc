angular
.module('app.directives.form-relation', [])


.directive('formRelation', function(){

  return {
  	scope: {
  		relation: '=',
  		actions: '=',
  		onremove: '&'
  	},
    templateUrl: 'javascript/directives/form-relation/template.html'    	
  }

})