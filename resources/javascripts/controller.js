var app = angular.module('goApp', []);


app.controller('goAppCtrl', ['$scope','$http', function($scope,$http){

		$scope.todos= [];

		$scope.title = "ToDo GO Application ";

		$scope.getTodos = function(){

			$http.get('http://localhost:4000/v1/todos').
			  success(function(data, status, headers, config) {
			    // this callback will be called asynchronously
			    // when the response is available
			    console.log(data);
			    if(data)
			    	$scope.todos = data;
			  }).
			  error(function(data, status, headers, config) {
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			  });
		}

		$scope.getTodos();

		$scope.addToDo = function(todo){
			var body = {
				name :todo,
				completed :false
			}

			$http.post('http://localhost:4000/v1/todos',body).
			  success(function(data, status, headers, config) {
			    // this callback will be called asynchronously
			    // when the response is available
			    console.log(data);
			    $scope.todos.push(data);
			    delete $scope.todo 
			  }).
			  error(function(data, status, headers, config) {
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			  });
		}

		$scope.deleteToDo = function(todo){
			$http.delete('http://localhost:4000/v1/todos/'+todo.id).
			  success(function(data, status, headers, config) {
			    // this callback will be called asynchronously
			    // when the response is available
			    console.log(headers);
			    // $scope.todos = data;
			    for(var i=0;i<$scope.todos.length;i++){
			    	if($scope.todos[i].id==todo.id){
			    		$scope.todos.splice(i,1);
			    		break;
			    	}

			    }
			  }).
			  error(function(data, status, headers, config) {
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			  });
		}

		$scope.updateToDo = function(todo,completed){

			$http.put('http://localhost:4000/v1/todos/'+todo.id,{completed:todo.completed}).
			  success(function(data, status, headers, config) {
			    // this callback will be called asynchronously
			    // when the response is available
			    console.log(data);
			     
			  }).
			  error(function(data, status, headers, config) {
			  	todo.completed = false;
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			  });
		}
}])