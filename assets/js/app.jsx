
class App extends React.Component {
    render() {
        return (
            <당첨번호/>
        )
    }
}

class 당첨번호 extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };
    }

    componentDidMount() {
        fetch("/api")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        items: result['당첨번호']
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            )
    }

    render() {
        const {error, isLoaded, items} = this.state;
        console.log(items)
        return (
            <div className="container">
                <div>
                    {items.map((value, index) => {
                        return <span className="number">{value}</span>
                    })}
                </div>
            </div>
        )
    }
}


ReactDOM.render(<App/>, document.getElementById('app'));
