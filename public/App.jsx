class App extends React.Component {
	render() {
		return (
			<div class="container">
				<TodoHeader />
				<TodoForm />
			</div>
		)
	}
}

class TodoHeader extends React.Component {
	render() {
		return (
			<div class="todoHeader">
				<p class="text-center fs-1 fw-bold">Todo App<span class="badge bg-warning">TEST</span></p>
			</div>
		)
	}
}

class TodoForm extends React.Component {
	constructor(props){
		super(props);
		this.state={
			list: [],
		};
	}
	componentDidMount() {

		var templist = [];
		var addItem = function(item){
			var temp = <Todos id={item.id} msg={item.message} />		
				return temp	
		}.bind(this)

		axios.get('/todos')
			.then(resp => {
				resp["data"].forEach(e => {
					var item = addItem(e)
					templist.push(item)
				});

				this.setState({
					list: templist
				})
				templist = [];
			});
	}
	render() {
		const list = this.state.list;
		$(document).on('click','.trash2',function(event) {
			var id = $(this).closest("li").attr('id');
			if(id){
				var $self = $(this);
				axios.delete('todos/'+id)
					.then(resp => {
						if (resp["data"].success === true){
							$self.parent().remove();
						}
					});
			}
		})

		function createHandler(event) {
			event.preventDefault();
			var item = $('#todoInput').val();
			if (item)
			{
				let form = new FormData();
				form.append('msg',item);
				axios.post("/todos",form)
					.then(resp => {
						addItem(resp["data"]);
						$('#todoInput').val("");
					})
			}
		}	

		var addItem = function(item){
			var temp = <Todos id={item.id} msg={item.message}  />		
				list.push(temp)
			this.setState({
				list: list
			})

		}.bind(this)



		return (

			<form id="puttodos" onSubmit= {createHandler}>
				<div class="mb-3">
					<label for="todoInput" class="form-label">Todos</label>
					<input type="text" class="form-control" id="todoInput" placeholder="Input your Todos"/>
				</div>
				<button type="submit" class="btn btn-primary">submit</button>
				<div class="todoList">
					<ul class="list-group" id="listTodo">
						<div>
							{this.state.list}
						</div>
					</ul>
				</div>
			</form>
		)
	}
}


class Todos extends React.Component {
	render() {
		return (
			<li id={this.props.id} class="list-group-item todoMsg"><input class="form-check-input" type="checkbox" value="" aria-label="..." /> {this.props.msg} <img class="trash2" src="./trash2.svg" alt="Bootstrap" width="24" height="24" /></li>
		)
	}	

}
ReactDOM.render(<App />, document.getElementById('root'));
