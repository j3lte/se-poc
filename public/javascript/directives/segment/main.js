angular
.module('app.directives.segment', [])


.directive('segmentContainer', function(){

  return {
    templateUrl: 'javascript/directives/segment/template.html',
    transclude: true,
    scope: {
    	title: '@',
    	array: '='
    },
    controller: function($scope) {
    	$scope.add = function() {
    		$scope.array.push({})
    	}
    }
  }

})

