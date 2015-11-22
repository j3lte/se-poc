angular
.module('app.directives.segment', [])


.directive('segmentContainer', function(){

  return {
    templateUrl: 'javascript/directives/segment/template.html',
    transclude: true,
    scope: {
    	title: '@',
    	onadd: '='
    },
    controller: function($scope) {
        if($scope.onadd) {
            $scope.showAdd = true
        }
    }
  }

})

