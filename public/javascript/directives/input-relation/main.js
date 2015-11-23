angular
.module('app.directives.input-relation', [])


.directive('inputRelation', function(People){

  return {
  	scope: {
  		ngModel: '='
  	},
    templateUrl: 'javascript/directives/input-relation/template.html',
    link: function($scope){

    	People.allSearch()
    	.success(function(data){
    		$scope.items = data;
    	})

    }
  }

})