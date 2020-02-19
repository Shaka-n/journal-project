import React from 'react';

class EssayForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      title: 'Tell me about your day.',
      body: 'Enter your thoughts here.'
    };


    this.handleChangeBody = this.handleChangeBody.bind(this);
    this.handleChangeTitle = this.handleChangeTitle.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }


  handleChangeTitle(event) {
    this.setState({title:event.target.value});
  }

  handleChangeBody(event){
    this.setState({body:event.target.value})
  }


  handleSubmit(event){
    alert('Thanks for sharing.');
    event.preventDefault();
    fetch('localhost:8080/entries/', {method: "POST", body: JSON.stringify(this.state)});
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <textarea id="title" value={this.state.title} onChange={this.handleChangeTitle}/>
        <p></p>
        <textarea id="body" value={this.state.body} onChange={this.handleChangeBody}/>
        <input type="submit" value="Submit" />
      </form>
    );
  }
}
export default EssayForm;
//ReactDOM.render(<EssayForm/>, document.getElementById('root'));
