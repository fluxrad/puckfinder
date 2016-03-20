var RinkListBox = React.createClass({
  getInitialState: function() {
    return {data: []};
  },

  loadRinksFromServer: function() {
    $.ajax({
      url: this.props.url,
      dataType: 'json',
      cache: false,
      success: function(data) {
        this.setState({data: data.rinks});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },

  componentDidMount: function() {
    this.loadRinksFromServer();
    setInterval(this.loadRinksFromServer, this.props.pollInterval);
  },

  render: function() {
    return (
      <div className="rinkListBox">
        <h1>Rinks</h1>
        <RinkList data={this.state.data} />
      </div>
    );
  }


});

var RinkList = React.createClass({
  render: function() {
    var rinkNodes = this.props.data.map(function(rink) {
      return (
        <RinkButton id={rink.id} >
          { rink.shortName }
        </RinkButton>
      );
    });

    return (
        <div className="rinkList">
          {rinkNodes}
        </div>
    );
  }
});

var RinkButton = React.createClass({
  render: function() {
    return (
      <div className="rink">
        <h2 className="rinkName">
          {this.props.rinkName}
        </h2>
        <a className="btn btn-default btn-sm" href="#" role="button">{this.props.children}</a>
      </div>
    )
  }
});

ReactDOM.render(
  <RinkListBox url="/api/rinks" pollInterval={60000} />,
  document.getElementById('content')
);
