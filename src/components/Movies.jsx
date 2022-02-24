import React, { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Movies extends Component {
	// defining local state
	state = { movies: []}

	// executes after a component is rendered to the screen
	componentDidMount() {
		// hardwiring state for testing
		this.setState({
			movies: [
				{id:1, title:"Pusher", runtime: 150},
				{id:2, title:"Taxi Driver", runtime: 154},
				{id:3, title:"American Psycho", runtime: 142},
			]
		})
	}

	render() {
		return (
			<Fragment>
				<h2> Choose a movie</h2>

				<ul>
					{this.state.movies.map(
						(m) => (
							<li key={m.id}>
								<Link to={`/movies/${m.id}`}>
									{m.title}
								</Link>
							</li>
						)
					)}
				</ul>
			</Fragment>
		)
	}
}
