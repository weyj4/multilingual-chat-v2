import React, { Component } from 'react'
import store from './messageStore'

export default class MessageInput extends Component {
  constructor(props) {
    super(props)
    this.state = { value: '' }
    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleChange = this.handleChange.bind(this)
  }

  handleSubmit(event) {
    event.preventDefault()
    store.addMessage(this.state.value)
    this.state.value = ''
  }

  handleChange(event) {
    this.setState({ value: event.target.value })
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Input Message:
          <input type='text' value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type='submit' value='Submit' />
      </form>
    )
  }
}
          

