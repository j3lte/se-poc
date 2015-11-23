angular
.module('app.directives.network-viewer', [])


.directive('networkViewer', function(){

  return {
    scope: {
        nodes: '=',
        edges: '='
    },
    templateUrl: 'javascript/directives/network-viewer/template.html',
    link: function($scope, $element) {
        // Recreate network once variables are updates
        $scope.$watch('nodes', createNetwork)
        $scope.$watch('edges', createNetwork)

        function createNetwork(){
            var color = 'gray';
            var len = undefined;

            // create a network
            var container = $element[0]
            var data = {
                nodes: $scope.nodes,
                edges: $scope.edges
            };
            var options = {
                height: '500px',
                nodes: {
                    shape: 'dot',
                    size: 20,
                    font: {
                        size: 15,
                        color: '#000'
                    },
                    borderWidth: 2
                },
                edges: {
                    width: 2
                },
                groups: {
                    people: {
                        shape: 'icon',
                        icon: {
                            face: 'Icons',
                            code: '\uf007',
                            color: '#000'
                        }
                    },
                    address: {
                        shape: 'icon',
                        icon: {
                            face: 'Icons',
                            code: '\uf015',
                            color: '#000'
                        }
                    },
                    account: {
                        shape: 'icon',
                        icon: {
                            face: 'Icons',
                            code: '\uf084',
                            color: '#000'
                        }
                    }

                }
            };
            network = new vis.Network(container, data, options);
        }
    }  
  }

})
